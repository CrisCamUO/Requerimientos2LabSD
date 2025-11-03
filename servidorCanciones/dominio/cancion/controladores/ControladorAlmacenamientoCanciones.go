package controladores

import (
	"fmt"
	"io"
	"net/http"

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

	r.ParseMultipartForm(50 << 20)
	file, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, _ := io.ReadAll(file)

	//leer los campos del DTO
	dto := dto.CancionAlmacenarDTOInput{
		Titulo:  r.FormValue("titulo"),
		Genero:  r.FormValue("genero"),
		Artista: r.FormValue("artista"),
		Idioma:  r.FormValue("idioma"),
	}

	thisC.fachada.GuardarCancion(dto, data)
}
