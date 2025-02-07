package main

import (
	"battle-game/internal/api/handlers"
	"battle-game/pkg/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"))

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao verificar a conexão do banco de dados:", err)
	}

	handlers.SetDatabase(db)

	router := mux.NewRouter()
	router.HandleFunc("/user/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/{userID}/characters", handlers.AddCharacters).Methods("POST")

	router.HandleFunc("/user/battle", handlers.BattleHandler).Methods("POST")

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
