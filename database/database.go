package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context) (*postgres, error) {
	pgOnce.Do(func() {
		config, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s dbname=%s", os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE")))
		if err != nil {
			log.Fatalf("unable to parse PostgreSQL configuration: %v", err)
		}

		db, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			// Check if the error is a PgError
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				log.Fatalf("PostgreSQL error - Code: %s, Message: %s", pgErr.Code, pgErr.Message)
			} else {
				// If it's not a PgError, log the general error
				log.Fatalf("unable to create connection pool: %v", err)
			}
		}

		pgInstance = &postgres{db}
	})

	return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
