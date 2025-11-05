package repositorio

import (
	//repoCanciones "servidor.local/grpc-servidorCancion/dominio/cancion/repositorio"
	dto "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
	canciones "servidor.local/grpc-servidorCancion/dominio/cancion/modelo"
	repoGenero "servidor.local/grpc-servidorCancion/dominio/genero/repositorio"
)

// (El nextID y VectorCanciones se mantienen en repositorioCancionesAlmacenadas.go)

// CargarCanciones inicializa el slice con datos de ejemplo
func CargarCanciones() {
	var objCancion1, objCancion2, objCancion3, objCancion4, objCancion5, objCancion6, objCancion7, objCancion8 canciones.Cancion

	// 1
	objCancion1.Id = nextID
	nextID++
	objCancion1.Titulo = "The Fate of Ophelia"
	objCancion1.Artista = "Taylor Swift"
	objCancion1.Duracion = "3:46"
	objCancion1.AnioLanzamiento = 2025
	objCancion1.Idioma = "Inglés"
	objCancion1.Genero = repoGenero.BuscarGeneroNombre("Pop").ObjGenero

	// 2
	objCancion2.Id = nextID
	nextID++
	objCancion2.Titulo = "OTRA NOCHE"
	objCancion2.Artista = "Los Ángeles Azules ft. Nicki Nicole"
	objCancion2.Duracion = "3:32"
	objCancion2.AnioLanzamiento = 2022
	objCancion2.Idioma = "Español"
	objCancion2.Genero = repoGenero.BuscarGeneroNombre("Cumbia").ObjGenero

	// 3
	objCancion3.Id = nextID
	nextID++
	objCancion3.Titulo = "Wind of Change"
	objCancion3.Artista = "Scorpions"
	objCancion3.Duracion = "5:10"
	objCancion3.AnioLanzamiento = 1991
	objCancion3.Idioma = "Inglés"
	objCancion3.Genero = repoGenero.BuscarGeneroNombre("Rock").ObjGenero

	// 4
	objCancion4.Id = nextID
	nextID++
	objCancion4.Titulo = "Oye mi amor"
	objCancion4.Artista = "Maná"
	objCancion4.Duracion = "4:27"
	objCancion4.AnioLanzamiento = 1992
	objCancion4.Idioma = "Español"
	objCancion4.Genero = repoGenero.BuscarGeneroNombre("Rock").ObjGenero

	// 5
	objCancion5.Id = nextID
	nextID++
	objCancion5.Titulo = "Hit The Road Jack"
	objCancion5.Artista = "Ray Charles"
	objCancion5.Duracion = "2:01"
	objCancion5.AnioLanzamiento = 1961
	objCancion5.Idioma = "Inglés"
	objCancion5.Genero = repoGenero.BuscarGeneroNombre("Jazz").ObjGenero

	// 6
	objCancion6.Id = nextID
	nextID++
	objCancion6.Titulo = "La vie en rose"
	objCancion6.Artista = "Louis Armstrong"
	objCancion6.Duracion = "3:26"
	objCancion6.AnioLanzamiento = 1950
	objCancion6.Idioma = "Francés"
	objCancion6.Genero = repoGenero.BuscarGeneroNombre("Jazz").ObjGenero

	// 7
	objCancion7.Id = nextID
	nextID++
	objCancion7.Titulo = "Gotas de lluvia"
	objCancion7.Artista = "Grupo Niche"
	objCancion7.Duracion = "4:50"
	objCancion7.AnioLanzamiento = 1990
	objCancion7.Idioma = "Español"
	objCancion7.Genero = repoGenero.BuscarGeneroNombre("Salsa").ObjGenero

	// 8 (mantener Lamento Boliviano)
	objCancion8.Id = nextID
	nextID++
	objCancion8.Titulo = "Lamento Boliviano"
	objCancion8.Artista = "Enanitos verdes"
	objCancion8.Duracion = "3:45"
	objCancion8.AnioLanzamiento = 1994
	objCancion8.Genero = repoGenero.BuscarGeneroNombre("Rock").ObjGenero

	VectorCanciones = append(VectorCanciones, objCancion1, objCancion2, objCancion3, objCancion4, objCancion5, objCancion6, objCancion7, objCancion8)
}

// BuscarCancion busca una canción por título
func BuscarCancion(titulo string) dto.RespuestaDTO {
	var respuesta dto.RespuestaDTO
	for i := 0; i < len(VectorCanciones); i++ {
		if VectorCanciones[i].Titulo == titulo {
			respuesta.ObjCancion = VectorCanciones[i]
			respuesta.Codigo = 200
			respuesta.Mensaje = "Cancion encontrada correctamente"
			return respuesta
		}
	}
	respuesta.Codigo = 404
	respuesta.Mensaje = "La cancion no se encontro"
	return respuesta
}
