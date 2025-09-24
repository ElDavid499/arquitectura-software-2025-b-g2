# Proyecto CRUD en Go — Comparación de Arquitecturas

Este proyecto implementa un CRUD en **Go con Gin** en dos escenarios diferentes:

1. **Escenario 1 — Arquitectura por capas (Layered Architecture)**  
2. **Escenario 2 — Arquitectura por módulos (Modular Architecture)**

---

## Escenario 1 — Arquitectura por capas

En este enfoque se organiza el código por **responsabilidades**. Cada capa cumple una función específica y se comunica con la siguiente.

###  Estructura de carpetas

```
escenario1/
├── controllers/ # Manejo de peticiones HTTP y llamadas a servicios
├── dto/ # Data Transfer Objects
├── entity/ # Modelos de dominio
├── repository/ # Acceso y gestión de datos
├── service/ # Lógica de negocio
├── utils/ # Funciones de apoyo
└── main.go # Punto de entrada de la aplicación
```

###  Endpoints CRUD de User

| Método | Ruta         | Descripción            |
|--------|-------------|------------------------|
| POST   | `/users`    | Crear usuario          |
| GET    | `/users`    | Listar usuarios        |
| GET    | `/users/:id`| Obtener usuario por ID |
| PUT    | `/users/:id`| Actualizar usuario     |
| DELETE | `/users/:id`| Eliminar usuario       |

---

## Escenario 2 — Arquitectura por módulos

En este enfoque el proyecto se divide en **módulos independientes**, cada uno con su propio mini-MVC (entidad, repositorio, servicio y handlers). Esto favorece la escalabilidad y la autonomía de cada módulo.

###  Estructura de carpetas

```
escenario2-by-module/
├── user/
│   ├── entity/
│   ├── repository/
│   ├── service/
│   ├── handlers/
│   └── main.go
├── product/
│   ├── entity/
│   ├── repository/
│   ├── service/
│   ├── handlers/
│   └── main.go
└── order/
    ├── entity/
    ├── repository/
    ├── service/
    ├── handlers/
    └── main.go
```

###  Ejecución por módulo

Cada módulo puede ejecutarse de manera independiente:

```bash
go run user/main.go
go run product/main.go
go run order/main.go
```

---

## Comparación entre escenarios

###  Ventajas

| Escenario 1 (por capas) | Escenario 2 (por módulos) |
|--------------------------|----------------------------|
| Claridad en separación de responsabilidades. | Cada módulo es autónomo y reutilizable. |
| Fácil de mantener en proyectos pequeños/medianos. | Facilita escalabilidad y microservicios. |
| Ideal para aplicaciones monolíticas. | Equipos pueden trabajar en módulos distintos. |

###  Desventajas

| Escenario 1 (por capas) | Escenario 2 (por módulos) |
|--------------------------|----------------------------|
| Código acoplado en un solo monolito. | Más código duplicado entre módulos. |
| Escalabilidad limitada. | Mayor complejidad inicial. |
| Menos flexible para microservicios. | Consumo extra de recursos. |

###  Tipos de proyectos recomendados

- **Escenario 1**: proyectos pequeños/medianos, monolitos.  
- **Escenario 2**: proyectos grandes, escalables, basados en microservicios.

###  Consumo de recursos

| Escenario   | Memoria/CPU           | Escalabilidad |
|-------------|------------------------|---------------|
| Escenario 1 | Bajo consumo inicial. | Limitada (difícil escalar monolito). |
| Escenario 2 | Mayor consumo inicial. | Alta (ideal para nube y balanceo). |

---

## Conclusión

- **Escenario 1**: recomendado para simplicidad y rapidez en proyectos pequeños.  
- **Escenario 2**: recomendado para modularidad y escalabilidad en proyectos grandes o distribuidos.  
