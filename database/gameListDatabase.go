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
	query := "INSERT INTO Gamelist (Descricao) VALUES(?)"
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

	query := "SELECT Gamelist.ID, Gamelist.Descricao FROM Gamelist WHERE ID = ?"
	err = db.QueryRow(query, id).Scan(&gamelist.ID, &gamelist.Descricao)
	if err == sql.ErrNoRows {

		return gamelist, err

	}
	return gamelist, nil
}
