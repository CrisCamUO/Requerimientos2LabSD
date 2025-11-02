package co.edu.unicauca.capaDeControladores;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.services.PreferenciasServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/preferencias")
public class ControladorPreferenciasRest {

    @Autowired
    private PreferenciasServiceImpl preferenciasService;

    /**
     * Endpoint REST para calcular las preferencias de un usuario.
     * Ejemplo: POST http://localhost:2021/preferencias/calcular?idUsuario=1
     */
    @PostMapping("/calcular")
    public PreferenciasDTORespuesta calcularPreferencias(@RequestParam("idUsuario") int idUsuario) {
        System.out.println(" Petición REST recibida: calcular preferencias del usuario " + idUsuario);
        return preferenciasService.getReferencias(idUsuario);
    }
}
