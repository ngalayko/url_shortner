package links

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ngalayko/url_shortner/server/dao/tables"
	"github.com/ngalayko/url_shortner/server/helpers"
	"github.com/ngalayko/url_shortner/server/logger"
	"github.com/ngalayko/url_shortner/server/schema"
	"go.uber.org/zap"
)

const (
	defaultExpire      = 24 * time.Hour
	defaultShortUrlLen = 6
	httpScheme         = "http"
)

// Links is a links service
type Links struct {
	logger *logger.Logger
	tables *tables.Tables

	//channel of link ids to increment views of
	viewsQueue chan uint64
}

func newLinks(ctx context.Context) *Links {
	l := &Links{
		logger: logger.FromContext(ctx),
		tables: tables.FromContext(ctx),

		viewsQueue: make(chan uint64),
	}

	go l.loop()

	return l
}

func (l *Links) loop() {
	for id := range l.viewsQueue {
		if err := l.incrementNextLink(id); err != nil {
			l.logger.Error("error incrementing link",
				zap.Error(err),
			)
		}
	}
}

func (l *Links) incrementNextLink(linkId uint64) error {
	link, err := l.tables.SelectLinkByFields(map[string]interface{}{"id": linkId})
	if err != nil {
		return err
	}

	link.Views += 1
	if err := l.tables.UpdateLink(link); err != nil {
		return err
	}

	l.logger.Info("link views incremented",
		zap.Uint64("id", link.ID),
		zap.Uint64("views", link.Views),
	)

	return nil
}

// CreateLink creates given link
func (l *Links) CreateLink(link *schema.Link) error {

	if err := prepareLink(link); err != nil {
		return err
	}

	return l.tables.InsertLink(link)
}

func prepareLink(link *schema.Link) error {
	if !strings.HasPrefix(link.URL, httpScheme) {
		link.URL = httpScheme + "://" + link.URL
	}

	uri, err := url.ParseRequestURI(link.URL)
	if err != nil {
		return err
	}

	now := time.Now()
	link.URL = uri.String()
	link.CreatedAt = now
	link.ShortURL = helpers.RandomString(defaultShortUrlLen)
	link.ExpiredAt = now.Add(defaultExpire)

	return nil
}

// QueryLinkByShortUrl returns link by short url
func (l *Links) QueryLinkByShortUrl(shortUrl string) (*schema.Link, error) {

	link, err := l.tables.SelectLinkByFields(map[string]interface{}{"short_url": shortUrl})
	if err != nil {
		return nil, err
	}

	if link.ExpiredAt.Before(time.Now()) {
		return link, fmt.Errorf("Link has expired at %s", link.ExpiredAt)
	}

	l.viewsQueue <- link.ID

	return link, nil
}
