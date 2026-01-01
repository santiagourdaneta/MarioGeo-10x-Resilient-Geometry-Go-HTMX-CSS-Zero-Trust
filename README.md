# 📐 MarioGeo: Sistema Geométrico 10x 

![CI](https://github.com/santiagourdaneta/MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust/actions/workflows/ci.yml/badge.svg)
![Status](https://img.shields.io/badge/Status-ALIVE-success?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Zero--Trust-red?style=for-the-badge)
![Tech](https://img.shields.io/badge/Tech-Go_%2B_HTMX-blue?style=for-the-badge)

**MarioGeo** es una plataforma de ingeniería geométrica de alto rendimiento. Diseñada para hardware restringido, utiliza una arquitectura de comunicación asíncrona y autocuración (Auto-healing) para garantizar disponibilidad constante.


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


🗺️ Roadmap de Evolución (Vision 2026)
El proyecto está diseñado para evolucionar hacia una plataforma de grado industrial siguiendo estos hitos:

🟢 Fase 1: Portabilidad Inmutable (Q1 2026)
Dockerización Ultra-Light: Creación de un Dockerfile multi-stage basado en scratch para generar imágenes de <15MB.

Orquestación Básica: Configuración de docker-compose para despliegue instantáneo con balanceo de carga.

🟡 Fase 2: Mobile-First & Offline (Q2 2026)
PWA (Progressive Web App): Implementación de Service Workers para que la calculadora funcione sin conexión a internet.

Manifest V3: Soporte para instalación nativa en dispositivos móviles y escritorio.

🟡 Fase 3: Gráficos de Alto Vértigo (Q3 2026)
Renderizado WebGL: Integración de Three.js o WebGL puro para visualizar los polígonos en 3D real mientras el usuario escribe.

Exportación CAD: Funcionalidad para descargar los cálculos en formatos vectoriales (.svg / .dxf).

🔴 Fase 4: Escalabilidad Global (Q4 2026)
Distribución en el Edge: Despliegue en Fly.io o Cloudflare Workers para latencias <50ms a nivel mundial.

API Pública Protegida: Implementación de Rate Limiting avanzado y claves API para uso por terceros.


Motivaciones del Roadmap de Evolución (Vision 2026):

Dockerización: Preparar el código para la nube moderna.

PWA: Experiencia del usuario (UX) y el acceso en áreas con mala conexión.

WebGL: Elevar el proyecto de "herramienta de texto" a "experiencia visual" (Vértigo 3D).


Este Roadmap es dinámico. Se priorizan las tareas que maximicen la resiliencia y minimicen el consumo de recursos.
