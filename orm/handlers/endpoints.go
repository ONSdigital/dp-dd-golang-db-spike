package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-golang-db-spike/orm/models"
	"github.com/go-pg/pg"
	"net/http"
	"os"
)

func GetDataResource(db *pg.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ds := r.URL.Query().Get(":dataResourceID")
		dr := &models.DataResource{DataResource: ds}

		if err := db.Select(dr); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(dr)
	}
}

func GetDataSet(db *pg.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get(":dataSetID")
		dataSet := &models.DataSet{Id: id}

		if err := db.Select(dataSet); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(dataSet)
	}
}
