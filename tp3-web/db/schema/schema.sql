CREATE TABLE usuario (
    id_usuario SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    contraseña VARCHAR(255) NOT NULL
);

CREATE TABLE entrenamiento (
    id_entrenamiento SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL REFERENCES usuario(id_usuario) ON DELETE CASCADE,
    fecha DATE NOT NULL,
    tipo VARCHAR(50) NOT NULL,
    distancia DECIMAL(5,2) NOT NULL,
    tiempo INT NOT NULL,
    ritmo_promedio DECIMAL(5,2),
    calorias INT,
    notas TEXT
);