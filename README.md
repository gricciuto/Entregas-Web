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

## **Correcciones**
```
Nicolás Angladette
Ramiro Ortiz
Gino Ricciuto

La entrega está bien documentada y en general bien estructurada. 
Sin embargo, hay algunos puntos importantes a mejorar.

Primero, no deberían ejecutar los comandos con sudo. 
Es una mala práctica y en este caso no es necesario, 
ya que todo corre correctamente sin privilegios de superusuario. 
Tuve que modificar algunas líneas para evitar que me pida la contraseña de sudo.

Además, no deberían tener el mapeo de db_data en el contenedor, 
ya que eso contiene datos binarios de la base. 
La creación de la base debe ocurrir al iniciar el contenedor, 
no mediante archivos persistentes mapeados.

Los tests con HURL están bien implementados y funcionan correctamente. 
En cuanto a la estructura, faltó separar la funcionalidad de los 
usuarios en un handler propio, ya que actualmente toda la lógica se encuentra 
en el main con una validación mínima de integridad.

En el Makefile, pueden simplificar los comandos invocando directamente los 
targets en lugar de llamarlos explícitamente con $(MAKE). 
Esto hará el flujo más claro y estándar.

Respecto a la aplicación, solo pude crear usuarios, pero no cargar entrenamientos. 
El README tampoco explica cómo hacerlo, ni cómo acceder a esa sección una vez 
registrado. Para la próxima entrega, asegúrense de agregar la lógica de uso 
completa y aprovechar las herramientas vistas en los prácticos 5 y 6, 
incorporando las mejoras sugeridas.

Vienen bien, sigan ajustando estos detalles para lograr una entrega más completa.

Saludos,`