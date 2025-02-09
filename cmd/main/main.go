package main

import (
	"battle-game/internal/api/controller"
	"battle-game/internal/router"
	"battle-game/pkg/config"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfig()

	db := config.InitDB()
	defer db.Close()

	controller.SetDatabase(db)
	router := router.New(db)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
