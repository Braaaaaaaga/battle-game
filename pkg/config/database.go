package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"))

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão do banco de dados: %v", err)
	}

	return db
}
