package main

import (
	"log"

	"github.com/joho/godotenv"

	"todoapp/internal/repository"
	"todoapp/internal/router"
)

func main() {
	_ = godotenv.Load()

	db, err := repository.NewDB()
	if err != nil {
		log.Fatalf("error while connect to db: %v", err)
	}
	defer db.Close()

	log.Println("connect to db is success")

	r := router.SetupRouter(db)
	r.Run(":8080")
}
