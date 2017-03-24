package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-golang-db-spike/basic/models/dataset"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetDataSets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Query(dataset.SELECT_ALL_QUERY)
		defer rows.Close()

		dataSets := make([]*dataset.DAO, 0)
		for rows.Next() {
			dataSets = append(dataSets, dataset.MapFromRow(rows))
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(dataSets)
	}
}

func SaveDataSet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stmt, _ := db.Prepare(dataset.INSERT_STMT)

		b, _ := ioutil.ReadAll(r.Body)
		s := string(b)
		fmt.Printf("\n%s\n", s)

		var d dataset.DAO
		err := json.NewDecoder(strings.NewReader(s)).Decode(&d)
		if err != nil {
			panic(err.Error())
		}
		defer r.Body.Close()

		_, err = stmt.Exec(d.Id, d.MajorLabel, d.MinorVersion, d.Metadata, d.MinorVersion,
			d.RevisionNotes, d.RevisionReason, d.S3URL, d.Status, d.Title, d.TotalRowCount)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
	}
}
