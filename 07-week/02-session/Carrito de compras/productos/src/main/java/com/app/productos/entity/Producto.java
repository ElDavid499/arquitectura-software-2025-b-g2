package com.app.productos.entity;

import jakarta.persistence.*;


@Entity
public class Producto {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String nombre;
    private Double precio;

    @ManyToOne
    private Categoria categoria;
}
