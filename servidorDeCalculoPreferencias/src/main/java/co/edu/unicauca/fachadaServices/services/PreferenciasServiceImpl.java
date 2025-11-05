
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

		if (objCanciones == null) {
			System.out.println(" No se recibieron canciones (objCanciones es NULL)");
		} else if (objCanciones.isEmpty()) {
			System.out.println("Lista de canciones vacía");
		} else {
			System.out.println("Canciones recibidas:");
			for (CancionDTOEntrada c : objCanciones) {
				System.out.println("   ID: " + c.getId() + " | Título: " + c.getTitulo() + 
								" | Género: " + c.getGenero() + " | Artista: " + c.getArtista());
			}
		}

		List<ReproduccionesDTOEntrada> reproduccionesUsuario = this.comunicacionServidorReproducciones.obtenerReproduccionesRemotas(id);
		System.out.println("Reproducciones obtenidas del servidor de reproducciones para el usuario  " + id);
		System.out.println("🟦 Llamando a obtenerReproduccionesRemotas con id = " + id);
		
		for (ReproduccionesDTOEntrada r : reproduccionesUsuario) {
        System.out.println("Repro encontrada -> user: " + r.getUserId() + " song: " + r.getSongId());
    	}
	
		return this.calculadorPreferencias.calcular(id, objCanciones, reproduccionesUsuario);
	}
}
