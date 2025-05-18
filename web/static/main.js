document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/estudiantes")
        .then(res => res.json())
        .then(data => {
            const tbody = document.querySelector("#tabla-estudiantes tbody");
            tbody.innerHTML = "";

            data.forEach(est => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${est.id}</td>
                    <td>${est.carrera}</td>
                    <td>${est.genero_accion}</td>
                    <td>${est.genero_ciencia_ficcion}</td>
                    <td>${est.genero_comedia}</td>
                    <td>${est.genero_terror}</td>
                    <td>${est.genero_documental}</td>
                    <td>${est.genero_romance}</td>
                    <td>${est.genero_musicales}</td>
                    <td>${est.poo}</td>
                    <td>${est.calculo_multivariado}</td>
                    <td>${est.ctd}</td>
                    <td>${est.ingenieria_software}</td>
                    <td>${est.bases_datos}</td>
                    <td>${est.control_analogo}</td>
                    <td>${est.circuitos_digitales}</td>
                    <td>${est.promedio}</td>
                `;
                tbody.appendChild(row);
            });
        })
        .catch(err => {
            console.error("Error al obtener estudiantes:", err);
        });
});
