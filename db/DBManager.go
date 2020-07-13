package db

import (
	"ToRead/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func FetchToReads() models.ToReads {
	db, err := sql.Open("mysql", "root:@to_read_db)")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	selectQuery, err := db.Query("SELECT title, referer FROM to_read_db.to_reads")

	if err != nil {
		panic(err)
	}

	toReads := models.ToReads{}

	for selectQuery.Next() {
		var title, referer string
		err = selectQuery.Scan(&title, &referer)

		if err != nil {
			panic(err)
		}

		toReads = append(toReads, models.ToRead{Title: title, Referrer: referer})
	}

	return toReads
}
