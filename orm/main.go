package main

import (
	"flag"
	"fmt"
	"github.com/ONSdigital/dp-dd-golang-db-spike/orm/handlers"
	"github.com/go-pg/pg"
	"github.com/gorilla/pat"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	dbUser := flag.String("db_user", "", "Database user name")
	dbPwd := flag.String("db_pwd", "", "Database password")
	dbName := flag.String("db_name", "", "Database name")
	flag.Parse()

	db := pg.Connect(&pg.Options{
		User:     *dbUser,
		Password: *dbPwd,
		Database: *dbName,
	})

	router := pat.New()
	router.Get("/dataResource/{dataResourceID}", handlers.GetDataResource(db))
	router.Get("/dataset/{dataSetID}", handlers.GetDataSet(db))

	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, os.Interrupt)

	go func() {
		for {
			s := <-sig_chan
			switch s {
			default:
				fmt.Println("\n> Attempting to close DB connection.")
				err := db.Close()
				if err != nil {
					fmt.Println("> Failed to close DB connection cleanly.")
					os.Exit(1)
				} else {
					fmt.Println("> DB connection closed cleanly.")
					os.Exit(0)
				}
			}
		}
	}()

	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
