# Database Module - PostgreSQL

Este módulo contiene los scripts de inicialización de la base de datos del proyecto.

## 📦 Configuración
- **Motor:** PostgreSQL
- **Versión recomendada:** 15 o superior
- **Puerto:** 5432
- **Usuario:** postgres
- **Contraseña:** postgres
- **Base de datos:** fasttrack_db

## 📂 Archivos
- `init.sql`: crea las tablas iniciales (usuarios, envíos, etc.)

## 🚀 Uso con Docker
El contenedor de PostgreSQL se levanta automáticamente al ejecutar:
```bash
docker-compose up -d
```
