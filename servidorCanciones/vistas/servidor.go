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

	//inciarServidorREST() // Inicia el servidor REST en una gorutina separada
	go iniciarServidorREST()

	// Iniciar el servidor gRPC

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

func iniciarServidorREST() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/canciones", capaControladoresCancion.ListarCancionesREST)

	port := 5051
	fmt.Printf("Servidor REST de Canciones escuchando en puerto %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatalf("Error iniciando servidor REST: %v", err)
	}
}