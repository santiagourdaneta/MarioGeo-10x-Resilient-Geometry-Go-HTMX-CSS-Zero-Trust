#!/bin/bash
# 🛡️ Watchdog 10x: Health & Memory Monitor
PORT=8080
MEM_LIMIT=100
COMMAND="go run main.go"

echo "🚀 Watchdog Senior iniciado. Límite de memoria: ${MEM_LIMIT}MB"

while true; do
  # 1. VERIFICAR SALUD (Protocolo HTTP)
  if ! curl -s --head http://localhost:$PORT/health | grep "200 OK" > /dev/null; then
    echo "⚠️ [$(date)] Servidor caído. Reiniciando..."
    taskkill //F //IM main.exe 2> /dev/null
    $COMMAND &
    sleep 5
  fi

  # 2. VERIFICAR MEMORIA (Prevención de fugas)
  # Buscamos el proceso main.exe y sumamos su consumo de memoria
  CURRENT_MEM=$(tasklist //FI "IMAGENAME eq main.exe" //FO CSV //NH | awk -F'","' '{gsub(/[. K]/,"",$5); print int($5/1024)}')

  if [ -n "$CURRENT_MEM" ] && [ "$CURRENT_MEM" -gt "$MEM_LIMIT" ]; then
    echo "🚨 [$(date)] Alerta: Memoria excedida ($CURRENT_MEM MB). Reiniciando por seguridad..."
    taskkill //F //IM main.exe
    $COMMAND &
    # Enviamos alerta manual a Sentry vía logs si fuera necesario
  fi

  sleep 15
done