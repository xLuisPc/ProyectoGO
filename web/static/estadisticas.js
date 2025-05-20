let chart = null;
let ultimaPrediccion = null;

async function cargarClusters() {
    const genero = document.getElementById("genero").value;
    const k = document.getElementById("k").value;
    const res = await fetch(`/api/estadisticas?genero=${genero}&k=${k}`);
    return await res.json();
}

function graficar(clusters, genero) {
    const colors = ['red', 'blue', 'green', 'purple', 'orange', 'cyan'];
    const datasets = clusters.map((cluster, i) => ({
        label: `Cluster ${i + 1}`,
        data: cluster.Personas.map(p => ({
            x: p[genero],
            y: p.promedio
        })),
        backgroundColor: colors[i % colors.length]
    }));

    if (chart) chart.destroy();

    chart = new Chart(document.getElementById("clusterChart"), {
        type: 'scatter',
        data: { datasets },
        options: {
            scales: {
                x: {
                    title: { display: true, text: genero.replace("genero_", "").toUpperCase() },
                    beginAtZero: true
                },
                y: {
                    title: { display: true, text: "Promedio" },
                    beginAtZero: true,
                    max: 5
                }
            },
            plugins: {
                legend: { position: 'bottom' }
            }
        }
    });
}

function generarResumen(clusters, genero) {
    const tbody = document.querySelector("#tablaResumen tbody");
    tbody.innerHTML = "";

    clusters.forEach((cluster, i) => {
        const n = cluster.Personas.length || 1;
        const promedio = campo => cluster.Personas.reduce((sum, p) => sum + p[campo], 0) / n;

        const fila = document.createElement("tr");
        fila.innerHTML = `
            <td><b>${i + 1}</b></td>
            <td>${n}</td>
            <td>${promedio(genero).toFixed(2)}</td>
            <td>${promedio("poo").toFixed(2)}</td>
            <td>${promedio("ctd").toFixed(2)}</td>
            <td>${promedio("calculo_multivariado").toFixed(2)}</td>
            <td>${promedio("ingenieria_software").toFixed(2)}</td>
            <td>${promedio("bases_datos").toFixed(2)}</td>
            <td>${promedio("control_analogo").toFixed(2)}</td>
            <td>${promedio("circuitos_digitales").toFixed(2)}</td>
            <td>${promedio("promedio").toFixed(2)}</td>
        `;
        tbody.appendChild(fila);
    });
}

async function actualizar() {
    const genero = document.getElementById("genero").value;
    const clusters = await cargarClusters();
    graficar(clusters, genero);
    generarResumen(clusters, genero);

    if (ultimaPrediccion) {
        await predecirConDatos(ultimaPrediccion);
    }
}

async function predecirCluster() {
    const datos = {
        carrera: document.getElementById("carreraPrediccion").value,
        genero_accion: parseInt(document.getElementById("genero_accion").value),
        genero_ciencia_ficcion: parseInt(document.getElementById("genero_ciencia_ficcion").value),
        genero_comedia: parseInt(document.getElementById("genero_comedia").value),
        genero_terror: parseInt(document.getElementById("genero_terror").value),
        genero_documental: parseInt(document.getElementById("genero_documental").value),
        genero_romance: parseInt(document.getElementById("genero_romance").value),
        genero_musicales: parseInt(document.getElementById("genero_musicales").value)
    };

    ultimaPrediccion = datos;
    await predecirConDatos(datos);
}

async function predecirConDatos(datos) {
    const res = await fetch("/api/prediccion", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(datos)
    });

    const resultado = await res.json();
    const clusterID = resultado.cluster;
    const promedio = resultado.promedio;

    const generoSeleccionado = document.getElementById("genero").value;
    const afinidad = datos[generoSeleccionado];

    chart.data.datasets = chart.data.datasets.filter(d => d.label !== "Predicción");

    chart.data.datasets.push({
        label: "Predicción",
        data: [{ x: afinidad, y: promedio }],
        backgroundColor: "black",
        pointRadius: 6,
        pointHoverRadius: 8
    });

    chart.update();

    document.getElementById("resultadoPrediccion").innerText =
        `🔍 Este perfil pertenece al Cluster ${clusterID + 1} | Estudiantes similares tienen un promedio de ${promedio.toFixed(2)}`;
}

function eliminarPrediccion() {
    ultimaPrediccion = null;
    chart.data.datasets = chart.data.datasets.filter(d => d.label !== "Predicción");
    chart.update();
    document.getElementById("resultadoPrediccion").innerText = "";
}

document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("genero").addEventListener("change", actualizar);
    document.getElementById("k").addEventListener("change", actualizar);
    document.querySelector("button[onclick='predecirCluster()']").addEventListener("click", predecirCluster);
    document.querySelector("button[onclick='eliminarPrediccion()']").addEventListener("click", eliminarPrediccion);
    actualizar();
});
