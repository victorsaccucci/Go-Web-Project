package database

import (
	"database/sql"
	"log"
	"modulo-go-project/models"
)

func CreateGameList(gameList models.GameList) error {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)
	query := "INSERT INTO gamelist (descricao) VALUES(?)"
	_, err = db.Exec(query, gameList.Descricao)
	if err != nil {
		return err
	}
	return nil
}

func GetGameListbyID(id int) (models.GameList, error) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	var gamelist models.GameList

	query := "SELECT gamelist.id, gamelist.descricao FROM gamelist WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&gamelist.ID, &gamelist.Descricao)
	if err == sql.ErrNoRows {

		return gamelist, err

	}
	return gamelist, nil
}
