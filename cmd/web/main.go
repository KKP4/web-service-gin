package main

import (
	"example.com/web-service-gin/pkg/models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load(".env")
	addr := os.Getenv("DATABASE_URL")
	fmt.Println(addr)

	models.DB = initDB(addr)
	defer models.DB.Close()

	r := setupRouter()

	err = r.Run(os.Getenv("HOST") + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("could not start server")
	}
}
