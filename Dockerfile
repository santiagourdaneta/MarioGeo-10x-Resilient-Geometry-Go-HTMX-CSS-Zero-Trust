# ETAPA 1: Compilación (El Constructor)
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
# Copiar archivos de dependencias
COPY go.mod ./
RUN go mod download
COPY go.sum ./
# Copiar el código fuente
COPY . .
# Compilar un binario estático (sin dependencias de C)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mariogeo .

# ETAPA 2: Producción (El Ejecutor Ultra-Ligero)
FROM scratch
WORKDIR /root/
# Copiamos el binario desde la etapa anterior
COPY --from=builder /app/mariogeo .
# Copiamos los archivos estáticos necesarios para la UI
COPY --from=builder /app/static ./static
# Puerto de salida
EXPOSE 8080
# Ejecución
CMD ["./mariogeo"]