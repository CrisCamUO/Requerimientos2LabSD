package repositorio

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	canciones "servidor.local/grpc-servidorCancion/dominio/cancion/modelo"

	dto "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
	repoGenero "servidor.local/grpc-servidorCancion/dominio/genero/repositorio"
)

// VectorCanciones ahora es un slice dinámico en vez de un arreglo fijo
var VectorCanciones []canciones.Cancion

// nextID se usa para asignar IDs únicos auto-incrementales dentro del proceso
var nextID int32 = 1

type RepositorioCanciones struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioCanciones
	once      sync.Once
)

// GetRepositorioCanciones aplica patrón Singleton
func GetRepositorioCanciones() *RepositorioCanciones {
	once.Do(func() {
		instancia = &RepositorioCanciones{}
	})
	return instancia
}

func (r *RepositorioCanciones) GuardarCancion(titulo string, genero string, artista string, idioma string, anioLanzamiento int32, duracion string, data []byte) error {
	r.GuardarAudioCancion(titulo, genero, artista, data)
	var objCancion canciones.Cancion
	objCancion.Id = 0
	objCancion.Titulo = titulo
	objCancion.Genero = repoGenero.BuscarGeneroNombre(genero).ObjGenero
	objCancion.Artista = artista
	objCancion.Idioma = idioma
	objCancion.AnioLanzamiento = anioLanzamiento
	objCancion.Duracion = duracion
	GuardarMetadatosCancionancion(objCancion)
	return nil
}

func (r *RepositorioCanciones) GuardarAudioCancion(titulo string, genero string, artista string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	//crear carpeta si no existe
	os.MkdirAll("Audios", os.ModePerm)

	//construir nombre de archivo
	fileName := fmt.Sprintf("%s_%s_%s.mp3", titulo, genero, artista)
	filePath := filepath.Join("Audios", fileName)

	//Guardar archivo fisico
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar archivo: %v", err)
	}

	//crear registro en memoria
	return nil
}
func GuardarMetadatosCancionancion(nuevaCancion canciones.Cancion) dto.RespuestaDTO {
	var respuesta dto.RespuestaDTO
	// Asignar ID automático si no viene uno
	if nuevaCancion.Id == 0 {
		nuevaCancion.Id = nextID
		nextID++
	}

	VectorCanciones = append(VectorCanciones, nuevaCancion)
	respuesta.Codigo = 201
	respuesta.Mensaje = "Cancion agregada correctamente"
	return respuesta
}
