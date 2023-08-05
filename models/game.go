package models

type Game struct {
	Idgame   int    `json:"idgame"`
	Titulo   string `json:"titulo"`
	Ano      int    `json:"ano"`
	Genero   string `json:"genero"`
	Gamelist int    `json:"gamelist"`
}
