# Pregunta 3: ¿Qué riesgos podrían surgir si se elige una metodología ágil sin adaptarla al entorno bancario?

## **Resumen – Reflexión**

Adoptar una metodología ágil como **Scrum** o **Kanban** sin adaptarla al contexto bancario puede generar múltiples riesgos debido a las exigencias propias del sector financiero.  
El entorno bancario requiere **trazabilidad total**, **documentación formal**, **cumplimiento normativo (compliance)**, **auditorías frecuentes**, y un alto nivel de **seguridad en la información y los procesos transaccionales**.

Si se aplica una metodología ágil “pura” (sin control documental ni gestión formal de riesgos), pueden surgir los siguientes problemas:

- **Pérdida de trazabilidad:** no hay un registro claro que vincule los requisitos del negocio con los entregables técnicos.  
- **Falta de documentación auditada:** se incumplen normativas internas o externas (como PCI DSS, ISO 27001).  
- **Ausencia de controles de seguridad:** despliegues frecuentes sin validaciones formales pueden exponer vulnerabilidades.  
- **Problemas de cumplimiento:** sin evidencia documental, las auditorías regulatorias pueden fallar.  
- **Riesgos en la integración:** los cambios rápidos y no gobernados pueden afectar la estabilidad del sistema central.

Por eso, en sistemas bancarios se recomienda **adaptar las metodologías ágiles** incorporando mecanismos de control, revisión de código, auditoría y seguridad continua.  
El modelo ideal combina la **flexibilidad ágil** con la **rigurosidad de marcos tradicionales** como **RUP o CMMI**, generando un entorno seguro, auditable y adaptable.

---

## **Respectivos Diagramas - Planteamiento**

### **Propósito del diagrama C4 (Nivel de Contenedores)**

Este diagrama describe la estructura técnica de alto nivel del **sistema bancario**, mostrando cómo los distintos contenedores (frontend, backend, base de datos) se comunican y dónde pueden surgir riesgos si se aplican metodologías ágiles sin control formal.

#### **Interpretación**

- **Cliente Bancario:** utiliza los servicios digitales para realizar transacciones en línea.  
- **App Web Bancaria (React/Angular):** contenedor de presentación que gestiona la interacción con el usuario.  
- **API Backend (Java Spring Boot):** maneja la lógica de negocio y transacciones, pero puede sufrir vulnerabilidades si las iteraciones ágiles no se validan adecuadamente.  
- **Base de Datos (PostgreSQL):** almacena información sensible; las actualizaciones rápidas sin trazabilidad pueden comprometer su integridad.

El diagrama C4 revela que un sistema bancario requiere un **balance entre agilidad y gobernanza técnica**, donde cada contenedor debe estar cubierto por políticas de **seguridad**, **documentación** y **trazabilidad** antes del despliegue.

---

### **Propósito del Diagrama de Componentes**

Este diagrama se enfoca en el **nivel organizacional y metodológico**, mostrando cómo la metodología ágil sin adaptación interactúa con el entorno bancario y los actores involucrados, destacando los puntos donde surgen los riesgos operativos.

#### **Interpretación**

**Actores principales:**
- **Equipo de desarrollo:** aplica prácticas ágiles con documentación mínima.  
- **Entidad bancaria:** depende del cumplimiento normativo y auditorías.  
- **Cliente/Regulador:** supervisa el cumplimiento y exige evidencia documental.

**Componentes:**
- Las **iteraciones rápidas y cambios frecuentes** representan la agilidad, pero también la falta de control formal.  
- El **entorno bancario** incluye *seguridad de datos*, *cumplimiento normativo* y *auditoría y trazabilidad*, que son vulnerables si no existe coordinación con el proceso ágil.

El **diagrama UML** complementa al **C4** al mostrar la dimensión organizacional y de procesos, evidenciando que la metodología ágil sin adaptación no solo genera **fallas técnicas**, sino también **brechas regulatorias** y **falta de trazabilidad institucional**.

---

### **Enlace de diagramas correspondientes a la tercera pregunta**
📂 [Ver Diagramas en Google Drive](https://drive.google.com/drive/folders/1mlBv4EscclOWirn34rFSWFCh0gNrRiHs?usp=sharing)

---

## **Bibliografía (Norma APA 7ª edición)**

Fowler, M. (2010). *Continuous Integration*. martinfowler.com.  
Bass, L., Clements, P., & Kazman, R. (2012). *Software Architecture in Practice* (3ª ed.). Addison-Wesley.  
ISO/IEC 27001. (2013). *Information technology — Security techniques — Information security management systems — Requirements*. International Organization for Standardization.  
Kruchten, P. (2003). *The Rational Unified Process: An Introduction* (3ª ed.). Addison-Wesley.
