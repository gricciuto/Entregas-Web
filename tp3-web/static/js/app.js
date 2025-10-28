const API_Usuarios = "http://localhost:8080/usuarios"
const API_Usuario = "http://localhost:8080/usuario/"
function cargarUsuarios(){
    fetch(API_Usuarios)
        .then(response => response.json())
        .then(data => {
            const contenedor = document.getElementById('ver-usuarios');

            if (data == null){
                contenedor.innerHTML = '<h2 style="color: white">No hay usuarios registrados</h2>';

            }
            else {
                contenedor.innerHTML = '<h2 style="padding: 1rem;background-color: rgba(255,255,255,40%); border-radius: 0.4rem">Ver Usuarios</h2>';
                data.forEach(usuario => {
                    const div = document.createElement('div');
                    div.className = 'usuario-item';
                    div.innerHTML = `
                    <div class="datos-usuario">
                        <h3>${usuario.nombre}</h3>
                        <p><strong>Contraseña:</strong> ${usuario.contraseña}</p>
                        <p><strong>Email:</strong> ${usuario.email}</p>
                        <button class="btn-borrar" data-id="${usuario.id_usuario}">Borrar</button>

                    </div>`
                    //<button class="btn-eliminar" data-id="${usuario.id}">Eliminar</button>
                    ;
                    contenedor.appendChild(div);
                });
            }
        })
        .catch(error => console.error('Error al cargar los usuarios:', error));
}
// Enviar nuevo usuario (POST)
const formu = document.getElementById('formulario')
formu.addEventListener('submit', function(event) {
    event.preventDefault();
    const nombreDoc = document.getElementById("usuario").value
    const emailDoc = document.getElementById("email").value
    const contrasenaDoc = document.getElementById("contrasena").value
    if (nombreDoc.trim() == '' || emailDoc.trim() == '' || contrasenaDoc.trim() == ''){
        alert('Te falto algun campito capito')
        return
    }

    const nuevoUsuario = {
        nombre: nombreDoc,
        email: emailDoc,
        contraseña: contrasenaDoc
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
                formu.reset()

            } else {
                console.error('Error al crear Usuario');
                alert('No se pudo crear el usuario')
                formu.reset()
            }
        })
        .catch(error => console.error('Error en POST:', error));
});

// Delegación de eventos para eliminar (DELETE)
const verusuarios = document.getElementById('ver-usuarios')
verusuarios.addEventListener('click', function(event){
    if (event.target.classList.contains('btn-borrar')){
        const id = event.target.getAttribute('data-id')
        fetch(`${API_Usuario}${id}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (response.ok) {
                    console.log(`Usuario ${id} eliminado`);
                    cargarUsuarios();
                } else {
                    console.error('Error al eliminar el usuario');
                }
            })
            .catch(error => console.error('Error en DELETE:', error));

    }
})


// Ejecutar al cargar la página
window.addEventListener('DOMContentLoaded', cargarUsuarios);