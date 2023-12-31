package main

import (
	"encoding/json"
	"fmt"
	"log"
	"modulo-go-project/database"
	"modulo-go-project/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	//Game
	router.HandleFunc("/game/{id}", handleGetGameByID).Methods("GET")
	router.HandleFunc("/games", handleAllGames).Methods("GET")
	router.HandleFunc("/insert", handleCreateGame).Methods("POST")
	router.HandleFunc("/delete/{id}", handleDeleteGame).Methods(http.MethodDelete)

	router.HandleFunc("/create", handleNewGame).Methods(http.MethodPost)

	//Gamelist(GameList GET)
	router.HandleFunc("/description/gameId/{id}", handleGetDescriptionByIDgame).Methods("GET")
	router.HandleFunc("/gamelist/{id}", handleGetGamelistByID).Methods("GET")

	port := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}

//Game

func handleNewGame(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao analisar o formulário", http.StatusInternalServerError)
			return
		}

		titulo := r.FormValue("titulo")
		anoStr := r.FormValue("ano")
		genero := r.FormValue("genero")
		gamelistStr := r.FormValue("gamelist")
		ano, err := strconv.Atoi(anoStr)

		if err != nil {
			http.Error(w, "Ano inválido", http.StatusBadRequest)
			return
		}

		gamelist, err := strconv.Atoi(gamelistStr)
		if err != nil {
			http.Error(w, "Gamelist inválido", http.StatusBadRequest)
			return
		}

		newGame := models.Game{
			Titulo:   titulo,
			Ano:      ano,
			Genero:   genero,
			Gamelist: gamelist,
		}

		err = database.CreateGame(newGame)
		if err != nil {
			http.Error(w, "Erro ao inserir o jogo no banco de dados", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Jogo cadastrado com sucesso"))

	} else {

		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
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

func handleAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := database.GetAllGames()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)

}

func handleCreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.CreateGame(game)
	if err != nil {
		http.Error(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Game created successfully"))
}

func handleDeleteGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		idGame := vars["id"]

		id, err := strconv.Atoi(idGame)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		err = database.DeleteGame(id)
		if err != nil {
			http.Error(w, "Failed to delete game", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Game deleted succesfully"))
	} else {
		http.Error(w, "Method not suported", http.StatusMethodNotAllowed)
	}
}

//GameList

func handleGetGamelistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idGameList := vars["id"]

	id, err := strconv.Atoi(idGameList)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	gamelist, err := database.GetGameListByID(id)
	if err != nil {
		if err == database.ErrNoRows {
			http.Error(w, "Game list not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gamelist)

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameList)
}
