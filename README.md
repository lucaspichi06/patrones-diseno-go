# Patrones de Diseño en GoLang

### *Implementación de patrones de diseño en GoLang*

* #### Singleton
* #### Factory
* #### Builder
    Permite separar la construcción de un objeto (complejo o no) de su representación con el fin de que el mismo proceso de creación pueda crear diferentes representaciones.
    * Product: objeto principal construido. Representa el objeto bajo construcción.
    * Builder: define la interface con los métodos que deberían cumplir los constructores.
    * Concrete Builder: implementa la itnerface builder para entregar el objeto concreto.
    * Director: construye el objeto utilizando la interface.
* #### Prototype
    Este patrón permite crear objetos a partir de instancias ya creadas.

    La idea es evitar el costo de creación de los objetos, utilizando un objeto creado previamente y clonándolo.

    Cuando hablamos del costo de creación de un objeto, hablamso de que para crearlo se pudo haber consultado varios elementos de una base de datos, consultando otros servicios, etc. Por lo que clonarlo es una idea genial para duplicar la información ya consultada.

    #### Actores
    * Prototipo: interface que deben implementar los objetos clonables.
    * Prototipo concreto: estructura o tipo que implementa el prototipo.
    * Cliente: crea el nuevo objeto solicitando al prototipo que se clone.
* #### Adapter
    Permite adaptar dos interfaces compatibles. Responsable de utilizar las funcionalidades de estructuras independientes o incompatibles.
    
    Nos permite hacer compatible componentes antiguos con sistemas nuevos.
* #### Facade (Fachada)
    Permite ocultar la complejidad de un sistema proveyendo una interfaz sencilla para el cliente.
    Unifica las diferentes interfaces del sistema en una sola, con esto la perspectiva del cliente es el uso de un sistema de fácil acceso.
    #### Ejemplo
    Haremos un proceso para registrar los comentarios de una persona en un post de un blog.

    Para realizar esto, el sistema debe ejecutar varios procesos:
    * Validar que el usuario esté logueado y activo.
    * Validar que el usuario tiene permiso de comentar.
    * Registrar el comentario.
    * Notificar al creador del post que llegó un comentario.

    Son demasiados los procesos para que el cliente los haga uno a uno, e incluso puede ser que se salte algunos. Por esta razón la fachada se encargará de enmascarar todo ese proceso en uno solo:
    * Registrar comentario.
* #### Bridge (Puente)
    Permite separar el comportamiento de la representación (abstracción de la implementación).

    Requiere de una clase abstracta y una interface, pero como en go no hay clases, se usan iterfaces para ambos objetivos.

    Esto permite combinar comportamiento con representaciones, utilizando menos cantidad de estructuras.
    #### Diferencias contra herencia (en cantidad de clases)
    | Representación | Implementación | Herencia | Bridge
    ---: | ---: | ---: | ---:
    2 | 2 | 4 | 4
    3 | 2 | 6 | 5
    4 | 4 | 16 | 8
    5 | 4 | 20 | 9