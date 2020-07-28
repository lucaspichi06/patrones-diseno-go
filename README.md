# Patrones de Diseño en GoLang

### *Implementación de patrones de diseño en GoLang*

## Creacionales
* #### Singleton
    Este patrón se utiliza cuando solo debe existir una única instancia de una estructura. Esta instancia única es llamada objeto singleton.
    Algunos de los casos en donde el objeto singleton es aplicable:
    * Instancia de base de datos: solo queremos crear una única instancia de objeto de base de datos y esa instancia se utilizará en toda la aplicación.
    * Instancia de usuario: solo se debe crear una única instancia del usuario y se debe usar en toda la aplicación.
* #### Factory
    Proporciona una forma de ocultar la lógica de creación de las instancias que se crean.
    El cliente solo interactúa con una estructura factory y le dice el tipo de instancias que deben crearse. La clase factory interactúa con las estructuras concretas correspondientes y devuelve la instancia correcta.
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
## Estructurales
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
* #### Composite
    Permite construir objetos complejos a partir de objetos más simples pero similares.

    Se debe utilizar la composición para crear estructuras de árbol y se debe usar recursividad para procesar la información de los objetos.
* #### Decorator
    Permite añadir funcionalidad a un objeto dinámicamente. Sin afectar el comportamiento de una clase/estructura completamente.

    Funciona como un wrapper. Envuelve el comportamiento de una función/estructura sin tener que modificarlo.
* #### Proxy
    Este patrón de diseño permite crear un objeto que se encargue de comunicarse con otros objetos para que el objeto cliente no tenga que comunicarse directamente con ellos.

    Nos permite simular la existencia de un objeto local cuando en realidad el objeto está remoto.

    También permite realizar procesos como caché de información.
## Comportamiento
* #### Observer
    Permite tener una relación de uno a muchos entre diferentes objetos con el objetivo de que podamos observar el comportamiento de un objeto y todos aquellos observadores que se encuentren registrados al comportamiento de dicho objeto sean notificados cuando ocurra un cambio.
* #### Chain of Responsibility
    Permite crear una cadena de manejadores de solicitudes. Para cada solicitud entrante, se pasa la misma a través de la cadena y cada uno de los controladores:
    * Procesa la solicitud u omite el procesamiento.
    * Decide si pasar la solicitud al siguiente controlador de la cadena.
* #### Strategy
    Permite cambiar el comportamiento de un objeto en tiempo de ejecución sin ningún cambio en la clase de este objeto.
    Cuándo utilizarlo:
    * Cuando un objeto necesita soportar diferentes comportamiento y es necesario cambiar ese comportamiento en tiempo de ejecución.
    * Cuando se tienen algoritmos diferentes que son similares y solo difieren en la forma en que ejecutan algún comportamiento.
