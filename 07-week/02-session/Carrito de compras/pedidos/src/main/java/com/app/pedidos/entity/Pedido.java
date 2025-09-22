package com.app.pedidos.entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Data
@Entity
public class Pedido {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private LocalDateTime fecha;

    @OneToOne
    private Carrito carrito;
}
