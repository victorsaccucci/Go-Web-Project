package database

import (
	"database/sql"
	"fmt"
	"log"
	"modulo-go-project/models"
)

func CreateGame(game models.Game) error {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	query := "INSERT INTO Game (Titulo, Ano, Genero, gamelist) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, game.Titulo, game.Ano, game.Genero, game.Gamelist)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("failed to insert a new game")
	}

	return nil
}

func GetGameByID(id int) (models.Game, error) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	var game models.Game

	query := "SELECT idgame, Titulo, Ano, Genero, gamelist FROM Game WHERE idgame = ?"

	err = db.QueryRow(query, id).Scan(&game.Idgame, &game.Titulo, &game.Ano, &game.Genero, &game.Gamelist)
	if err == sql.ErrNoRows {

		return game, err

	}
	return game, nil
}

func GetAllGames() ([]models.Game, error) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	var games []models.Game

	query := "SELECT idgame, Titulo, Ano, Genero, gamelist FROM Game"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.Idgame, &game.Titulo, &game.Ano, &game.Genero, &game.Gamelist)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}
