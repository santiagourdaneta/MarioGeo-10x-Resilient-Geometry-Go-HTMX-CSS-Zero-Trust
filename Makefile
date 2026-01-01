# 🛠️ Makefile Maestro MarioGeo

.PHONY: all clean lint test run dev

# Tarea por defecto
all: lint test build

# 1. Limpiar binarios antiguos
clean:
	@echo "🧹 Limpiando..."
	@rm -f main.exe main
	@go clean

# 2. Formatear y Linter (Go + Web básica)
lint:
	@echo "🔍 Analizando integridad del codigo..."
	@gofmt -w .
	@golangci-lint run ./...
	@echo "🔍 Validando Frontend..."
	# Validador mejorado: Cuenta aperturas y cierres de forma independiente
	@OPEN=$$(grep -c "<script" static/index.html); \
	 CLOSE=$$(grep -c "</script>" static/index.html); \
	 if [ $$OPEN -ne $$CLOSE ]; then echo "❌ Error: Etiquetas <script> ($$OPEN vs $$CLOSE)"; exit 1; fi
	@OPEN_S=$$(grep -c "<style" static/index.html); \
	 CLOSE_S=$$(grep -c "</style>" static/index.html); \
	 if [ $$OPEN_S -ne $$CLOSE_S ]; then echo "❌ Error: Etiquetas <style> ($$OPEN_S vs $$CLOSE_S)"; exit 1; fi
	@echo "✅ Frontend validado con exito."

# 3. Ejecutar Tests
test:
	@echo "🧪 Corriendo Tests..."
	@go test -v ./...

# 4. Compilar
build:
	@echo "⚙️ Compilando..."
	@go build -o main.exe .

# 5. EL COMANDO MAESTRO: Limpiar, Testear y Watchdog
dev: clean lint test
	@echo "🚀 Lanzando Watchdog de salud y memoria..."
	@chmod +x watchdog.sh
	@./watchdog.sh