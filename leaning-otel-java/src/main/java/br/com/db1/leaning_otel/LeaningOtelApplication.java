package br.com.db1.leaning_otel;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class LeaningOtelApplication {

    public static void main(String[] args) {
        SpringApplication.run(LeaningOtelApplication.class, args);
    }

}
