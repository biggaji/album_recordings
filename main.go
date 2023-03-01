package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	configs := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error

	dbb, err := sql.Open("mysql", configs.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := dbb.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database Connected")

	// fetch albums for artist
	artistName := "John Coltrane"
	albums, err := fetchAlbumByArtist(artistName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(albums)
	// albLenght := len(albums)
	// var albPlural string

	// if albLenght > 1 {
	// 	albPlural = "albums"
	// } else {
	// 	albPlural = "album"
	// }

	// fmt.Printf("%v %v found for %v:", albLenght, albPlural, artistName)
}

func fetchAlbumByArtist(artist string) ([]Album, error) {
	// var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)

	if err != nil {
		return []Album{}, fmt.Errorf("fetchAlbumByArtist %q: %v", artist, err)
	}
	fmt.Print(rows)

	// defer rows.Close()

	// for rows.Next() {
	// 	var alb Album
	// 	fmt.Print(alb)
	// 	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	// 		return nil, fmt.Errorf("fetchAlbumByArtist %q: %v", artist, err)
	// 	}

	// 	albums = append(albums, alb)
	// }

	// if err := rows.Err(); err != nil {
	// 	return nil, fmt.Errorf("fetchAlbumByArtist %q: %v", artist, err)
	// }

	return []Album{}, nil
}
