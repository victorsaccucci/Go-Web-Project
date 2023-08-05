package models

type Game struct {
	ID       int    `json:"id"`
	Titulo   string `json:"titulo"`
	Ano      int    `json:"ano"`
	Genero   string `json:"genero"`
	GameList int    `json:"gamelist"`
}
