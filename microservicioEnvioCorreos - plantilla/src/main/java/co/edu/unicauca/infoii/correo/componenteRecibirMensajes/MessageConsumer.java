package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;

import org.springframework.amqp.rabbit.annotation.RabbitListener;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@Service
public class MessageConsumer {
    @RabbitListener(queues = "notificaciones_canciones")
    public void recibirMensaje(CancionAlmacenarDTOInput objClienteCreado) {
        // Lógica para procesar el mensaje recibido
        System.out.println("Datos de la canción recibidos");
        System.out.println("Enviando correo electrónico...");
        Simulacion.simular(10000,"Enviando correo...");
        System.out.println("Correo enviado al cliente con los siguientes datos:");
        System.out.println("Título: " + objClienteCreado.getTitulo());
        System.out.println("Artista: " + objClienteCreado.getArtista());
        System.out.println("Género: " + objClienteCreado.getGenero());
        System.out.println("Año de Lanzamiento: " + objClienteCreado.getAnio_lanzamiento());
        System.out.println("Duración: " + objClienteCreado.getDuracion());
        System.out.println("Idioma: " + objClienteCreado.getIdioma());
        System.out.println("-------------------------------------");
        // Imprimir fecha y hora actual
        LocalDateTime ahora = LocalDateTime.now();
        DateTimeFormatter formato = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
        System.out.println("Fecha y hora en que fue agregada: " + ahora.format(formato));

        // Mensaje inspirador al final
        System.out.println("\nMensaje inspirador: La música no solo escucha tus emociones — las transforma. Sigue creando y compartiendo sonido.");
    }
}
    