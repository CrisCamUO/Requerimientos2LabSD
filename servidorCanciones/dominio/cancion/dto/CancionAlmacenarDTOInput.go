package dto

type CancionAlmacenarDTOInput struct {
	Titulo  string `json:"titulo"`
	Genero  string `json:"genero"`
	Artista string `json:"artista"`
	Idioma 	string `json:"idioma"`
}
