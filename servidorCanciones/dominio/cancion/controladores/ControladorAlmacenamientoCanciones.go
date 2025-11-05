package controladores

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	dto "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
	fachadacancionesservices "servidor.local/grpc-servidorCancion/dominio/cancion/fachadaCancionesServices"
)

type ControladorAlmacenamientoCanciones struct {
	fachada *fachadacancionesservices.FachadaAlmacenamiento
}

// Constructor del Controlador
func NuevoControladorAlmacenamientoCanciones() *ControladorAlmacenamientoCanciones {
	return &ControladorAlmacenamientoCanciones{
		fachada: fachadacancionesservices.NuevaFachadaAlmacenamiento(),
	}
}

func (thisC *ControladorAlmacenamientoCanciones) AlmacenarCancion(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Almacenando canción...\n")
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error leyendo el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// leer y convertir anio_lanzamiento a int32
	var anio int32 = 0
	if anioStr := r.FormValue("anio_lanzamiento"); anioStr != "" {
		v, err := strconv.ParseInt(anioStr, 10, 32)
		if err != nil {
			http.Error(w, "anio_lanzamiento inválido", http.StatusBadRequest)
			return
		}
		anio = int32(v)
	}

	// leer los campos del DTO
	dtoInput := dto.CancionAlmacenarDTOInput{
		Titulo:          r.FormValue("titulo"),
		Genero:          r.FormValue("genero"),
		Artista:         r.FormValue("artista"),
		Idioma:          r.FormValue("idioma"),
		Duracion:        r.FormValue("duracion"),
		AnioLanzamiento: anio,
	}

	if err := thisC.fachada.GuardarCancion(dtoInput, data); err != nil {
		http.Error(w, "Error guardando la canción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Canción almacenada correctamente"))
}
