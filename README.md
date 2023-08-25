# starshipCommsResolver API
starshipCommsResolver es un microservicio desarrollado en GO con el framework GIN, su desarrollo se realizo en una estructura hexagonal. Esta API proporciona una serie de ENDPOINT'S para desencriptar mensajes y triangular la ubiación de una nave respecto a 3 satelites del comando estelar acargo de Han Solo. 

## Documentación detallada en Swagger
En el repositorio se encuentra una carpeta llamada swagger, dentro hay un archivo en formato yaml que contiene la documentación en detalle de la API, para acceder se recomienda utilizar el siguiente editor: https://editor.swagger.io/

## Recibir información de los satelites y desencriptar
POST /topsecret/kenobi

Este endpoint recibe la información de los satelites incluida la distancia y el mensaje. Triangula la distancia respecto a la ubicación de los satelites y descifra el mensaje. 

### Ejemplo de solicitud:
json
Copy code
{
    "satellites": [
        {
            "name": "kenobi",
            "distance": 100.0,
            "message": [
                "este",
                "",
                "",
                "mensaje",
                ""
            ]
        },
        {
            "name": "skywalker",
            "distance": 115.5,
            "message": [
                "",
                "es",
                "",
                "",
                "secreto"
            ]
        },
        {
            "name": "sato",
            "distance": 142.7,
            "message": [
                "este",
                "",
                "un",
                "",
                ""
           ]
        }
    ]
}
## Obtener la información almacenada de Kenobi - Skywalker - Sato
GET /topsecret_split/kenobi

Este endpoint devuelve la información almacenada previamente para el satélite Kenobi, incluida la distancia y el mensaje.

## Almacenar la información de Kenobi - Skywalker - Sato
POST /topsecret_split/kenobi

Este endpoint permite almacenar la información del satélite Kenobi para su uso posterior en la resolución de la ubicación y el mensaje de la nave espacial desconocida.

### Ejemplo de solicitud:

json
Copy code
{ "distance": 100.0, "message": ["", "este", "", "mensaje", ""] }
Actualizar la información almacenada de Skywalker

## Ejecución local
Para ejecutar este servicio localmente, sigue los siguientes pasos:

- Clona este repositorio en tu máquina local.
- Asegúrate de tener Go instalado y configurado en tu sistema.
- Ejecuta go build para compilar el proyecto.
- Ejecuta ./nombre-de-tu-binario para iniciar el servidor.
- Accede a http://lcambiar para interactuar con el servicio localmente.

## Video Presentación


### Autor
#### Juan Jose Ochoa Ortiz

