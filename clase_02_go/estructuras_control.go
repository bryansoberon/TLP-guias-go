package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== ESTRUCTURAS IF/ELSE ===")

	// IF BÁSICO
	edad := 25

	if edad >= 18 {
		fmt.Println("✅ Mayor de edad")
	}

	// IF-ELSE
	temperatura := 22

	if temperatura > 25 {
		fmt.Println("🌡 Hace calor")
	} else {
		fmt.Println("🌡 Temperatura agradable")
	}

	// IF-ELSE-IF (cadena)
	puntuacion := 85
	if puntuacion >= 90 {
		fmt.Println("🏆 Excelente")
	} else if puntuacion >= 75 {
		fmt.Println("👍 Bueno")
	} else if puntuacion >= 60 {
		fmt.Println("😐 Regular")
	} else {
		fmt.Println("😞 Necesita mejorar")
	}

	// IF CON INICIALIZACIÓN (patrón muy común en Go)
	if hora := time.Now().Hour(); hora < 12 {
		fmt.Println("🌅 Buenos días")
	} else if hora < 18 {
		fmt.Println("☀️ Buenas tardes")
	} else {
		fmt.Println("🌙 Buenas noches")
	}

	// VERIFICACIÓN DE ERRORES (patrón idiomático)
	if numero, err := strconv.Atoi("123"); err != nil {
		fmt.Printf("❌ Error de conversión: %v\n", err)
	} else {
		fmt.Printf("✅ Número convertido: %d\n", numero)
	}
	// MÚLTIPLES CONDICIONES
	usuario := "admin"
	password := "secret123"

	if usuario == "admin" && password == "secret123" {
		fmt.Println("🔑 Acceso concedido")
	} else {
		fmt.Println("🚫 Acceso denegado")
	}

	// CONDICIONES COMPLEJAS
	estado := "activo"
	ultimoAcceso := time.Now().Add(-24 * time.Hour)

	if estado == "activo" && time.Since(ultimoAcceso) < 30*24*time.Hour {
		fmt.Println("👤 Usuario activo y reciente")
	} else if estado == "activo" {
		fmt.Println("⚠️ Usuario activo pero inactivo por tiempo")
	} else {
		fmt.Println("❌ Usuario inactivo")
	}

	// CASOS PRÁCTICOS
	demonstrarCasosPracticosIf()
}
func demonstrarCasosPracticosIf() {
	fmt.Println("\n--- Casos prácticos con if ---")

	// 1. Validación de entrada
	email := "usuario@dominio.com"

	if len(email) == 0 {
		fmt.Println("❌ Email vacío")
	} else if !strings.Contains(email, "@") {
		fmt.Println("❌ Email inválido: falta @")
	} else if !strings.Contains(email, ".") {
		fmt.Println("❌ Email inválido: falta dominio")
	} else {
		fmt.Println("✅ Email válido")
	}

	// 2. Categorización de rangos
	velocidad := 75 // km/h
	limite := 60

	if velocidad <= limite {
		fmt.Println("🚗 Velocidad normal")
	} else if velocidad <= limite+10 {
		fmt.Println("⚠️ Ligero exceso de velocidad")
	} else if velocidad <= limite+20 {
		fmt.Println("🚨 Exceso moderado - multa")
	} else {
		fmt.Println("🚔 Exceso grave - suspensión")
	}

	// 3. Lógica de negocio con múltiples factores
	edad := 25
	experiencia := 3 // años
	certificaciones := 2

	if edad >= 21 && experiencia >= 2 && certificaciones >= 1 {
		fmt.Println("✅ Candidato calificado para posición senior")
	} else if edad >= 18 && (experiencia >= 1 || certificaciones >= 1) {
		fmt.Println("✅ Candidato calificado para posición junior")
	} else if edad >= 18 {
		fmt.Println("⚠️ Candidato para posición de entrenamiento")
	} else {
		fmt.Println("❌ No cumple requisitos mínimos")
	}

	// 4. Manejo de casos especiales
	valor := 0.0
	if valor > 0 {
		fmt.Printf("Valor positivo: %.2f\n", valor)
	} else if valor < 0 {
		fmt.Printf("Valor negativo: %.2f\n", valor)
	} else {
		// Caso especial: exactamente cero
		fmt.Println("Valor es exactamente cero")
	}

	// 5. Verificación de recursos
	memoryUsage := 85.5 // porcentaje
	cpuUsage := 70.2
	diskUsage := 45.0

	alertLevel := "normal"

	if memoryUsage > 90 || cpuUsage > 90 || diskUsage > 95 {
		alertLevel = "crítico"
	} else if memoryUsage > 80 || cpuUsage > 80 || diskUsage > 85 {
		alertLevel = "warning"
	}

	switch alertLevel {
	case "crítico":
		fmt.Println("🚨 ALERTA CRÍTICA: Recursos del sistema agotados")
	case "warning":
		fmt.Println("⚠️ ADVERTENCIA: Alto uso de recursos")
	default:
		fmt.Println("✅ Recursos del sistema normales")
	}
}
