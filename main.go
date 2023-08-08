package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
)

var version string = "(unset)"

type Opts struct {
	User     string `short:"u" long:"user" description:"MySQL user name" value-name:"user_name" default:"root"`
	Password string `short:"p" long:"password" description:"MySQL user password" value-name:"password"`
	Host     string `short:"h" long:"host" description:"Host to connect to the MySQL server" value-name:"host_name" default:"127.0.0.1"`
	Port     uint   `short:"P" long:"port" description:"Port used for the connection" value-name:"port_num" default:"3306"`
	Help     bool   `long:"help" description:"Show this help"`
	Version  bool   `long:"version" description:"Show this version"`
}

func main() {
	var opts Opts
	parser := flags.NewParser(&opts, flags.None)
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if opts.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	dsn := buildDSN(opts)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, "SELECT 1"); err != nil {
		log.Fatal(err)
	}
	log.Println("Ping successful")
}

func buildDSN(opts Opts) string {
	cfg := mysql.NewConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	cfg.Net = "tcp"
	cfg.User = opts.User
	cfg.Passwd = opts.Password
	return cfg.FormatDSN()
}
