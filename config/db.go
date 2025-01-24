package config

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type (
	DBConfigParams struct {
		Dsn string
	}

	DBConfig interface {
		NewDB() *bun.DB
	}
)

func NewDB(DBConfig *DBConfigParams) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DBConfig.Dsn)))

	if err := sqldb.Ping(); err != nil {
		log.Debug("DB Connection failed", "error", err)
	} else {
		log.Debug("DB Connection successful")
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
