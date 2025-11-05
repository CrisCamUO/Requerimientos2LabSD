package modelo

import (
	"servidor.local/grpc-servidorCancion/dominio/genero/modelo"
)

type Cancion struct {
	Id              int32
	Titulo          string
	Artista         string
	AnioLanzamiento int32
	Duracion        string
	Genero          modelo.Genero
	Idioma          string
}
