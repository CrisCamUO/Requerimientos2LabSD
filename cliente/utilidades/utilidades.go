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

type PreferenciaGenero struct {
	NombreGenero       string `json:"nombreGenero"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

type PreferenciaArtista struct {
	NombreArtista      string `json:"nombreArtista"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

type PreferenciaIdioma struct {
	NombreIdioma       string `json:"nombreIdioma"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

type PreferenciasRespuesta struct {
	IdUsuario            int                  `json:"idUsuario"`
	PreferenciasGeneros  []PreferenciaGenero  `json:"preferenciasGeneros"`
	PreferenciasArtistas []PreferenciaArtista `json:"preferenciasArtistas"`
	PreferenciasIdiomas  []PreferenciaIdioma  `json:"preferenciasIdiomas"`
}

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

func RecibirCancion(stream pb.AudioService_EnviarCancionMedianteStreamClient, writer *io.PipeWriter, canalSincronizacion chan struct{}, idUsuario string, idCancion string) {
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
			log.Printf("Se detuvo la lectura en pipe: %v", err)
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

func VerPreferencias(userID string) {
	url := fmt.Sprintf("http://localhost:2021/preferencias/calcular?idUsuario=%s", userID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("❌ Error llamando al servidor de preferencias: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var prefs PreferenciasRespuesta
	if err := json.Unmarshal(body, &prefs); err != nil {
		fmt.Println("⚠️ No se pudo procesar la respuesta del servidor:")
		fmt.Println(string(body))
		return
	}

	fmt.Printf("\n🎧 Preferencias del usuario #%d 🎧\n", prefs.IdUsuario)
	fmt.Println("───────────────────────────────────────")

	if len(prefs.PreferenciasGeneros) == 0 && len(prefs.PreferenciasArtistas) == 0 {
		fmt.Println("🚫 No hay suficientes reproducciones para calcular preferencias.")
		return
	}

	if len(prefs.PreferenciasGeneros) > 0 {
		fmt.Println("\n🎵 Géneros preferidos:")
		for i, g := range prefs.PreferenciasGeneros {
			fmt.Printf("   %d. %s (%d reproducciones)\n", i+1, g.NombreGenero, g.NumeroPreferencias)
		}
	}

	if len(prefs.PreferenciasArtistas) > 0 {
		fmt.Println("\n🎤 Artistas preferidos:")
		for i, a := range prefs.PreferenciasArtistas {
			fmt.Printf("   %d. %s (%d reproducciones)\n", i+1, a.NombreArtista, a.NumeroPreferencias)
		}
	}
	if len(prefs.PreferenciasIdiomas) > 0 {
		fmt.Println("\n🗣️ Idiomas preferidos:")
		for i, id := range prefs.PreferenciasIdiomas {
			fmt.Printf("   %d. %s (%d reproducciones)\n", i+1, id.NombreIdioma, id.NumeroPreferencias)
		}
	}

	fmt.Println("───────────────────────────────────────\n")
}

func RegistrarReproduccion(idUsuario string, idCancion string) {
	url := "http://localhost:2020/reproducciones"

	// Crear la estructura de la reproducción
	reproduccion := map[string]interface{}{
		"userId":    idUsuario,
		"songId":    idCancion,
		"fechaHora": time.Now().Format("2006-01-02 15:04:05"),
	}

	jsonData, _ := json.Marshal(reproduccion)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error enviando reproducción: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
