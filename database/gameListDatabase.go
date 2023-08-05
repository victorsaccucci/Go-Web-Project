package database

import (
	"database/sql"
	"log"
	"modulo-go-project/models"
)

func GetGameListByID(id int) (models.GameList, error) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	var gamelist models.GameList

	query := "SELECT gamelist.id, gamelist.descricao from gamelist where id = ?"
	err = db.QueryRow(query, id).Scan(&gamelist.Id, &gamelist.Descricao)

	if err == sql.ErrNoRows {
		return gamelist, err
	}
	return gamelist, nil
}

func GetDescriptionListByGameID(id int) (models.GameList, error) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	var gamelist models.GameList

	query := `SELECT gamelist.id, gamelist.descricao
	FROM game
	INNER JOIN gamelist ON
	game.gamelist = gamelist.id
	WHERE game.idgame = ?;`

	err = db.QueryRow(query, id).Scan(&gamelist.Id, &gamelist.Descricao)

	if err == sql.ErrNoRows {
		return gamelist, err
	}
	return gamelist, nil
}
