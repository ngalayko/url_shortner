package migrate

import (
	"context"

	"go.uber.org/zap"
	reform "gopkg.in/reform.v1"

	"github.com/ngalayko/url_shortner/server/db"
	"github.com/ngalayko/url_shortner/server/logger"
)

const (
	ctxKey migrationsCtxKey = "migrations_ctx_key"
)

type migrationsCtxKey string

// Migrate is a service to apply db migrations
type Migrate struct {
	Db *db.Db

	logger logger.ILogger
}

// NewContext stores Migrate in context
func NewContext(ctx context.Context, migrations interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, ok := migrations.(*Migrate); !ok {
		migrations = newMigrate(ctx)
	}

	return context.WithValue(ctx, ctxKey, migrations)
}

// FromContext returns Migrate from context
func FromContext(ctx context.Context) *Migrate {
	if migrations, ok := ctx.Value(ctxKey).(*Migrate); ok {
		return migrations
	}

	return newMigrate(ctx)
}

func newMigrate(ctx context.Context) *Migrate {
	return &Migrate{
		Db:     db.FromContext(ctx),
		logger: logger.FromContext(ctx),
	}
}

// Apply applies all migrations that were not applied
func (m *Migrate) Apply() error {
	if err := m.applyInitMigrations(); err != nil {
		return err
	}

	appliedMigrations, err := m.applied()
	if err != nil {
		return err
	}

	m.logger.Info("Applying migrations...")

	var count int
	for _, migration := range migrations() {
		if _, ok := appliedMigrations[migration.Name]; ok {
			continue
		}

		if err = m.applyMigration(migration); err != nil {
			return err
		}

		count++
	}

	m.logger.Info("Appled migrations",
		zap.Int("count", count),
	)

	return nil
}

// Flush flushes migrations
func (m *Migrate) Flush() error {
	migrations := append(initMigrations(), migrations()...)
	for i := len(migrations) - 1; i >= 0; i-- {
		m.logger.Info("Flushing migration",
			zap.String("name", migrations[i].Name),
			zap.String("query", migrations[i].FlushSQL),
		)

		if len(migrations[i].FlushSQL) == 0 {
			continue
		}

		if _, err := m.Db.Exec(migrations[i].FlushSQL); err != nil {
			m.logger.Error("error flushing migration",
				zap.String("name", migrations[i].Name),
				zap.String("query", migrations[i].FlushSQL),
			)
		}
	}

	return nil
}

func (m *Migrate) applyInitMigrations() error {
	for _, migration := range initMigrations() {
		if _, err := m.Db.Exec(migration.RawSQL); err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrate) applied() (map[string]*migration, error) {
	rows, err := m.Db.SelectRows(migrationTable, "")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := map[string]*migration{}
	for {
		migration := &migration{}
		if err := m.Db.NextRow(migration, rows); err != nil {
			break
		}

		result[migration.Name] = migration
	}

	return result, nil
}

func (m *Migrate) applyMigration(migration *migration) error {
	return m.Db.InTransaction(func(tx *reform.TX) error {

		m.logger.Info("applying migration",
			zap.String("name", migration.Name),
			zap.String("query", migration.RawSQL),
		)

		if _, err := tx.Exec(migration.RawSQL); err != nil {
			return err
		}

		if _, err := tx.Exec(`INSERT INTO migrations (name) VALUES ($1)`, migration.Name); err != nil {
			return err
		}

		m.logger.Info("migration applied",
			zap.String("name", migration.Name),
		)

		return nil
	})
}
