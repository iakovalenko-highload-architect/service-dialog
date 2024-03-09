package cmd

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	mustInitEnv()
}

func mustInitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func MustInitPostgresql() *sqlx.DB {
	conn, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_SSLMODE"),
		),
	)
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}

	return conn
}
