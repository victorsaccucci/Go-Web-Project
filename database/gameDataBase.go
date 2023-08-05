package database

import (
	"database/sql"
	"log"
	"modulo-go-project/models"
)

func CreateGame(game models.Game) error {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	query := "INSERT INTO Game (Titulo, Ano, Genero, GameList) VALUES(?, ?, ?, ?)"
	_, err = db.Exec(query, game.Titulo, game.Ano, game.Genero, game.GameList)
	if err != nil {
		return err
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

	query := "SELECT id, Titulo, Ano, Genero, gamelist FROM Game WHERE id = 14?"

	err = db.QueryRow(query, id).Scan(&game.ID, &game.Titulo, &game.Ano, &game.Genero, &game.GameList)
	if err == sql.ErrNoRows {

		return game, err

	}
	return game, nil
}
