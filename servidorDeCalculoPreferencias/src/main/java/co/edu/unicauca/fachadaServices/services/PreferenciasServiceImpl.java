
package co.edu.unicauca.fachadaServices.services;

import org.springframework.stereotype.Service;
import java.util.List;

import co.edu.unicauca.fachadaServices.DTO.CancionDTOEntrada;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias.CalculadorPreferencias;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones.ComunicacionServidorCanciones;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones.ComunicacionServidorReproducciones;

@Service
public class PreferenciasServiceImpl implements IPreferenciasService {
	
	private ComunicacionServidorCanciones comunicacionServidorCanciones;
	private ComunicacionServidorReproducciones comunicacionServidorReproducciones;
	private CalculadorPreferencias calculadorPreferencias;

	public PreferenciasServiceImpl() {
		this.comunicacionServidorCanciones = new ComunicacionServidorCanciones();
		this.comunicacionServidorReproducciones = new ComunicacionServidorReproducciones();
		this.calculadorPreferencias = new CalculadorPreferencias();
	}

	@Override
	public PreferenciasDTORespuesta getReferencias(Integer id) {
		System.out.println("Obteniendo preferencias para el usuario con ID: " + id);
		List<CancionDTOEntrada> objCanciones = this.comunicacionServidorCanciones.obtenerCancionesRemotas();
		System.out.println("Canciones obtenidas del servidor de canciones");

		for (CancionDTOEntrada cancion : objCanciones){
			System.out.println("Cancion obtenida: " + cancion.getTitulo());
			System.out.println("Genero: " + cancion.getGenero());
			System.out.println("Artista: " + cancion.getArtista());
			
		}
		List<ReproduccionesDTOEntrada> reproduccionesUsuario = this.comunicacionServidorReproducciones.obtenerReproduccionesRemotas(id);
		System.out.println("Reproducciones obtenidas del servidor de reproducciones para el usuario  " + id);
		for (ReproduccionesDTOEntrada reproduccion : reproduccionesUsuario){
			System.out.println(reproduccion.getUserId() + " " + reproduccion.getSongId());
		}
	
		return this.calculadorPreferencias.calcular(id, objCanciones, reproduccionesUsuario);
	}
}
