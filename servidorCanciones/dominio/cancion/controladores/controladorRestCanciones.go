package controladores

import (
	"encoding/json"
	"net/http"

	fachada "servidor.local/grpc-servidorCancion/dominio/cancion/fachadaCancionesServices"
)

func ListarCancionesREST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	canciones := fachada.ObtenerCancionesParaREST() // Llamamos al nuevo método

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(canciones)
}
