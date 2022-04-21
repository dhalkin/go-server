package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type documents struct {
	DocumentList []Document
}

type Document struct {
	ID       int
	UniqueID string
	Filename string
}

func GetContainer(c echo.Context) error {

	db, err := sql.Open("mysql", "signnow:ba99wu8ck2@/signnow")
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT id, unique_id, filename FROM documents limit 10`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	documents := documents{}
	for rows.Next() {
		document := Document{}
		err = rows.Scan(&document.ID, &document.UniqueID, &document.Filename)
		if err != nil {
			panic(err)
		}
		documents.DocumentList = append(documents.DocumentList, document)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, documents)
}
