package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

// SystemStatus representa las métricas de salud y rendimiento del servidor.
type SystemStatus struct {
	Uptime       string `json:"uptime"`
	GoVersion     string `json:"go_version"`
	MemoryAlloc  string `json:"memory_alloc_mb"`
	NumGoroutines int    `json:"num_goroutines"`
	APIHealth    string `json:"api_health"` 
}

var startTime = time.Now()

// Middleware de Logs
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		// Forzamos el log para que veas que funciona (aunque no haya variable DEBUG)
		log.Printf("METHOD=%s | PATH=%s | DURATION=%s | IP=%s",
			r.Method, r.URL.Path, time.Since(start), r.RemoteAddr)
	})
}

// Middleware de Timeouts
func withTimeouts(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		r = r.WithContext(ctx)
		done := make(chan bool)
		go func() {
			next.ServeHTTP(w, r)
			done <- true
		}()
		select {
		case <-done:
			return
		case <-ctx.Done():
			w.WriteHeader(http.StatusGatewayTimeout)
			fmt.Fprint(w, "Timeout del servidor")
		}
	})
}

func main() {
	_ = godotenv.Load()

	// 1. Crear el Multiplexor
	mux := http.NewServeMux()

	// 2. Definir Rutas
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/api/status", handleStatus)
	mux.HandleFunc("/api/calculate", handleCalculate)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") })
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) })

	mux.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
	    http.ServeFile(w, r, "static/manifest.json")
	})
	mux.HandleFunc("/sw.js", func(w http.ResponseWriter, r *http.Request) {
	    http.ServeFile(w, r, "static/sw.js")
	})

	// 3. Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 4. Envolver con Middlewares (ORDEN: Timeouts -> Logs)
	// Esto crea la "cebolla" de procesamiento
	finalHandler := withTimeouts(loggerMiddleware(mux))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto local
	}

	log.Printf("🚀 Sistema 10x corriendo en http://localhost:%s", port)
	
	// 5. Iniciar Servidor usando el handler envuelto
	err := http.ListenAndServe(":"+port, finalHandler)
	if err != nil {
		log.Fatal(err)
	}
}

// --- HANDLERS ---

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Error cargando template", 500)
		return
	}
	// Ejecutar una sola vez y capturar error
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error ejecutando template: %v", err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	status := SystemStatus{
		Uptime:        time.Since(startTime).String(),
		GoVersion:     runtime.Version(),
		MemoryAlloc:   m.Alloc / 1024 / 1024,
		NumGoroutines: runtime.NumGoroutine(),
		ApiHealth:     "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	// Encode escribe directamente en el ResponseWriter una sola vez
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {

	log.Printf("📥 Solicitud recibida: %s %s desde %s", r.Method, r.URL.Path, r.RemoteAddr)

	query := r.URL.Query().Get("sides")
	sides, _ := strconv.Atoi(query)

	if sides < 3 {
		fmt.Fprint(w, "La geometría requiere al menos 3 lados.")
		return
	}

	// Usamos la función que acabamos de crear
	nombre := getPolygonName(sides)
	triangulos := sides - 2

	fmt.Fprintf(w, "Este es un <strong>%s</strong>. <br> 📐 Se puede dividir en <strong>%d</strong> triángulos.", nombre, triangulos)
}
