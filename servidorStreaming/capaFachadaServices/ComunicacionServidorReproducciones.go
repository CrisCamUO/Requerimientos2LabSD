package capafachadaservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura que se enviará al microservicio de reproducciones
type ReproduccionDTO struct {
	IdUsuario int32 `json:"idUsuario"`
	IdCancion int32 `json:"idCancion"`
}

// Envía una reproducción al servidor de reproducciones (Spring Boot)
func EnviarReproduccion(idUsuario int32, idCancion int32) error {
	url := "http://localhost:2020/reproducciones"

	body, _ := json.Marshal(ReproduccionDTO{
		IdUsuario: idUsuario,
		IdCancion: idCancion,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error enviando reproducción: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("el servidor respondió con código: %v", resp.StatusCode)
	}

	fmt.Println("Reproducción registrada con éxito en el servidor de reproducciones")
	return nil
}