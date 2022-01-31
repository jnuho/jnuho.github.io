package com.spring.book.springboot;


import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

import java.util.Collections;

@EnableJpaAuditing
@SpringBootApplication
public class Application {

  public static void main(String[] args) {
//    SpringApplication.run(Application.class, args);
    SpringApplication app = new SpringApplication(Application.class);
    app.setDefaultProperties(
        Collections.singletonMap("server.port", "8081")
    );
    app.run(args);
  }

}
