package utilidades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	pb "servidor.local/grpc-servidor/serviciosStreaming"
)

func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Fatalf("error decodificando MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

func RecibirCancion(stream pb.AudioService_EnviarCancionMedianteStreamClient, writer *io.PipeWriter, canalSincronizacion chan struct{}, idUsuario int, idCancion int) {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Canción recibida completa.")
			writer.Close()
			break
		}
		if err != nil {
			log.Fatalf("Error recibiendo chunk: %v", err)
		}
		noFragmento++
		fmt.Printf("\n Fragmento #%d recibido (%d bytes) reproduciendo...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			log.Printf("Error escribiendo en pipe: %v", err)
			break
		}
	}
	// Esperar hasta que termine la reproducción
	<-canalSincronizacion
	fmt.Println("Reproducción finalizada.")

	// Cerrar el writer aquí, no dentro del bucle
	writer.Close()

	//registrar la reproduccion
	RegistrarReproduccion(idUsuario, idCancion)
}

func VerPreferencias(userID int) {
	url := fmt.Sprintf("http://localhost:2021/preferencias/calcular?idUsuario=%d", userID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error llamando al servidor de preferencias: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("\nPreferencias recibidas:")
	fmt.Println(string(body))
}

func RegistrarReproduccion(idUsuario int, idCancion int) {
	url := "http://localhost:2020/reproducciones"

	// Crear la estructura de la reproducción
	reproduccion := map[string]interface{}{
		"userId":    fmt.Sprintf("%d", idUsuario),
		"songId":    fmt.Sprintf("%d", idCancion),
		"fechaHora": time.Now().Format("2006-01-02 15:04:05"),
	}

	// Convertir el mapa a JSON
	body, err := json.Marshal(reproduccion)
	if err != nil {
		log.Printf("Error convirtiendo a JSON: %v", err)
		return
	}

	// Enviar al servidor de reproducciones
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error enviando la reproducción: %v", err)
		return
	}
	defer resp.Body.Close()

	// Leer respuesta
	respuesta, _ := io.ReadAll(resp.Body)
	fmt.Println("\nReproducción registrada correctamente en el servidor de reproducciones.")
	fmt.Println(string(respuesta))
}
