# 📐 MarioGeo: Sistema Geométrico 10x 

## 🌐 Live Demo

Puedes probar el sistema en tiempo real aquí: [https://mariogeo-10x-resilient-geometry-go-htmx.onrender.com/](https://mariogeo-10x-resilient-geometry-go-htmx.onrender.com/)

![CI](https://github.com/santiagourdaneta/MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust/actions/workflows/ci.yml/badge.svg)
![Status](https://img.shields.io/badge/Status-ALIVE-success?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Zero--Trust-red?style=for-the-badge)
![Tech](https://img.shields.io/badge/Tech-Go_%2B_HTMX-blue?style=for-the-badge)

**MarioGeo** es una plataforma de ingeniería geométrica de alto rendimiento. Diseñada para hardware restringido, utiliza una arquitectura de comunicación asíncrona y autocuración (Auto-healing) para garantizar disponibilidad constante.

## 🏗️ Principios de Diseño

El proyecto se rige bajo el principio de Separación de Responsabilidades (SoC), garantizando un código limpio, testeable y escalable:

Backend (Go): Lógica de negocio pura. Los cálculos geométricos están aislados en geometry.go, mientras que la infraestructura de red reside en main.go.

Frontend Modular: * index.html: Estructura semántica pura.

style.css: Estética validada por Stylelint (orden alfabético y fallbacks de compatibilidad).

app.js: Comportamiento y lógica de cliente desacoplada.

Observabilidad Nativa: Implementación de un Watchdog Sentinel en Bash que monitorea proactivamente la salud del proceso y el consumo de RAM.

Zero-Dependency Frontend: Uso de HTMX para mantener la interactividad sin la carga cognitiva ni el peso de los frameworks de JS modernos.

🛠️ Quality Assurance (QA)

Para mantener los estándares de 2026, el repositorio integra un pipeline de análisis estático:

Linting: Validado con Revive (Go), Stylelint (CSS) y JSHint (JS).

Security: Escaneo de vulnerabilidades en contenedores mediante el Dockerfile multi-stage.

Resiliencia: Capacidad de auto-recuperación ante fugas de memoria o caídas críticas.


PROS: 

No hay código espagueti: Cada tecnología está en su sitio.

Cinsciencia del rendimiento: Zero-Dependency.

DevOps: Observabilidad y Auto-recuperación.


## 🏗️ Arquitectura del Sistema

El sistema se basa en tres pilares de ingeniería:

1. **Backend Modular (Go):** Separación estricta entre el servidor HTTP (`main.go`) y la lógica de cálculo (`geometry.go`).
2. **Frontend Ultra-Light (HTMX):** Carga dinámica de componentes sin recargas de página, reduciendo el consumo de ancho de banda en un 80%.
3. **Defensa en Profundidad:** Validación triple (CSS/HTML5 en el Edge, AWK en el Linter y Go en el Core).

## 🛠️ Herramientas de Calidad (No-Node Stack)

Para mantener el sistema ligero, he eliminado la dependencia de Node.js, utilizando herramientas nativas de Unix/GNU:

- **Linter de Go:** `golangci-lint` para análisis estático profundo.
- **Linter de Frontend:** Scripts personalizados en `AWK` para validar integridad de etiquetas `<script>` y `<style>`.
- **Git Hooks:** - `pre-commit`: Formatea y limpia el código automáticamente.
  - `pre-push`: Bloquea la subida a GitHub si los tests unitarios fallan.

## 📊 Panel de Observabilidad

El sistema incluye telemetría en tiempo real:
- **Watchdog en Terminal:** Monitorea RAM y CPU cada 15s.
- **Dashboard Web:** Integración con `/api/status` para visualizar el estado del servidor directamente en la UI mediante HTMX.


## 📦 Instalación y Desarrollo Rápido

Si tienes instalado `make` y `Go`, el flujo es instantáneo. El uso de `make dev` simplifica la vida de cualquier otro desarrollador que vea este proyecto.

1. **Clonar y Preparar:**
   ```bash
   git clone [https://github.com/santiagourdaneta/MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust.git](https://github.com/santiagourdaneta/MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust.git)
   cd MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust

2. Comando Maestro (Dev Mode): Este comando limpia, ejecuta linters, corre tests y lanza el sistema de autocuración:

  make dev

3. Ejecución Manual:

  go run .

🧪 Testing

Garantizo la precisión matemática mediante tests unitarios:

  make test

Cubre: Lógica de polígonos, cálculo de triángulos y manejo de errores.

Ingeniería de Santiago Urdaneta | Zero-Trust Architecture | 2026
