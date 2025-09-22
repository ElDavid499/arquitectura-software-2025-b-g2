package com.app.pedidos.entity;

import jakarta.persistence.*;

@Data
@Entity
public class ItemCarrito {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private int cantidad;

    @ManyToOne
    private Carrito carrito;

    @ManyToOne
    private Producto producto;
}
