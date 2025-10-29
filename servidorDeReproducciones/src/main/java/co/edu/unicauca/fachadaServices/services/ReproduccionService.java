package co.edu.unicauca.fachadaServices.services;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class ReproduccionService {

    private List<ReproduccionesDTOEntrada> lista = new ArrayList<>();

    public void registrar(ReproduccionesDTOEntrada dto) {
        lista.add(dto);
        System.out.println("[ECO][Java] Registro guardado. Total reproducciones: " + lista.size());
    }

    public List<ReproduccionesDTOEntrada> listarPorUsuario(String userId) {
        List<ReproduccionesDTOEntrada> res = new ArrayList<>();
        for (ReproduccionesDTOEntrada r : lista) {
            if (r.getUserId().equals(userId)) {
                res.add(r);
            }
        }
        return res;
    }
}

