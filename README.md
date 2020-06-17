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