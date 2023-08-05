package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"modulo-go-project/database"
)

func main() {

	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB(db)

	// Testar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/game/{id}", handleGetGameByID)
	router.HandleFunc("/games", handleAllGames)
	router.HandleFunc("/description/gameId/{id}", handleGetDescriptionByIDgame)

	port := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
func handleAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := database.GetAllGames()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)

}

func handleGetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idGame := vars["id"]

	id, err := strconv.Atoi(idGame)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	game, err := database.GetGameByID(id)
	if err != nil {
		if err == database.ErrNoRows {
			http.Error(w, "Game not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
func handleGetDescriptionByIDgame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idGame := vars["id"]

	id, err := strconv.Atoi(idGame)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	gameList, err := database.GetDescriptionListByGameID(id)
	if err != nil {
		if err == database.ErrNoRows {
			http.Error(w, "Game not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json") // Corrigido: "Content-Type" em vez de "Content-Typer"
	json.NewEncoder(w).Encode(gameList)
}
