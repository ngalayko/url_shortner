package user_token

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	. "gopkg.in/check.v1"

	"github.com/ngalayko/url_shortner/server/cache"
	"github.com/ngalayko/url_shortner/server/config"
	"github.com/ngalayko/url_shortner/server/dao/migrate"
	"github.com/ngalayko/url_shortner/server/logger"
	"github.com/ngalayko/url_shortner/server/schema"
)

type TestTokensSuite struct {
	ctx context.Context

	service *Service

	usersCount uint64
}

func Test(t *testing.T) { TestingT(t) }

var suite *TestTokensSuite

var _ = Suite(&TestTokensSuite{})

func (s *TestTokensSuite) SetUpSuite(c *C) {
	suite = &TestTokensSuite{
		ctx: context.Background(),
	}

	s.init()

	m := migrate.FromContext(s.ctx)
	if err := m.Flush(); err != nil {
		c.Fatal(err)
	}

	if err := m.Apply(); err != nil {
		log.Panicf("error applying migrations: %s", err)
	}
}

func (s *TestTokensSuite) init() {
	s.ctx = cache.NewContext(nil, cache.NewStubCache())
	s.ctx = logger.NewContext(s.ctx, logger.NewTestLogger())
	s.ctx = config.NewContext(s.ctx, config.NewTestConfig())
	s.ctx = migrate.NewContext(s.ctx, nil)

	s.service = FromContext(s.ctx)
}

func (s *TestTokensSuite) Test_CreateUserToken__should_create_user_token(c *C) {
	user, err := s.createUser()
	c.Assert(err, IsNil)

	token, err := s.service.CreateUserToken(user)
	if err != nil {
		c.Fatal(err)
	}

	selected, err := s.service.GetUserToken(token.Token)
	if err != nil {
		c.Fatal(err)
	}

	c.Assert(token.UserID, Equals, user.ID)
	c.Assert(token.Token, Equals, selected.Token)
	c.Assert(token.ID, Equals, selected.ID)
}

func (s *TestTokensSuite) Test_DeleteUserToken__should_delete_user_token(c *C) {
	token, err := s.createToken()
	c.Assert(err, IsNil)

	if err := s.service.DeleteUserToken(token.UserID, token.Token); err != nil {
		c.Fatal(err)
	}

	_, err = s.service.GetUserToken(token.Token)
	c.Assert(err, Equals, sql.ErrNoRows)
}

// helpers

func (s *TestTokensSuite) createToken() (*schema.UserToken, error) {
	user, err := s.createUser()
	if err != nil {
		return nil, err
	}

	return s.service.CreateUserToken(user)
}

func (s *TestTokensSuite) createUser() (*schema.User, error) {
	user := &schema.User{
		FirstName:  fmt.Sprintf("name %d", s.usersCount),
		LastName:   fmt.Sprintf("last name %d", s.usersCount),
		FacebookID: fmt.Sprintf("facebook id %d", s.usersCount),
	}

	if err := s.service.db.Insert(user); err != nil {
		return nil, err
	}

	s.usersCount++
	return user, nil
}
