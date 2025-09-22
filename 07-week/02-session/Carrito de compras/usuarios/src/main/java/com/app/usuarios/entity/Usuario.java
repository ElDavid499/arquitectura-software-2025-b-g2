package com.app.usuarios.entity;

import jakarta.persistence.*;

@Data
@Entity
public class Usuario {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String nombre;
    private String email;

    @ManyToOne
    private Rol rol;
}
