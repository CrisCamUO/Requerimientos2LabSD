package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import feign.Feign;
import feign.jackson.JacksonDecoder;
import feign.jackson.JacksonEncoder;

import java.util.ArrayList;
import java.util.List;

public class ComunicacionServidorReproducciones {

   private static final String BASE_URL = "http://localhost:2020";
   private final ReproduccionesRemoteClient client;

    public ComunicacionServidorReproducciones(){
        this.client = Feign.builder()
        .encoder(new JacksonEncoder())
        .decoder(new JacksonDecoder())
        .target(ReproduccionesRemoteClient.class,BASE_URL);
    }
    
    public List<ReproduccionesDTOEntrada> obtenerReproduccionesRemotas(Integer idUsuario){
        try {

            if (idUsuario == null) {
            System.out.println("idUsuario es NULL. No se puede consultar reproducciones.");
            return new ArrayList<>();
        }
            System.out.println("Consultando reproducciones remotas para el usuario " + idUsuario);
            List<ReproduccionesDTOEntrada> reproducciones = client.obtenerReproducciones(idUsuario);
            System.out.println("Respuesta del servidor de reproducciones: " + reproducciones);
            return reproducciones != null ? reproducciones : new ArrayList<>(); 
        
        } catch (Exception e) {
            System.err.println(" Error consultando reproducciones: " + e.getMessage());
            return new ArrayList<>();
        }
    }
    
}


