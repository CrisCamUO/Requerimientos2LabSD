package repositorio

import (
	repoCanciones "servidor.local/grpc-servidorCancion/dominio/cancion/capaAccesoADatos"
	dto "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
	canciones "servidor.local/grpc-servidorCancion/dominio/cancion/modelo"
	repoGenero "servidor.local/grpc-servidorCancion/dominio/genero/repositorio"
)

// nextID se usa para asignar IDs únicos auto-incrementales dentro del proceso
var nextID int32 = 1

// CargarCanciones inicializa el slice con datos de ejemplo
func CargarCanciones() {
	var objCancion1, objCancion2, objCancion3, objCancion4 canciones.Cancion

	objCancion1.Id = nextID
	nextID++
	objCancion1.Titulo = "La Vida"
	objCancion1.Artista = "Carlos Vives"
	objCancion1.Duracion = "3:45"
	objCancion1.AnioLanzamiento = 1998
	objCancion1.Idioma = "Español"
	objCancion1.Genero = repoGenero.BuscarGeneroNombre("Salsa").ObjGenero

	objCancion2.Id = nextID
	nextID++
	objCancion2.Titulo = "La Bicicleta"
	objCancion2.Artista = "Shakira"
	objCancion2.Duracion = "3:38"
	objCancion2.AnioLanzamiento = 2016
	objCancion2.Genero = repoGenero.BuscarGeneroNombre("Pop").ObjGenero

	objCancion3.Id = nextID
	nextID++
	objCancion3.Titulo = "Ojos Así"
	objCancion3.Artista = "Shakira"
	objCancion3.Duracion = "4:12"
	objCancion3.AnioLanzamiento = 2000
	objCancion3.Genero = repoGenero.BuscarGeneroNombre("Pop").ObjGenero

	objCancion4.Id = nextID
	nextID++
	objCancion4.Titulo = "LamentoBoliviano"
	objCancion4.Artista = "Los enanitos verdes"
	objCancion4.Duracion = "4:12"
	objCancion4.AnioLanzamiento = 2000
	objCancion4.Genero = repoGenero.BuscarGeneroNombre("Rock").ObjGenero

	repoCanciones.VectorCanciones = append(repoCanciones.VectorCanciones, objCancion1, objCancion2, objCancion3, objCancion4)
}

// BuscarCancion busca una canción por título
func BuscarCancion(titulo string) dto.RespuestaDTO {
	var respuesta dto.RespuestaDTO
	for i := 0; i < len(repoCanciones.VectorCanciones); i++ {
		if repoCanciones.VectorCanciones[i].Titulo == titulo {
			respuesta.ObjCancion = repoCanciones.VectorCanciones[i]
			respuesta.Codigo = 200
			respuesta.Mensaje = "Cancion encontrada correctamente"
			return respuesta
		}
	}
	respuesta.Codigo = 404
	respuesta.Mensaje = "La cancion no se encontro"
	return respuesta
}
