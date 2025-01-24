package testhelpers

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorNevola/config"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
)

type (
	PostgresContainer struct {
		*postgres.PostgresContainer
		ConnectionString string
	}
)

func createPostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.Run(
		ctx,
		"postgres:13",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(60*time.Second),
		),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

func runMigrations(db *bun.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "/home/nevola/personal/fidelis/scripts/db/migrations",
	}

	migrate.SetTable("migrations")
	migrate.SetIgnoreUnknown(true)
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Printf("Applied %d migrations!\n", n)

	return nil
}

func ConnectionToDB(ctx context.Context) (*bun.DB, func(), error) {
	postgresContainer, err := createPostgresContainer(ctx)
	if err != nil {
		panic(err)
	}

	db := config.NewDB(&config.DBConfigParams{
		Dsn: postgresContainer.ConnectionString,
	})

	if err := runMigrations(db); err != nil {
		return nil, nil, err
	}

	return db, func() {
		postgresContainer.Terminate(ctx)
	}, nil
}
