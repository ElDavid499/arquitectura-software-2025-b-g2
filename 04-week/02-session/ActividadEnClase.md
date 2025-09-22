**POO:** Paradigma de programación que modela el entorno mediante objetos que combinan estado y comportamiento. Permite representar conceptos del dominio de forma reutilizable y modular. Sus pilares son:

**Herencia**: Permite que una clase (subclase) herede atributos y métodos de otra clase (superclase), promoviendo la reutilización de código. Ejemplo: Una clase Usuario que hereda de una clase Persona, reutilizando atributos como nombre y documento.

**Polimorfismo**: Habilidad de un objeto para tomar diferentes formas, permitiendo que un mismo método se comporte de manera distinta según el contexto. Ejemplo: Un método autenticar() que puede implementarse de manera diferente en clases Usuario y Administrador.

**Encapsulamiento**: Restricción del acceso directo a los atributos de un objeto, proporcionando métodos para acceder y modificar su estado. Ejemplo: Usar métodos getSaldo() y setSaldo() en una clase Cuenta para proteger el atributo saldo.

**Abstracción**: Proceso de ocultar los detalles de implementación y mostrar solo las funcionalidades esenciales. Ejemplo: Una clase abstracta Figura con un método calcularArea() que es implementado por clases concretas como Circulo y Rectangulo.

plantuml
@startuml

@enduml

**UML:** Lenguaje de modelado visual estándar para especificar, visualizar, construir y documentar los artefactos de un sistema software. Se usa para representar diferentes vistas del sistema (estructura, comportamiento, interacciones).

**Diagrama de clases:** representa la estructura estática (clases, atributos, métodos y relaciones).

# Diagrama UML en PlantUML

```plantuml
@startuml
class Persona{
    - id: int
    - tipo_documento : String
    - documento : String
    - nombre : String
    + setNombre(nombre: String): void
    + getNombre(): String
}

class Usuario extends Persona{ 
    - id: int
    - nombre: String
    + autenticar(): boolean
}

class Cuenta {
    - numero: String
    - saldo: float
    + depositar(monto: float): void
    + retirar(monto: float): boolean
}

abstract class Operacion {
    - monto: float
    + ejecutar(cuenta: Cuenta): void
}

'Case 1: Dónde extends corrresponde al concepto herencia, también se conoce como la generalización.

'Case 2: Dónde entidad Persona el atributo nombre cuando se aplique los métodos set y get, esto corresponde a la encapsulación. setNombre(nombre: String): void, getNombre(): String

'Case 3: 

@enduml

Herencia 

- En class Usuario extends Persona se ve claramente el uso de herencia/generalización (Usuario hereda de Persona).

**Encapsulamiento**

- En Persona usas atributos privados (- nombre) y métodos setNombre y getNombre, lo cual es encapsulamiento: proteges el acceso directo a los atributos y usas métodos públicos para acceder/modificar.

Abstracción

- consiste en definir solo lo esencial de un objeto y ocultar los detalles innecesarios.

Se enfoca en qué hace un objeto, no en cómo lo hace.

Eso es abstracción → te muestran solo lo esencial

La clase Persona es genérica.

No se crea directamente, solo sirve de modelo base para otras clases como Usuario, Cliente, Empleado.

POLIMORFISMO

Qué es:“muchas formas”.
Es la capacidad de que mismo método tenga comportamientos diferentes según el objeto que lo use.

Se logra mediante herencia (sobrescritura) o interfaces.

El mensaje es el mismo (“x mensaje”), pero cada objeto responde distinto.
Eso es polimorfismo: mismo método → diferentes formas de hacerlo.