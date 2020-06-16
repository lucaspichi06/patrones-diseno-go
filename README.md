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