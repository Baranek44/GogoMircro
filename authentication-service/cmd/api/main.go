package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Authentication service has been started")

	// Try to connect to DataBase
	conn := connectToDB()
	if conn == nil {
		log.Panic("Failed connecting to PG")
	}

	// Set up all config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		// Addr string need to be empty
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

//Make sure its available before make connection to DB
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Waiting for postgres...")
			count++
		} else {
			log.Println("Successed! Connected to PG")
			return connection
		}

		if count > 8 {
			log.Println(err)
			return nil
		}

		log.Println("Refreshing for 4s")
		time.Sleep(4 * time.Second)
		continue
	}
}
