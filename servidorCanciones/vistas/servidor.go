package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	capaControladoresCancion "servidor.local/grpc-servidorCancion/dominio/cancion/controladores"
	pb "servidor.local/grpc-servidorCancion/serviciosCancion"
)

func main() {

	go iniciarSevidorREST()

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiciosCancionesServer(grpcServer, &capaControladoresCancion.ControladorCanciones{})
	fmt.Printf("Servidor de Canciones escuchando en %s...\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

}

func iniciarSevidorREST() {
	ctrl := capaControladoresCancion.NuevoControladorAlmacenamientoCanciones()

	http.HandleFunc("/canciones/almacenamiento", ctrl.AlmacenarCancion)
	// TODO: registrar el handler de metadatos cuando se implemente, por ejemplo:
	// http.HandleFunc("/canciones/metadatos", ctrl.GuardarMetadatosCancion)

	fmt.Println("✅ Servicio de Tendencias escuchando en el puerto 5000...")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error iniciando el servidor:", err)
	}
}
