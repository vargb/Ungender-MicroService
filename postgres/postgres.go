package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	logrus.Info(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
