package models

import (
	"context"
	"fmt"
)

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}


func GetAllAlbums()(albums []Album, err error) {

	rows, err := DB.Query(context.Background(), SQL.getAllAlbums)
	if err != nil {
		return nil, err

	}
	if rows != nil {
		for rows.Next(){
			var album Album
			err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price)

			if err != nil {
				fmt.Println(err)

				return nil, err
			}
			albums = append(albums, album)

		}
	}
	return albums, nil
}

func InsertAlbum(title string, artist string, price float64)(err error) {
	_, err = DB.Exec(context.Background(), SQL.insertAlbum, title, artist, price)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}