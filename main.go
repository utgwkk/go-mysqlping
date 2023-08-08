package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	dsn := buildDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, "SELECT 1"); err != nil {
		log.Fatal(err)
	}
	log.Println("Ping successful")
}

func buildDSN() string {
	cfg := mysql.NewConfig()
	// TODO: give options from CLI args
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Passwd = "root"
	return cfg.FormatDSN()
}
