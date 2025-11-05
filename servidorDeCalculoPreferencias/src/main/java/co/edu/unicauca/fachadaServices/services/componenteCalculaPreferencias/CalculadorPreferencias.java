package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.CancionDTOEntrada;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;

public class CalculadorPreferencias {

    
    public PreferenciasDTORespuesta calcular(Integer idUsuario,
                                            List<CancionDTOEntrada> canciones,
                                             List<ReproduccionesDTOEntrada> reproducciones) {

        // Crear mapa de canciones por ID (para acceder rápido)
        Map<Integer, CancionDTOEntrada> mapaCanciones = canciones.stream()
            .filter(Objects::nonNull)
            .filter(c -> c.getId() != null)
            .collect(Collectors.toMap(CancionDTOEntrada::getId, c -> c, (a, b) -> a));

        // Contadores de generos y artistas
        Map<String, Integer> contadorGeneros = new HashMap<>();
        Map<String, Integer> contadorArtistas = new HashMap<>();
    Map<String, Integer> contadorIdiomas = new HashMap<>();

        System.out.println("Iniciando calculo de preferencias...");

        // Recorrer las reproducciones del usuario
        for (ReproduccionesDTOEntrada r : reproducciones) {
            if (r.getSongId() == null) continue;

            try {
                Integer idCancionInt = Integer.parseInt(r.getSongId());
                CancionDTOEntrada c = mapaCanciones.get(idCancionInt);

                if (c == null) {
                    System.out.println("No se encontró la cancion con ID " + idCancionInt);
                    continue;
                }

                String genero = c.getGenero() == null ? "Desconocido" : c.getGenero();
                String artista = c.getArtista() == null ? "Desconocido" : c.getArtista();
                String idioma = "Desconocido";
                try {
                    idioma = c.getIdioma() == null ? "Desconocido" : c.getIdioma();
                } catch (Exception ex) {
                    idioma = "Desconocido";
                }

                contadorGeneros.put(genero, contadorGeneros.getOrDefault(genero, 0) + 1);
                contadorArtistas.put(artista, contadorArtistas.getOrDefault(artista, 0) + 1);
                contadorIdiomas.put(idioma, contadorIdiomas.getOrDefault(idioma, 0) + 1);

                System.out.println(" Procesada reproducción -> " + artista + " / " + genero);

            } catch (NumberFormatException e) {
                System.out.println("ID de cancion invalido: " + r.getSongId());
            }
        }

        // Convertir los contadores a listas ordenadas
        
        List<PreferenciaGeneroDTORespuesta> preferenciasGeneros = contadorGeneros.entrySet().stream()
            .map(e -> new PreferenciaGeneroDTORespuesta(e.getKey(), e.getValue()))
            .sorted(Comparator.comparingInt(PreferenciaGeneroDTORespuesta::getNumeroPreferencias).reversed()
                    .thenComparing(PreferenciaGeneroDTORespuesta::getNombreGenero))
            .collect(Collectors.toList());

        List<PreferenciaArtistaDTORespuesta> preferenciasArtistas = contadorArtistas.entrySet().stream()
            .map(e -> new PreferenciaArtistaDTORespuesta(e.getKey(), e.getValue()))
            .sorted(Comparator.comparingInt(PreferenciaArtistaDTORespuesta::getNumeroPreferencias).reversed()
                    .thenComparing(PreferenciaArtistaDTORespuesta::getNombreArtista))
            .collect(Collectors.toList());

    List<PreferenciaIdiomaDTORespuesta> preferenciasIdiomas = contadorIdiomas.entrySet().stream()
        .map(e -> new PreferenciaIdiomaDTORespuesta(e.getKey(), e.getValue()))
        .sorted(Comparator.comparingInt(PreferenciaIdiomaDTORespuesta::getNumeroPreferencias).reversed()
            .thenComparing(PreferenciaIdiomaDTORespuesta::getNombreIdioma))
        .collect(Collectors.toList());

        // Armar el DTO final

        PreferenciasDTORespuesta respuesta = new PreferenciasDTORespuesta();
        respuesta.setIdUsuario(idUsuario);
        respuesta.setPreferenciasGeneros(preferenciasGeneros);
        respuesta.setPreferenciasArtistas(preferenciasArtistas);
    respuesta.setPreferenciasIdiomas(preferenciasIdiomas);

        System.out.println(" Cálculo de preferencias completado.");
        return respuesta;
    }
}
