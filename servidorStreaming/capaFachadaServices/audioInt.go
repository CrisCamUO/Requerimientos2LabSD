package capafachadaservices

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func StreamAudioFile(idCancion int32, funcionParaEnviarFragmento func([]byte) error) error {
	log.Printf("Canción solicitada: %d", idCancion)
	file, err := os.Open("canciones/" + strconv.FormatInt(int64(idCancion), 10) + ".mp3")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 64*1024) // 64 KB se envian por fragmento
	fragmento := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Println("Cancion enviada completamente desde la fachada")
			break
		}

		if err != nil {
			return fmt.Errorf("error leyendo el archivo: %w", err)
		}
		fragmento++
		log.Printf("Fragmento #%d leido (%d bytes) y enviando", fragmento, n)
		// Ejecutamos la función para enviar el fragmento al cliente
		err = funcionParaEnviarFragmento(buffer[:n])
		if err != nil {
			return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
		}
	}
	return nil
}

// StreamAudioFileByName busca el archivo usando el campo 'Nombre' recibido en la petición
// y lo envía en fragmentos al invocador.
func StreamAudioFileByName(nombre string, funcionParaEnviarFragmento func([]byte) error) error {
	log.Printf("Canción solicitada por nombre: %s", nombre)

	// Construir posibles nombres/paths a probar
	candidates := []string{}
	// Tal como llega
	candidates = append(candidates, nombre+".mp3")
	// Reemplazar '_' por ' '
	nombreSpaces := strings.ReplaceAll(nombre, "_", " ")
	if nombreSpaces != nombre {
		candidates = append(candidates, nombreSpaces+".mp3")
	}

	// Rutas a probar explícitas
	pathsToTry := []string{}
	for _, c := range candidates {
		pathsToTry = append(pathsToTry, filepath.Join("..", "servidorCanciones", "Audios", c))
		pathsToTry = append(pathsToTry, filepath.Join("canciones", c))
	}

	// Buscar coincidencias parciales en los directorios listados para mayor robustez
	searchDirs := []string{filepath.Join("..", "servidorCanciones", "Audios"), "canciones"}
	lowerQuery := strings.ToLower(strings.ReplaceAll(nombre, "_", " "))
	for _, dir := range searchDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			log.Printf("No se pudo leer directorio %s: %v", dir, err)
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			fname := e.Name()
			fnameLower := strings.ToLower(fname)
			if strings.Contains(fnameLower, lowerQuery) || strings.Contains(fnameLower, strings.ToLower(nombre)) {
				pathsToTry = append(pathsToTry, filepath.Join(dir, fname))
			}
		}
	}

	var file *os.File
	var err error
	for _, p := range pathsToTry {
		file, err = os.Open(p)
		if err == nil {
			log.Printf("Abriendo archivo de audio en: %s", p)
			break
		}
		log.Printf("No se pudo abrir en %s: %v", p, err)
	}
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo en ninguna ruta posible: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 64*1024) // 64 KB se envian por fragmento
	fragmento := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Println("Cancion enviada completamente desde la fachada (por nombre)")
			break
		}

		if err != nil {
			return fmt.Errorf("error leyendo el archivo: %w", err)
		}
		fragmento++
		log.Printf("Fragmento #%d leido (%d bytes) y enviando", fragmento, n)
		// Ejecutamos la función para enviar el fragmento al cliente
		err = funcionParaEnviarFragmento(buffer[:n])
		if err != nil {
			return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
		}
	}
	return nil
}
