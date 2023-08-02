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

	query := "INSERT INTO Game (Titulo, Ano, Genero, gamelist) VALUES(?, ?, ?, ?)"
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
	query := "SELECT game.idgame, game.titulo, game.ano, game.genero, game.gamelist FROM Game WHERE idgame = ?"
	err = db.QueryRow(query, id).Scan(&game.ID, &game.Titulo, &game.Ano, &game.Genero, &game.GameList)
	if err == sql.ErrNoRows {

		return game, err

	}
	return game, nil
}
