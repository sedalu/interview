/*
package pgsql implements a PostgreSQL store

tables:
  interview.challenges:
    columns:
      - id uuid primary key
      - name text not null unique
      - created_at timestamp without time zone not null
*/
package pgsql

import (
	"database/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

	"github.com/sedalu/interview/internal/store"
)

type Store struct {
	*sql.DB
}

func Open(dsn string) (*Store, error) {
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	return &Store{
		stdlib.OpenDB(*config),
	}, nil
}

var _ store.Store = (*Store)(nil)
