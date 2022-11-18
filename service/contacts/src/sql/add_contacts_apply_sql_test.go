package sql

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Exponential backoff
// const sleepTime = (2 ** retryCount) * 100
//                    + Math.ceil(Math.random() * 100);
//                 await sleep(sleepTime);

func TestSql(t *testing.T) {

	// config, err := pgxpool.ParseConfig("postgresql://luojm:ccccc123@crdb.env0.luojm.com:9080/contacts?sslmode=verify-ca&sslrootcert=/etc/app-0/cert/ca.crt&pool_max_conns=40")
	ctx := context.Background()
	config, err := pgxpool.ParseConfig("postgresql://luojm:ccccc123@crdb.env0.luojm.com:9080/contacts?sslmode=verify-ca&sslrootcert=/etc/app-0/cert/ca.crt")
	config.MaxConns = 4
	config.MaxConnIdleTime = 3 * time.Second
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	dbpool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer dbpool.Close()

	connection, err := dbpool.Acquire(ctx)
	defer connection.Release()
	if err != nil {
		log.Fatal("dbpool acquire fail: ", err)
	}

	// updateTime := time.Now().Format(time.RFC3339)
	// result, err := connection.Exec(ctx, UpsertAddContactsApply, "3", "2", "please  me", updateTime)
	// if err != nil{
	// 	log.Fatal("exec: ", UpsertAddContactsApply, " fail, err: ",  err)
	// }
	// fmt.Println(result)
	//bulk insert
	//https://github.com/jackc/pgx/issues/764#issuecomment-685249471 

	rows, err := connection.Query(ctx, AddContactsApply, "2")
	if err != nil {
		log.Fatal("query: ", AddContactsApply, " fail, err: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user_id string
		var message string
		var update_time time.Time
	
		err = rows.Scan(&user_id, &message, &update_time)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}
		fmt.Printf("user: %s, send: %s in %s\n", user_id, message, update_time)
	}
	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return
	}
}
