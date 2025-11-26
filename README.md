# Entrega TP6.
## Para ir a la carpeta donde esta el proyecto
    cd tp3-web
Luego los siguientes comandos
# Inicio del Servidor
    make start

## Ingreso al portal
    http://localhost:8080
Se ingresa con email y contrasena, si no esta registrado se puede registrar.

Al ingresar se redirige a la pagina de entrenamientos, que permite agregar y borrar entrenamientos.

Se puede cerrar la sesion pulsando el boton de la barra lateral que dice "Cerrar sesion" (lo unico que hace es redirigir a localhost:8080/login)


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
Por ejemplo en layout se tiene la estructura base del **html** y dos paginas *login page y entrenamiento page*.

Ahora en vez de enviar con **method="POST"** y **action="url"** lo que hicimos fue integrar htmx en la vista de login que sigue redirigiendo a la url de entrenamientos, pero si haces algun cambio en la lista de entrenamientos _(Ya sea borrar o agregar)_ esto se hace de manera dinamica sin recargar toda la pagina y sin usar javascript
Usamos hx-target y hx-swap para indicar que parte de la pagina especifica queremos actualizar.

Siguen estando en la carpeta static los archivos css que utilizan los templ y la imagen para el login.