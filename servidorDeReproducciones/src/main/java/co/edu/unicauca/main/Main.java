package co.edu.unicauca.main;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication(scanBasePackages = "co.edu.unicauca")
public class Main {

    public static void main(String[] args) {
        SpringApplication.run(Main.class, args);
        System.out.println("[ECO][Java] ServidorDeReproducciones iniciado correctamente en puerto 2020 ");
    }
}


