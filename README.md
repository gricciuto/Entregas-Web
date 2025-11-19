# Entrega TP5.
## Para ir a la carpeta donde esta el proyecto
    cd tp3-web
Luego los siguientes comandos
# Inicio del Servidor
    make start

## Ingreso al portal
    http://localhost:8080
Por el momento tenemos implementado un login con un usuario.
-   Usuario:    profe
-   Email:      profe@gmail.com
  - Contraseña: 1234

Al ingresar se redirige a la pagina de entrenamientos, que permite agregar y borrar entrenamientos sin recargar toda la pagina.

------------------

## Logica del proyecto
El usuario se registra en la pagina y cada vez que sale a correr registra un nuevo entrenamiento que se va a almacenar en la base de datos. Para el calculo de las calorias quemadas se usa un peso fijo de 70kg

## Estructuras del proyecto
- Tabla Usuarios: Lleva registro de los usuarios de la plataforma
    
  - Contiene: 
    - ID Usuario
    - Usuario
    - Contrasena
    - Email
    
- Tabla Entrenamientos:

  - Contiene:
    
    - ID Entrenamiento
    - ID Usuario
    - Fecha
    - Tipo
    - Distancia
    - Tiempo
    - Ritmo Promedio
    - Calorias
    - Notas


### Main.go

Contiene el servidor que escucha en el puerto **8080**, y las funciones que al recibir /login o /entrenamientos o entrenamientos/delete llaman a los handlers que a su vez utilizan templ para interacturar con el front.

### Logs
En el archivo /logs/server.log se encuentran los logs de la ejecucion del servidor
### Views
Contiene los archivos templ. Cada uno contiene una funcion especifica con un trozo de html
Por ejemplo en layout se tiene la estructura base del **html** y dos paginas *login page y entrenamiento page* 
Siguen estando en la carpeta static los archivos css que utilizan los templ y la imagen para el login.