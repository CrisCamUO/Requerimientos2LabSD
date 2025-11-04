package co.edu.unicauca.main;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = "co.edu.unicauca")
public class MainCalculoPreferencias {
    public static void main(String[] args) {
        SpringApplication.run(MainCalculoPreferencias.class, args);
        System.out.println("[ECO][Java] Servidor de Cálculo de Preferencias iniciado correctamente en puerto 2021");
    }
}