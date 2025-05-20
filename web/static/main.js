document.addEventListener("DOMContentLoaded", () => {
    // Cargar estudiantes y luego inicializar DataTable
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

            // Inicializar DataTables DESPUÉS de llenar la tabla
            $('#tabla-estudiantes').DataTable({
                pageLength: 10,
                lengthChange: false,
                language: {
                    url: "https://cdn.datatables.net/plug-ins/1.13.6/i18n/es-ES.json"
                }
            });
        })
        .catch(err => {
            console.error("Error al obtener estudiantes:", err);
        });

    // Evento de cambio de carrera
    document.getElementById("carrera").addEventListener("change", function () {
        const carrera = this.value;
        const campos = document.getElementById("campos-restantes");

        if (carrera) {
            campos.style.display = "block";
        } else {
            campos.style.display = "none";
        }

        const show = id => document.getElementById(id).closest(".mb-3").style.display = "block";
        const hide = id => {
            const el = document.getElementById(id);
            el.value = "0";
            el.closest(".mb-3").style.display = "none";
        };

        show("ingenieria_software");
        show("bases_datos");
        show("control_analogo");
        show("circuitos_digitales");

        if (carrera === "Ingeniería de Sistemas") {
            hide("control_analogo");
            hide("circuitos_digitales");
        } else if (carrera === "Ingeniería Electrónica") {
            hide("ingenieria_software");
            hide("bases_datos");
        }
    });

    // Envío de formulario
    document.getElementById("form-estudiante").addEventListener("submit", async (e) => {
        e.preventDefault();

        const getValue = id => parseFloat(document.getElementById(id).value) || 0;

        const estudiante = {
            carrera: document.getElementById("carrera").value,
            genero_accion: parseInt(document.getElementById("genero_accion").value),
            genero_ciencia_ficcion: parseInt(document.getElementById("genero_ciencia_ficcion").value),
            genero_comedia: parseInt(document.getElementById("genero_comedia").value),
            genero_terror: parseInt(document.getElementById("genero_terror").value),
            genero_documental: parseInt(document.getElementById("genero_documental").value),
            genero_romance: parseInt(document.getElementById("genero_romance").value),
            genero_musicales: parseInt(document.getElementById("genero_musicales").value),
            poo: getValue("poo"),
            calculo_multivariado: getValue("calculo_multivariado"),
            ctd: getValue("ctd"),
            ingenieria_software: getValue("ingenieria_software"),
            bases_datos: getValue("bases_datos"),
            control_analogo: getValue("control_analogo"),
            circuitos_digitales: getValue("circuitos_digitales")
        };

        const validNota = n => n >= 0 && n <= 5;
        const validGenero = g => g >= 0 && g <= 10;

        const notas = [
            estudiante.poo, estudiante.calculo_multivariado, estudiante.ctd,
            estudiante.ingenieria_software, estudiante.bases_datos,
            estudiante.control_analogo, estudiante.circuitos_digitales
        ];
        const generos = [
            estudiante.genero_accion, estudiante.genero_ciencia_ficcion,
            estudiante.genero_comedia, estudiante.genero_terror,
            estudiante.genero_documental, estudiante.genero_romance,
            estudiante.genero_musicales
        ];

        if (!notas.every(validNota) || !generos.every(validGenero)) {
            alert("Verifica que las notas estén entre 0-5 y los géneros entre 0-10.");
            return;
        }

        try {
            const res = await fetch("/api/estudiantes", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(estudiante)
            });

            if (res.ok) {
                alert("✅ Estudiante añadido correctamente");
                window.location.reload();
            } else {
                const msg = await res.text();
                alert("❌ Error al guardar: " + msg);
            }
        } catch (err) {
            alert("❌ Error de conexión con el servidor");
            console.error(err);
        }
    });

    // Reset modal al abrir
    document.getElementById("modalEstudiante").addEventListener("show.bs.modal", function () {
        document.getElementById("form-estudiante").reset();
        document.getElementById("campos-restantes").style.display = "none";
    });
});
