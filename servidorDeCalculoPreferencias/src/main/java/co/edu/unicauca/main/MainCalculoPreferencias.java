package co.edu.unicauca.main;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class MainCalculoPreferencias {
    public static void main(String[] args) {
        System.out.println("[ECO][Java] Servidor de Cálculo de Preferencias iniciado correctamente en puerto 2021");
        SpringApplication.run(MainCalculoPreferencias.class, args);
    }
}