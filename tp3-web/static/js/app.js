const API_Usuarios = "http://localhost:8080/usuarios"
console.log("Cargalo porque te mato")
function cargarUsuarios(){
    fetch(API_Usuarios)
        .then(response => response.json())
        .then(data => {
            const contenedor = document.getElementById('ver-usuarios');
            contenedor.innerHTML = '<h2>Ver Usuarios</h2>'; // reinicia contenido

            data.forEach(usuario => {
                const div = document.createElement('div');
                div.className = 'usuario-item';
                div.innerHTML = `
                    <div class="datos-usuario">
                        <h3>${usuario.nombre}</h3>
                        <p><strong>Contraseña:</strong> ${usuario.contraseña}</p>
                        <p><strong>Email:</strong> ${usuario.email}</p>

                    </div>`
                //<button class="btn-eliminar" data-id="${usuario.id}">Eliminar</button>
                ;
                contenedor.appendChild(div);
            });
        })
        .catch(error => console.error('Error al cargar los usuarios:', error));
}
// Enviar nuevo usuario (POST)
const boton = document.getElementById('btn_registrarse')
boton.addEventListener('click', function(event) {
    event.preventDefault();

    const nuevoUsuario = {
        nombre: document.getElementById("usuario").value,
        email: document.getElementById("email").value,
        contraseña: document.getElementById("contrasena").value
    };
    fetch(API_Usuarios, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nuevoUsuario)
    })
        .then(response => {
            if (response.ok) {
                console.log('Usuario creado');
                cargarUsuarios();
                event.target.reset();
            } else {
                console.error('Error al crear Usuario');
            }
        })
        .catch(error => console.error('Error en POST:', error));
});

// Delegación de eventos para eliminar (DELETE)
/*document.getElementById('ver-entrenamientos').addEventListener('click', function(event) {
    if (event.target.classList.contains('btn-eliminar')) {
        const id = event.target.getAttribute('data-id');

        fetch(`${API_URL}/${id}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (response.ok) {
                    console.log(`Entrenamiento ${id} eliminado`);
                    cargarEntrenamientos();
                } else {
                    console.error('Error al eliminar entrenamiento');
                }
            })
            .catch(error => console.error('Error en DELETE:', error));
    }
});
*/
// Ejecutar al cargar la página
window.addEventListener('DOMContentLoaded', cargarUsuarios);