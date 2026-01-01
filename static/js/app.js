/* jshint esversion: 6 */

// Registro del Service Worker para PWA
if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/static/sw.js')
    .then(() => console.log('🚀 Service Worker registrado'))
    .catch(err => console.error('❌ Error al registrar SW:', err));
}

/**
 * Valida la entrada del usuario en la calculadora
 * @param {HTMLInputElement} input 
 */
function validateInput(input) {
  const resultBox = document.getElementById("res");
  if (input.value < 3 && input.value !== "") {
    resultBox.innerHTML = "<span style='color:#ff3e3e'>La geometría requiere al menos 3 lados.</span>";
  }
}

// Escuchar eventos de HTMX para telemetría
document.body.addEventListener('htmx:afterRequest', (evt) => {
  if (evt.detail.pathInfo.requestPath === "/api/status") {
    const metricsContainer = document.getElementById("system-metrics");
    
    if (evt.detail.xhr.status === 200) {
      try {
        const data = JSON.parse(evt.detail.xhr.responseText);
        metricsContainer.innerHTML = `
          <p style="margin:5px 0">⏱️ Uptime: ${data.uptime.split(".")[0]}s</p>
          <p style="margin:5px 0">🧠 RAM: ${data.memory_alloc_mb} MB</p>
          <p style="margin:5px 0">🟢 <span class="status-badge">${data.api_health}</span></p>
        `;
      } catch (e) {
        console.error("Error parseando JSON de status:", e);
      }
    } else {
      metricsContainer.innerHTML = `
        <p style="color: #ff3e3e">⚠️ Error de conexión (${evt.detail.xhr.status})</p>
      `;
    }
  }
});