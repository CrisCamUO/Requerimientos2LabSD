package dto

type CancionAlmacenarDTOInput struct {
	Titulo          string `json:"titulo"`
	Genero          string `json:"genero"`
	Artista         string `json:"artista"`
	Idioma          string `json:"idioma"`
	AnioLanzamiento int32  `json:"anio_lanzamiento"`
	Duracion        string `json:"duracion"`
}
