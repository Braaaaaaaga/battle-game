package router

import (
	"battle-game/internal/api/controller"
	"database/sql"

	"github.com/gorilla/mux"
)

func New(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.ListUsers).Methods("GET")
	router.HandleFunc("/user/battle", controller.BattleHandler).Methods("POST")
	router.HandleFunc("/user/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/users/{userID}/characters", controller.AddCharacters).Methods("POST")
	router.HandleFunc("/users/{userID}/characters/edit", controller.EditCharacter).Methods("POST")

	return router
}
