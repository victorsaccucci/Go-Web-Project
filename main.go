package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/gowebprojectdb")

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	insert, err := db.Query("INSERT INTO Game (Titulo, Ano, Genero, GameList) VALUES ('The Witcher 3', 2015, 'RPG', 4);")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successful Connection to Database!")
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"

// 	"database/sql"
// 	"modulo-go-project/database"
// )

// func main() {

// 	// Abrir conexão com o banco de dados
// 	db, err := database.OpenDB()
// 	if err != nil {
// 		log.Fatal("Erro ao estabelecer a conexão com o banco de dados:", err)
// 	}
// 	defer database.CloseDB(db) // Fechar a conexão no final do programa

// 	// Verificar se a conexão foi bem sucedida
// 	if err = db.Ping(); err != nil {
// 		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
// 	}

// 	// Resto do seu código aqui...

// 	port := ":8080"
// 	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
// 	log.Fatal(http.ListenAndServe(port, nil))
// }

// func handleGetGameByID(w http.ResponseWriter, r *http.Request) {
// 	varHttp := mux.Vars(r)
// 	idGame := varHttp["id"]

// 	id, err := strconv.Atoi(idGame)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	game, err := database.GetGameByID(id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.Error(w, "Game not found", http.StatusNotFound)
// 		} else {
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(game)
// }
