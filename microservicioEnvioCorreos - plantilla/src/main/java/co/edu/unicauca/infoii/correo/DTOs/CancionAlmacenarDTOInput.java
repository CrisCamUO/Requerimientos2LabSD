package co.edu.unicauca.infoii.correo.DTOs;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class CancionAlmacenarDTOInput {
    private String titulo;
    private String artista;
    private String genero;
    private String anio_lanzamiento;
    private String duracion;
    private String idioma;

    public CancionAlmacenarDTOInput() {
    }

    public String getTitulo() {
        return titulo;
    }
    public void setTitulo(String titulo) {
        this.titulo = titulo;
    }
    public String getArtista() {
        return artista;
    }
    public void setArtista(String artista) {
        this.artista = artista;
    }
    public String getGenero() {
        return genero;
    }
    public void setGenero(String genero) {
        this.genero = genero;
    }
    public String getAnio_lanzamiento() {
        return anio_lanzamiento;
    }
    public void setAnio_lanzamiento(String anio_lanzamiento) {
        this.anio_lanzamiento = anio_lanzamiento;
    }
    public String getDuracion() {
        return duracion;
    }
    public void setDuracion(String duracion) {
        this.duracion = duracion;
    }
    public String getIdioma() {
        return idioma;
    }
    public void setIdioma(String idioma) {
        this.idioma = idioma;
    }
    
}


	