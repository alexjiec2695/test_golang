
## Tabla de contenido

* [Pre-requisitos ](#Pre-requisitos)
* [Instalación](#Instalación)


### Pre-requisitos 📋

* Para poder utilizar este aplicativo es necesario instalar [Go.](https://golang.org/doc/install)

* Instalar [MAKE.](https://www.gnu.org/software/make/) de forma global. 


### Instalación

* Clonar el repositorio

````
git clone https://github.com/alexjiec2695/test_golang.git
````

### Ejecución

* Tener [Docker.](https://www.docker.com/) instalado y corriendo.
* Desde la raiz del proyecto ejecutar el comando `Make run` 
* En caso de no tener ``Make`` instalado, desde la raiz del proyecto ejecutar el comando `docker-compose up` y luego navegar a la carpeta `cmd`, _ejemplo_ `cd cmd` y ejecutar el comando `go run main.go`

## Construido con 🛠️

* [Go](https://golang.org/) - Lenguaje de programación base del proyecto.
* [Fiber ](https://docs.gofiber.io/) - Librería web usada para la definición de los endpoints REST.
* [Testify](https://github.com/stretchr/testify) - Librería que permite realizar pruebas unitarias.
