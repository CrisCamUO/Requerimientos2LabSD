package fachadacancionesservices

import (
	"fmt"

	capaaccesoadatos "servidor.local/grpc-servidorCancion/dominio/cancion/capaAccesoADatos"
	componnteconexioncola "servidor.local/grpc-servidorCancion/dominio/cancion/componnteConexionCola"
	dtos "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccesoadatos.RepositorioCanciones
	conexionCola *componnteconexioncola.RabbitPublisher
}

// Constructor de la fachada
func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println("🔧 Inicializando fachada de almacenamiento...")

	repo := capaaccesoadatos.GetRepositorioCanciones()

	conexionCola, err := componnteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println(" Error al conectar con RabbitMQ:", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenarDTOInput, data []byte) error {
	thisF.conexionCola.PublicarNotificacion(componnteconexioncola.NotificacionCancion{
		Titulo:          objCancion.Titulo,
		Genero:          objCancion.Genero,
		Artista:         objCancion.Artista,
		Idioma:          objCancion.Idioma,
		AnioLanzamiento: objCancion.AnioLanzamiento,
		Duracion:        objCancion.Duracion,
		Mensaje:         "Nueva cancion almacenada: " + objCancion.Titulo + " de " + objCancion.Artista,
	})
	// Guardar archivo y registro en memoria
	// delegar en el repositorio

	return thisF.repo.GuardarCancion(objCancion.Titulo, objCancion.Genero, objCancion.Artista, objCancion.Idioma, objCancion.AnioLanzamiento, objCancion.Duracion, data)
}
