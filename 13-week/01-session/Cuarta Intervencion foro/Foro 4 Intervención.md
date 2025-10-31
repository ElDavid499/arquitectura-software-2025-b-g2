# Foro: Mejores prácticas de experiencia de usuario en apps móviles  
## **Pregunta 4:**  
**En términos arquitectónicos, ¿qué se esperaría de una metodología aplicada a sistemas financieros?**

---

##  **Resumen – Reflexión**

En los sistemas financieros, la arquitectura de software debe garantizar **seguridad, trazabilidad, disponibilidad y cumplimiento normativo**. Por ello, una metodología aplicada en este contexto debe integrar **principios de ingeniería segura**, **control de versiones trazable**, **documentación formal** y **procesos de validación rigurosos**.

Desde una perspectiva arquitectónica, se espera que dicha metodología fomente el **diseño modular y desacoplado**, permitiendo la evolución del sistema sin comprometer su estabilidad. Asimismo, debe incorporar etapas claras de **gestión de riesgos**, **análisis de impacto** y **control de acceso**, asegurando que cada componente cumpla con las políticas de auditoría interna y las regulaciones del sector (como PCI DSS o ISO 27001).

El enfoque arquitectónico ideal combina **procesos iterativos** con una **visión estructurada del ciclo de vida**, tal como lo hace **RUP (Rational Unified Process)** o su adaptación en entornos financieros mediante marcos híbridos (Ágil + RUP). Esto permite mantener **la flexibilidad del desarrollo ágil**, pero sin perder **el control formal requerido** en un entorno de alta criticidad.

En resumen, una metodología para sistemas financieros debe garantizar:
- **Seguridad desde el diseño** (security by design).  
- **Arquitectura escalable y resiliente** (microservicios, contenedores, observabilidad).  
- **Gestión rigurosa de requisitos y trazabilidad**.  
- **Cumplimiento normativo y auditoría permanente**.  
- **Alta disponibilidad y tolerancia a fallos**.  

Todo esto respaldado por una **documentación técnica robusta**, asegurando que el sistema no solo funcione, sino que pueda **probar su integridad y confiabilidad** ante auditorías y organismos de control.

---

## **Representación arquitectónica (C4 – Nivel 2: Diagrama de Contenedores)**

### **Sistema Bancario (visión por contenedores)**

```text
+-------------------------------------------------------+
|                     Sistema Bancario                  |
|-------------------------------------------------------|
|  [Cliente]                                             |
|        |                                                |
|        v                                                |
|  +----------------------+       +-------------------+   |
|  | Aplicación Web/Móvil | <---> | API Gateway / BFF |   |
|  +----------------------+       +-------------------+   |
|             |                              |            |
|             v                              v            |
|   +-------------------+          +-------------------+   |
|   | Servicios Core    | <------> | Servicio de Pagos |   |
|   | Bancarios         |          | Externo/Interno   |   |
|   +-------------------+          +-------------------+   |
|             |                              |            |
|             v                              |            |
|   +-------------------+          +-------------------+   |
|   | Base de Datos     |          | Módulo Seguridad  |   |
|   | Transaccional     |          | y Auditoría       |   |
|   +-------------------+          +-------------------+   |
|                                                       |
|   --> Integración con sistemas externos (ACH, SWIFT)  |
+-------------------------------------------------------+
```

**Tipo de diagrama:** C4 – **Container Diagram (Nivel 2)**  
**Propósito:** Mostrar los principales contenedores (apps, APIs, servicios, bases de datos) que conforman el sistema bancario y sus relaciones.

---

## **Bibliografía**

- Booch, G., Rumbaugh, J., & Jacobson, I. (1999). *The Unified Software Development Process*. Addison-Wesley.  
- Simon Brown. (2018). *The C4 Model for Visualising Software Architecture*. [https://c4model.com](https://c4model.com)  
- ISO/IEC 27001:2022 – *Information Security Management Systems*.  
- OWASP Foundation. (2024). *OWASP Top 10 – Security Risks for Web Applications*.  
- PCI Security Standards Council. (2023). *Payment Card Industry Data Security Standard (PCI DSS) v4.0*.  
- Sommerville, I. (2020). *Ingeniería del Software* (10ª ed.). Pearson Educación.

---

##  **Conclusión**

Una metodología arquitectónica aplicada a un sistema financiero debe equilibrar **seguridad, control y flexibilidad**. Su propósito no es solo desarrollar software, sino **garantizar confianza y estabilidad operativa** en un entorno donde la precisión y la trazabilidad son críticas.  
El modelo C4 apoya este enfoque al ofrecer **una visión estructurada de la arquitectura**, permitiendo documentar y comunicar claramente la complejidad del sistema desde el nivel contextual hasta el de código.

---

