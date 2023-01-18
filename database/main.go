package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID string
	Title string
	Artist string
	Price float32
}

func getAlbums() ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album")
	if err != nil {
		return albums, fmt.Errorf("getAlbums: %v", err)
	}

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return albums, fmt.Errorf("getAlbums: %v", err)
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func main() {
	config := mysql.Config{
		User: "root",
		Passwd: "",
		DBName: "godb",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	albums, err := getAlbums()
	if err != nil {
		panic(err.Error())
	}

	for _, al := range albums {
		fmt.Println(al.Title)
	}
}
