package capaaccesoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

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

func (r *RepositorioCanciones) GuardarCancion(titulo string, genero string, artista string, data []byte) error {
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
