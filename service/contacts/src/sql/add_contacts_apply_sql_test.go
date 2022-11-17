package sql

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestSql(t *testing.T) {

	config, err := pgxpool.ParseConfig("postgresql://luojm:ccccc123@crdb.env0.luojm.com:9080/contacts?sslmode=verify-ca&sslrootcert=/etc/app-0/cert/ca.crt&pool_max_conns=40")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer dbpool.Close()


	// dsn := "postgresql://luojm:ccccc123@crdb.env0.luojm.com:9080/defaultdb?sslmode=verify-ca&sslrootcert=../../cert/ca.crt"
	// ctx := context.Background()
	// conn, err := pgx.Connect(ctx, dsn)
	// defer func() { conn.Close(context.Background()) }()
	// if err != nil {
	// 	log.Fatal("failed to connect database", err)
	// }

	var text string
	err = dbpool.QueryRow(context.Background(),
		"SELECT CONCAT('Hello from CockroachDB at ',CAST (NOW() as STRING))").Scan(&text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(text)
}
