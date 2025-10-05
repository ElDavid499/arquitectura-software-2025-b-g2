# 🧱 Publicación de la Estructura del Proyecto — Patrón Arquitectónico

Este documento describe la **estructura general del proyecto**, identificando los componentes que conforman su arquitectura y su respectiva organización dentro del framework asociado.

---

## ⚙️ Backend

El backend del sistema está diseñado bajo una arquitectura **basada en microservicios**, en donde cada módulo funciona de manera independiente, con su propio entorno, dependencias y configuración Docker.  
Los servicios se comunican mediante un **API Gateway** central (`api-gateway.py`).

**Estructura:**

```
backend/
├── auth-service/             # Servicio de autenticación y gestión de usuarios
│   ├── main.py
│   ├── requirements.txt
│   ├── Dockerfile
│   └── README.md
│
├── fleet-service/            # Servicio de gestión de flotas
├── notifications-service/    # Servicio de envío de notificaciones
├── shipments-service/        # Servicio de gestión de envíos
└── tracking-service/         # Servicio de rastreo de paquetes
```

Cada microservicio cuenta con:
- `main.py` → punto de entrada principal.  
- `requirements.txt` → dependencias de Python.  
- `Dockerfile` → definición del contenedor.  
- `README.md` → documentación específica del servicio.

---

## 💻 Frontend

El frontend está implementado como una aplicación web moderna, estructurada dentro del directorio:

```
frontend/
└── web-app/
    ├── package.json          # Dependencias y scripts del proyecto
    ├── README.md             # Documentación específica del frontend
    └── src/ (si aplica)      # Componentes, vistas y estilos
```

Esta aplicación se conecta a las APIs del backend para presentar interfaces dinámicas al usuario final.

---

## 🗄️ Database

Aunque el proyecto no incluye una carpeta `database/` explícita, las bases de datos están configuradas **dentro del archivo `docker-compose.yml`** y pueden ser independientes para cada servicio, siguiendo el patrón **Database per Service**.

**Ejemplo de estructura propuesta:**

```
database/
├── init.sql                      # Script de creación de tablas y datos iniciales
└── docker-compose.override.yml   # Configuración adicional para bases de datos
```

Cada microservicio podría conectarse a su propia base de datos (por ejemplo, PostgreSQL o MongoDB), definida en las variables de entorno del contenedor correspondiente.

---

## 🐳 DevOps / Docker

El sistema está completamente **contenedorizado**.  
Cada servicio cuenta con su propio `Dockerfile` y el archivo `docker-compose.yml` en la raíz orquesta todos los contenedores del proyecto.

**Archivos relacionados:**
```
docker-compose.yml          # Define los servicios y dependencias
backend/*/Dockerfile        # Imagen de cada microservicio
frontend/web-app/Dockerfile # (opcional) Imagen del frontend
```

**Comandos básicos de despliegue:**

```bash
# Construir e iniciar los contenedores
docker-compose up -d

# Verificar estado de los servicios
docker ps

# Detener todos los servicios
docker-compose down
```

Esto permite implementar una infraestructura reproducible, escalable y fácilmente desplegable.

---

## 📂 Diagrams y Tests

```
diagrams/
├── fasttrack_components.png
├── fasttrack_components.puml

tests/
└── test_shipments.py
```

- Los diagramas muestran la arquitectura del sistema.  
- Las pruebas unitarias verifican el comportamiento de los módulos del backend.

---

## 🧩 Conclusión

El proyecto **Patrón Arquitectónico** presenta una estructura sólida y modular, que permite mantener la independencia entre componentes, facilitar el mantenimiento y soportar prácticas modernas de despliegue continuo (CI/CD).  
Esta separación clara en **Backend**, **Frontend**, **Database** y **DevOps/Docker** garantiza una alta escalabilidad y flexibilidad para entornos productivos.

