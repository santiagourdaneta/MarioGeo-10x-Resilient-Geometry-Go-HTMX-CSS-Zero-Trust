package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)

type SystemStatus struct {
	Uptime        string `json:"uptime"`
	GoVersion     string `json:"go_version"`
	MemoryAlloc   uint64 `json:"memory_alloc_mb"`
	NumGoroutines int    `json:"num_goroutines"`
	ApiHealth     string `json:"api_health"`
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

	// 3. Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 4. Envolver con Middlewares (ORDEN: Timeouts -> Logs)
	// Esto crea la "cebolla" de procesamiento
	finalHandler := withTimeouts(loggerMiddleware(mux))

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

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
		http.Error(w, "No se encontró static/index.html", 500)
		return
	}
	t.Execute(w, nil)
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
	json.NewEncoder(w).Encode(status)
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("sides")
	sides, _ := strconv.Atoi(query)
	if sides < 3 {
		fmt.Fprint(w, "Mínimo 3 lados")
		return
	}
	fmt.Fprintf(w, "📐 Triángulos: <strong>%d</strong>", sides-2)
}