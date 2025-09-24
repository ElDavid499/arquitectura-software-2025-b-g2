# Taller Arquitectura & Esqueleto 

## 1. Documentación de los casos

### Caso 1 – Arquitectura monolítica tradicional  
- **Ventajas**  
  - Desarrollo más sencillo y rápido en fases iniciales.  
  - Menos complejidad en la configuración de la infraestructura.  
  - Ideal para proyectos pequeños o MVP.  

- **Desventajas**  
  - Dificultad para escalar partes específicas.  
  - Si un módulo falla, puede afectar todo el sistema.  
  - Despliegues más lentos porque se debe compilar y desplegar toda la aplicación.  

- **Proyectos ideales**  
  - Aplicaciones pequeñas o de mediana escala.  
  - Sistemas internos sin altas exigencias de disponibilidad.  

- **Consumo de recursos**  
  - Menor consumo en servidores, ya que todo corre en una sola aplicación.  
  - Fácil de desplegar en un solo contenedor/VM.  

---

### Caso 2 – Arquitectura modular / orientada a microservicios  
- **Ventajas**  
  - Cada módulo o servicio puede escalar de manera independiente.  
  - Favorece la reutilización y la separación de responsabilidades.  
  - Facilita la transición hacia una arquitectura de microservicios.  
  - Mayor flexibilidad tecnológica (cada módulo podría evolucionar distinto).  

- **Desventajas**  
  - Requiere mayor configuración e infraestructura.  
  - Mayor complejidad en la comunicación entre módulos.  
  - Puede ser excesivo para proyectos muy pequeños.  

- **Proyectos ideales**  
  - Aplicaciones con visión de crecimiento rápido.  
  - Plataformas que se espera tengan muchos usuarios concurrentes.  
  - Casos donde se piensa evolucionar hacia microservicios en el corto plazo.  

- **Consumo de recursos**  
  - Mayor demanda de recursos de infraestructura (servidores, contenedores, balanceadores).  
  - Implica monitoreo y orquestación (Docker, Kubernetes, etc.) si escala a microservicios.  

---

### Comparación  
| Aspecto              | Caso 1: Monolítico        | Caso 2: Modular/Microservicios |
|----------------------|--------------------------|--------------------------------|
| **Complejidad**      | Baja                     | Media / Alta                   |
| **Escalabilidad**    | Limitada                 | Alta                           |
| **Despliegue**       | Rápido, único artefacto  | Varias unidades independientes |
| **Mantenimiento**    | Puede volverse difícil   | Más organizado por módulos     |
| **Infraestructura**  | Menos exigente           | Más exigente                   |

---

### Elección para proyectos de crecimiento rápido  
Si el proyecto **tendrá un crecimiento acelerado** y a futuro se quiere migrar a **microservicios**, **el Caso 2 es el ideal**, ya que la arquitectura modular te prepara desde temprano para separar responsabilidades y escalar partes específicas del sistema.

---

## 2. Proyecto planteado: **Shopping Cart**

Se construyó un **esqueleto en Spring Boot** bajo un enfoque **multi-módulo**, con **3 módulos** (`usuarios`, `productos`, `pedidos`) y **7 entidades**.

### Estructura de carpetas
```
shopping-cart/
│── pom.xml
│
├── usuarios/
│   ├── pom.xml
│   └── src/main/java/com/app/usuarios/entity/
│       ├── Usuario.java
│       └── Rol.java
│
├── productos/
│   ├── pom.xml
│   └── src/main/java/com/app/productos/entity/
│       ├── Producto.java
│       └── Categoria.java
│
└── pedidos/
    ├── pom.xml
    └── src/main/java/com/app/pedidos/entity/
        ├── Carrito.java
        ├── ItemCarrito.java
        └── Pedido.java
```

### Módulos
- **Usuarios** → gestiona usuarios y roles.  
- **Productos** → catálogo de productos y categorías.  
- **Pedidos** → carrito de compras, items y pedidos.  

### Entidades (7 en total)
- `Usuario`, `Rol`  
- `Producto`, `Categoria`  
- `Carrito`, `ItemCarrito`, `Pedido`  

---

## 3. Ejemplo de configuración con Maven (`pom.xml`)

### `pom.xml` principal (raíz del proyecto)
```xml
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">

    <modelVersion>4.0.0</modelVersion>
    <groupId>com.app</groupId>
    <artifactId>shopping-cart</artifactId>
    <version>1.0.0</version>
    <packaging>pom</packaging>

    <modules>
        <module>usuarios</module>
        <module>productos</module>
        <module>pedidos</module>
    </modules>

    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-dependencies</artifactId>
                <version>3.3.0</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>

</project>
```

### `usuarios/pom.xml`
```xml
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">

    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>com.app</groupId>
        <artifactId>shopping-cart</artifactId>
        <version>1.0.0</version>
    </parent>

    <artifactId>usuarios</artifactId>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-jpa</artifactId>
        </dependency>
    </dependencies>
</project>
```

---

## 4. Entidades por módulo

### **Módulo: usuarios**
```java
@Entity
public class Usuario {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String nombre;
    private String email;

    @ManyToOne
    private Rol rol;
}
```

```java
@Entity
public class Rol {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String nombre;
}
```

### **Módulo: productos**
```java
@Entity
public class Producto {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String nombre;
    private Double precio;

    @ManyToOne
    private Categoria categoria;
}
```

```java
@Entity
public class Categoria {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String nombre;
}
```

### **Módulo: pedidos**
```java
@Entity
public class Carrito {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToMany(mappedBy = "carrito", cascade = CascadeType.ALL)
    private List<ItemCarrito> items;
}
```

```java
@Entity
public class ItemCarrito {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private int cantidad;

    @ManyToOne
    private Carrito carrito;

    @ManyToOne
    private Producto producto;
}
```

```java
@Entity
public class Pedido {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private LocalDateTime fecha;

    @OneToOne
    private Carrito carrito;
}
```

---

## 5. Diagrama de módulos y entidades

### **Diagrama de paquetes (módulos)**
```
shopping-cart
 ├── usuarios
 │     └── Usuario, Rol
 ├── productos
 │     └── Producto, Categoria
 └── pedidos
       └── Carrito, ItemCarrito, Pedido
```

### **Relaciones principales entre entidades**
- `Usuario` → tiene un `Rol`.  
- `Producto` → pertenece a una `Categoria`.  
- `Carrito` → contiene muchos `ItemCarrito`.  
- `ItemCarrito` → une `Carrito` con `Producto`.  
- `Pedido` → se genera a partir de un `Carrito`.  
