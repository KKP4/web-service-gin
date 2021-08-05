package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func initDB(addr string) *pgxpool.Pool {
	db, err := pgxpool.Connect(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connected\n")
	}
	return db
}
