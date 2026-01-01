#!/bin/bash
# ========================================================
# MARIO-GEO WATCHDOG 
# Optimizado para: Windows MINGW64 / Go 1.21+
# ========================================================

# Configuración
PORT=8080
BINARY="main.exe"
MEM_LIMIT=120  # Límite de seguridad en MB
LOG_FILE="sentinel.log"

# Colores para la terminal
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

clear
echo -e "${CYAN}🚀 Iniciando Sentinel Watchdog...${NC}"

# 1. Asegurar que el binario existe (Compilación fresca)
echo -e "${YELLOW}⚙️ Compilando binario de producción...${NC}"

if ! go build -o "$BINARY" .; then
    echo -e "${RED}❌ Error crítico: El código no compila.${NC}"
    exit 1
fi

while true; do
    # 2. Verificar si el proceso existe
    # Usamos tasklist y filtramos por el nombre exacto del binario
    PROC_DATA=$(tasklist //FI "IMAGENAME eq $BINARY" //NH //FO CSV 2>/dev/null | tr -d '"')
    
    if [[ $PROC_DATA == *$BINARY* ]]; then
        # Extraer memoria: El 5to campo es la RAM (ej: 15.420 K)
        RAW_MEM=$(echo "$PROC_DATA" | cut -d',' -f5 | sed 's/[^0-9]//g')
        MEM_MB=$((RAW_MEM / 1024))
        STATUS="${GREEN}RUNNING (ALIVE)${NC}"
    else
        MEM_MB=0
        STATUS="${RED}DOWN (RESTARTING)${NC}"
        # Reiniciar proceso en segundo plano
        ./$BINARY >> $LOG_FILE 2>&1 &
        echo -e "$(date) - [SENTINEL] Servidor reiniciado por caída." >> $LOG_FILE
        sleep 2
    fi

    # 3. Interfaz Visual (TUI)
    clear
    echo -e "${CYAN}========================================================${NC}"
    echo -e "   📊 ${CYAN}MARIO-GEO TELEMETRY${NC} (Refresco: 5s)             "
    echo -e "${CYAN}========================================================${NC}"
    echo -e "  ESTADO:    $STATUS"
    echo -e "  MEMORIA:   ${YELLOW}$MEM_MB MB${NC} / $MEM_LIMIT MB"
    echo -e "  PUERTO:    $PORT"
    echo -e "  LOGS:      tail -f $LOG_FILE"
    echo -e "${CYAN}--------------------------------------------------------${NC}"
    echo -e "${CYAN}========================================================${NC}"

    # 4. Guardian de Memoria (Anti-Leak)
    if [ "$MEM_MB" -gt "$MEM_LIMIT" ]; then
        echo -e "${RED}🚨 CRITICAL: Memory Leak detected ($MEM_MB MB).${NC}"
        echo -e "$(date) - [SENTINEL] Kill por exceso de memoria." >> $LOG_FILE
        taskkill //F //IM $BINARY > /dev/null 2>&1
    fi

    sleep 5
done