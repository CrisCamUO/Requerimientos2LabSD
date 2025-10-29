package co.edu.unicauca.fachadaServices.DTO;

import java.io.Serializable;
import java.util.List;


import lombok.Data;

@Data
public class PreferenciasDTORespuesta implements Serializable {

    private int idUsuario;
    private List<PreferenciaArtistaDTORespuesta> PreferenciasArtistas;
    private List<PreferenciaGeneroDTORespuesta> preferenciasGeneros;
}
