package models

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"path/filepath"
)

var DB *pgxpool.Pool

type sql struct {
	getAllAlbums string
	insertAlbum string
}
/**
Returns the contents of a file
 */
func ReadSql(folder string, filename string) string {

	absPath, err := filepath.Abs("../../pkg/models/sql/")
	fmt.Println(absPath)
	if err != nil{
		log.Fatalf("could not read sql folder path %v\n", absPath)
	}
	path := filepath.Join(absPath, folder, filename)
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Println(ioErr)
		log.Fatalf("could not read sql file %v\n", path)
	}
	return string(c)
}

var SQL = &sql{
	getAllAlbums: ReadSql("select","getAllAlbums.sql"),
	insertAlbum: ReadSql("insert", "insertAlbum.sql"),
}