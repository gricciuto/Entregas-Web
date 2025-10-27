# 1. Obtener todos los usuarios
GET http://localhost:8080/usuarios
HTTP 200

# 2. Crear un usuario
POST http://localhost:8080/usuarios
Content-Type: application/json

{
    "Nombre": "Roman",
    "Email": "RomanElOriginalContrataciones@hotmail.com",
    "Contraseña": "PepeElSapo"
    }
HTTP 200
[Captures]
id: jsonpath "$.id_usuario"

# 3. Obtener un usuario particular por ID
GET http://localhost:8080/usuario/{{id}}
HTTP 200

# 4. Crear un entrenamiento
POST http://localhost:8080/entrenamientos
Content-Type: application/json

{
    "id_usuario"        : {{id}},
    "Fecha"             : "2025-10-20T14:30:00Z",
    "Tipo"              : "Carrera",
    "Distancia"         : 5.6,
    "Tiempo"            : 3

}
HTTP 200
[Captures]
id_entrenamiento: jsonpath "$.id_entrenamiento"

# 5. Ver entrenamientos
GET http://localhost:8080/entrenamientos
HTTP 200

# 6. Modificar un entrenamiento
UPDATE http://localhost:8080/entrenamiento/{{id_entrenamiento}}
Content-Type: application/json

{
    "Tipo" : "Caminata"
}
HTTP 200

# 7. Borrar un entrenamiento
DELETE http://localhost:8080/entrenamiento/{{id_entrenamiento}}

# 4. Borrar un usuario
DELETE http://localhost:8080/usuario/{{id}}
HTTP 200
