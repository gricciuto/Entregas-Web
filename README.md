# Entrega TP3 y TP4.
## Para ir a la carpeta donde esta el proyecto
    cd tp3-web
Luego los siguientes comandos
# Inicio del Servidor
    sudo make start

## Pruebas HURL
    sudo make hurl-test
## Ingreso al portal
    http://localhost:8080
Puede crear usuarios con su email y contrasena y al presionar registrarse estos se almacenan en la base de datos siendo antes chequeados que sean correctos.

Una vez creados, se agregan a la lista que se muestra a la derecha, cada uno con un boton que lo elimina.

**ACLARACION:**
 La lista que muestra los usuarios en la proxima entrega sera eliminada, solo esta para mostrar los cambios de la base de datos.
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


### Main_Test.go

Codigo que corre pruebas en la base de datos agregando y borrando usuarios y entrenamientos. Lo incluimos en esta entrega debido a que no fue incluido en la entrega anterior.
### Main.go

Contiene el servidor que escucha en el puerto **8080**, ademas de los handlers para cada operacion de usuarios y de entrenamientos (CRUD).

    Metodos de la api:
     http://localhost:8080/usuarios -> GET Devuelve todos los usuarios
                                    -> POST Crea un nuevo usuario

     http://localhost:8080/usuario/id -> GET Devuelve el usuario con el id
                                      -> DELETE Borra el usuario con el id
     http://localhost:8080/entrenamientos -> GET Devuelve todos los entrenamientos
                                          -> POST Crea un entrenamiento
     http://localhost:8080/entrenamiento/id -> PUT Actualiza un entrenamiento
                                            -> DELETE Borra un entrenamiento
                                            -> GET Obtiene un entrenamiento por id

### Logs
En el archivo /logs/server.log se encuentran los logs de la ejecucion del servidor
### Logic
Se verifican en logic la correcta escritura de los usuarios y los entrenamientos

