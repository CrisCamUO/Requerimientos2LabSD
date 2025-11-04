package capacontroladores

import (
	"fmt"
	capafachadaservices "servidor.local/grpc-servidor/capaFachadaServices"
	pb "servidor.local/grpc-servidor/serviciosStreaming"
	
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// Implementación del procedimiento remoto que recibe el título de una canción y envia el archivo de audio en fragmentos mediante un stream.
func (s *ControladorServidor) EnviarCancionMedianteStream(req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {
	// Enviar los fragmentos de audio
	err := capafachadaservices.StreamAudioFile(
		req.Id,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})
	if err != nil {
		fmt.Println("Error al enviar el audio:", err)
		return err
	}

	// Registrar la reproducción en el servidor de reproducciones
	err = capafachadaservices.EnviarReproduccion(req.IdUsuario, req.Id)
	if err != nil {
		fmt.Println("No se pudo registrar la reproducción:", err)
	} else {
		fmt.Println("Reproducción registrada correctamente")
	}

	return nil
}
