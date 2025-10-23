# Entrega tp3.

# Inicio del Servidor
    sudo make start
## Correr pruebas HURL
*Tiene que estar corriendo el servidor*

    sudo make hurl-test



------------------


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

