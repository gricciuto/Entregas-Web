// URL base de tu API REST en Go
const API_URL = "http://localhost:8080/entrenamientos";

// Cargar lista de entrenamientos al inicio
function cargarEntrenamientos() {
    fetch(API_URL)
        .then(response => response.json())
        .then(data => {
            const contenedor = document.getElementById('ver-entrenamientos');
            contenedor.innerHTML = '<h2>Ver entrenamientos</h2>'; // reinicia contenido

            data.forEach(entrenamiento => {
                const div = document.createElement('div');
                div.className = 'entrenamiento-item';
                div.innerHTML = `
                    <div class="datos-entrenamiento">
                        <h3>${entrenamiento.tipo}</h3>
                        <p><strong>Fecha:</strong> ${entrenamiento.fecha}</p>
                        <p><strong>Distancia:</strong> ${entrenamiento.distancia} km</p>
                        <p><strong>Tiempo:</strong> ${entrenamiento.tiempo} min</p>
                        <p><strong>Ritmo:</strong> ${entrenamiento.ritmo || '-'} min/km</p>
                        <p><strong>Notas:</strong> ${entrenamiento.notas || 'Sin notas'}</p>
                    </div>
                    <button class="btn-eliminar" data-id="${entrenamiento.id}">Eliminar</button>
                `;
                contenedor.appendChild(div);
            });
        })
        .catch(error => console.error('Error al cargar los entrenamientos:', error));
}

// Enviar nuevo entrenamiento (POST)
document.getElementById('crear-entrenamiento').addEventListener('submit', function(event) {
    event.preventDefault();

    const nuevoEntrenamiento = {
        fecha: document.getElementById('fecha-entrenamiento').value,
        tipo: document.getElementById('tipo-entrenamiento').value,
        distancia: parseFloat(document.getElementById('distancia-entrenamiento').value),
        tiempo: parseInt(document.getElementById('tiempo-entrenamiento').value),
        ritmo: parseFloat(document.getElementById('ritmo-entrenamiento').value) || null,
        notas: document.getElementById('notas-entrenamiento').value
    };

    fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nuevoEntrenamiento)
    })
        .then(response => {
            if (response.ok) {
                console.log('Entrenamiento creado');
                cargarEntrenamientos();
                event.target.reset();
            } else {
                console.error('Error al crear entrenamiento');
            }
        })
        .catch(error => console.error('Error en POST:', error));
});

// Delegación de eventos para eliminar (DELETE)
document.getElementById('ver-entrenamientos').addEventListener('click', function(event) {
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

// Ejecutar al cargar la página
window.addEventListener('DOMContentLoaded', cargarEntrenamientos);