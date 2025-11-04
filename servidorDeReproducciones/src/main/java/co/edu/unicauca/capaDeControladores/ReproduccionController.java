package co.edu.unicauca.capaDeControladores;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import co.edu.unicauca.fachadaServices.services.ReproduccionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/reproducciones")
public class ReproduccionController {

    @Autowired
    private ReproduccionService service;

    @PostMapping
    public ResponseEntity<String> registrar(@RequestBody ReproduccionesDTOEntrada dto) {
        System.out.println("[ECO][Java] Reproducción recibida -> Usuario: " 
            + dto.getUserId() + " | Canción: " + dto.getSongId() + " | Fecha y Hora: " + dto.getFechaHora());
        
        service.registrar(dto);
        return ResponseEntity.ok("Reproducción registrada correctamente");
    }
}

