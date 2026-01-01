# 📐 MarioGeo: Sistema Geométrico 10x

![Status](https://img.shields.io/badge/Status-ALIVE-success?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Zero--Trust-red?style=for-the-badge)
![Tech](https://img.shields.io/badge/Tech-Go_%2B_HTMX-blue?style=for-the-badge)
![CI](https://github.com/santiagourdaneta/MarioGeo-10x-Resilient-Geometry-Go-HTMX-CSS-Zero-Trust/actions/workflows/ci.yml/badge.svg)

**MarioGeo** es una plataforma de cálculo geométrico que aplica principios avanzados de ingeniería de software para garantizar la máxima disponibilidad y seguridad con el mínimo consumo de recursos.

## 🚀 Características Principales

- **Arquitectura Async-First:** Comunicación eficiente entre el frontend y backend mediante HTMX, evitando recargas innecesarias.
- **Zero-Trust & Least Privilege:** Implementación de seguridad defensiva, validaciones en el borde (Edge) y reducción de superficie de ataque.
- **Auto-Healing Watchdog:** Sistema vigilante en Bash que monitorea la salud y el consumo de memoria, reiniciando el servicio automáticamente ante anomalías.
- **Resiliencia & Backpressure:** Control de timeouts y manejo de carga para evitar el colapso del sistema en hardware de bajos recursos.

## 🛠️ Stack Tecnológico

- **Backend:** Go (Net/HTTP, Goroutines, Timeouts Middleware).
- **Frontend:** HTML5, CSS3 (Bento Grid, Scroll-driven animations, 3D Transforms), HTMX.
- **DevOps:** Bash Watchdog, Health-check endpoints.

## 📦 Instalación y Uso

1. **Clonar el repositorio:**
   ```bash
   git clone [https://github.com/tu-usuario/mario-geo-landing.git](https://github.com/tu-usuario/mario-geo-landing.git)
   cd mario-geo-landing

2. Configurar variables de entorno: Crea un archivo .env basado en el entorno de desarrollo:

   PORT=8080
   DEBUG=true

3. Ejecutar con Auto-Healing:

   chmod +x watchdog.sh
   ./watchdog.sh

4. Acceso: Abre tu navegador en http://localhost:8080

📊 SLOs (Service Level Objectives)

Disponibilidad: 99.9% mediante el watchdog de autocuración.

Latencia: < 200ms en el procesamiento geométrico.

Memoria: Límite estricto de 100MB para entornos restringidos.

Diseñado con ❤️ por Santiago Urdaneta | 2025