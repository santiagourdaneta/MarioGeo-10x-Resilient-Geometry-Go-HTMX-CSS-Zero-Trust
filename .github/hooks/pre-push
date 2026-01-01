#!/bin/sh
echo "🔍 Validando calidad de código antes del Push..."

# Ejecutar el formateador de Go
go fmt ./...

# Ejecutar el Linter de Go
if ! golangci-lint run; then
    echo "❌ Error: El Linter encontró problemas. Push cancelado."
    exit 1
fi

echo "✅ Código perfecto. Procediendo al envío..."