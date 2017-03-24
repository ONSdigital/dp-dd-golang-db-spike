package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ONSdigital/dp-dd-golang-db-spike/basic/handler"
	"github.com/gorilla/pat"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
)

const db_connection_fmt = "user=%s password=%s dbname=%s sslmode=disable"

func main() {
	dbUsr := flag.String("db_user", "", "Database username.")
	dbPwd := flag.String("db_pwd", "", "Database password.")
	dbName := flag.String("db_name", "", "Database name.")
	flag.Parse()

	db, err := sql.Open("postgres", fmt.Sprintf(db_connection_fmt, *dbUsr, *dbPwd, *dbName))
	if err != nil {
		fmt.Println("Error connecting to the DB: " + err.Error())
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Ping failed. " + err.Error())
		os.Exit(1)
	}

	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, os.Interrupt)

	go func() {
		for {
			s := <-sig_chan
			switch s {
			default:
				fmt.Println("\nAttempting to close DB connection.")
				err := db.Close()
				if err != nil {
					fmt.Println("Failed to close DB connection cleanly.")
					os.Exit(1)
				} else {
					fmt.Println("DB connection closed cleanly.")
					os.Exit(0)
				}
			}
		}
	}()

	router := pat.New()
	router.Get("/datasets", handler.GetDataSets(db))
	router.Post("/dataset", handler.SaveDataSet(db))

	fmt.Println("\tGET -> http:localhost:8000/datasets")
	fmt.Println("\tPOST Dataset request BODY -> http:localhost:8000/dataset ")

	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println("Failed to start.")
		os.Exit(1)
	}
}
