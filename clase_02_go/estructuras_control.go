package main

//"strconv"
//"strings"
//"runtime"
//"math/rand"
//"os"
//"time"
import (
	"fmt"
)

func main() {
	fmt.Println("=== PANIC Y RECOVER ===")
	// PANIC BÁSICO
	demonstrarPanicBasico()
	// RECOVER PARA MANEJAR PANIC
	demonstrarRecover()
	// CASOS PRÁCTICOS
	demonstrarCasosPracticosPanicRecover()
}
func demonstrarPanicBasico() {
	fmt.Println("--- Panic básico ---")
	// defer se ejecuta incluso con panic
	defer fmt.Println("3. Defer ejecutándose durante panic")
	fmt.Println("1. Antes del panic")
	fmt.Println("2. Justo antes del panic")
	// Este panic terminará el programa si no se recupera
	// panic("¡Algo salió terriblemente mal!")
	fmt.Println("Esta línea nunca se ejecutaría")
}
func demonstrarRecover() {
	fmt.Println("\n--- Recover para manejar panic ---")
	// Función que puede hacer panic
	funcionPeligrosa := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf(" 🚨 Panic recuperado: %v\n", r)
				fmt.Println(" 🔄 Continuando ejecución normal...")
			}
		}()
		fmt.Println(" ⚙️ Iniciando operación peligrosa...")
		panic("¡Error crítico simulado!")

		fmt.Println(" Esta línea nunca se ejecutaría")
	}
	fmt.Println("1. Antes de función peligrosa")
	funcionPeligrosa()
	fmt.Println("2. Después de función peligrosa (recuperada)")
	fmt.Println("3. El programa continúa normalmente")
}
func demonstrarCasosPracticosPanicRecover() {
	fmt.Println("\n--- Casos prácticos ---")
	// 1. Servidor web que no debe caerse
	fmt.Println("1. Simulación de servidor web:")
	simularServidorWeb()
	// 2. Validación estricta
	fmt.Println("\n2. Validación con panic/recover:")
	testValidacion()
	// 3. Procesamiento de datos con recovery
	fmt.Println("\n3. Procesamiento de lote con recovery:")
	procesarLoteDatos()
	// 4. División segura
	fmt.Println("\n4. División segura:")
	testDivisionSegura()
}
func simularServidorWeb() {
	// Simular múltiples requests
	requests := []string{"GET /users", "POST /users", "GET /invalid", "DELETE /users/1"}
	for i, request := range requests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf(" 🚨 Request %d falló: %v\n", i+1,
						r)
					fmt.Println(" 📝 Logging error y continuando...")
				}
			}()
			fmt.Printf(" 📥 Procesando request %d: %s\n", i+1,
				request)
			// Simular error en request específico
			if request == "GET /invalid" {
				panic("endpoint no válido")
			}
			fmt.Printf(" ✅ Request %d completado exitosamente\n",
				i+1)
		}()
	}
	fmt.Println(" 🌐 Servidor continúa funcionando")
}
func testValidacion() {
	usuarios := []struct {
		Nombre string
		Edad   int
		Email  string
	}{
		{"Ana", 25, "ana@email.com"},
		{"", 30, "luis@email.com"},       // Error: nombre vacío
		{"María", -5, "maria@email.com"}, // Error: edad negativa
		{"Carlos", 35, "carlos@email.com"},
	}
	for i, usuario := range usuarios {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf(" ❌ Usuario %d inválido: %v\n",
						i+1, r)
				}
			}()
			validarUsuario(usuario.Nombre, usuario.Edad, usuario.Email)
			fmt.Printf(" ✅ Usuario %d válido: %s\n", i+1,
				usuario.Nombre)
		}()
	}
}
func validarUsuario(nombre string, edad int, email string) {
	if nombre == "" {
		panic("nombre no puede estar vacío")
	}
	if edad < 0 {
		panic("edad no puede ser negativa")
	}
	if email == "" {
		panic("email no puede estar vacío")
	}
}
func procesarLoteDatos() {
	datos := []interface{}{1, "texto", 3.14, []int{1, 2, 3}, nil, 42}
	resultados := make([]string, 0)
	for i, dato := range datos {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf(" ⚠️ Error procesando elemento %d: %v\n", i, r)
					resultados = append(resultados,
						fmt.Sprintf("ERROR_%d", i))
				}
			}()
			resultado := procesarDato(dato)
			resultados = append(resultados, resultado)
			fmt.Printf(" ✅ Elemento %d procesado: %s\n", i,
				resultado)
		}()
	}
	fmt.Printf(" 📊 Resultados finales: %v\n", resultados)
}
func procesarDato(dato interface{}) string {
	switch v := dato.(type) {
	case int:
		return fmt.Sprintf("INT_%d", v*2)
	case string:
		return fmt.Sprintf("STR_%s", v)
	case float64:
		return fmt.Sprintf("FLOAT_%.2f", v)
	case nil:
		panic("no se puede procesar nil")
	default:
		panic(fmt.Sprintf("tipo no soportado: %T", v))
	}
}
func testDivisionSegura() {
	operaciones := []struct {
		a, b float64
	}{
		{10, 2},
		{15, 3},
		{20, 0}, // División por cero
		{25, 5},
	}
	for i, op := range operaciones {
		resultado := divisionSegura(op.a, op.b)
		fmt.Printf(" %.1f ÷ %.1f = %s\n", op.a, op.b, resultado)
	}
}
func divisionSegura(a, b float64) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error en división: %v", r)
		}
	}()
	if b == 0 {
		panic("división por cero")
	}
	resultado := a / b
	return fmt.Sprintf("%.2f", resultado)
}

// Función utilitaria para demostrar panic con stack trace
func demonstrarStackTrace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recuperado: %v\n", r)
			// En aplicaciones reales, aquí podrías imprimir el stack trace completo
		}
	}()
	funcionNivel1()
}
func funcionNivel1() {
	funcionNivel2()
}
func funcionNivel2() {
	funcionNivel3()
}
func funcionNivel3() {
	panic("Error en función nivel 3")
}

/*
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









//SWITCH












func main() {
	fmt.Println("=== ESTRUCTURAS SWITCH ===")
	// SWITCH BÁSICO
	dia := time.Now().Weekday()
	switch dia {
	case time.Monday:
		fmt.Println("😴 Lunes - Inicio de semana")
	case time.Tuesday:
		fmt.Println("💪 Martes - A trabajar")
	case time.Wednesday:
		fmt.Println("🐪 Miércoles - Mitad de semana")
	case time.Thursday:
		fmt.Println("🚀 Jueves - Casi llegamos")
	case time.Friday:
		fmt.Println("🎉 Viernes - ¡Fin de semana próximo!")
	case time.Saturday, time.Sunday:
		fmt.Println("🏖 Fin de semana")

	default:
		fmt.Println("🤔 Día desconocido")
	}
	// SWITCH CON INICIALIZACIÓN
	switch mes := time.Now().Month(); mes {
	case time.December, time.January, time.February:
		fmt.Println("❄️ Época de verano (Hemisferio Sur)")
	case time.March, time.April, time.May:
		fmt.Println("🍂 Otoño")
	case time.June, time.July, time.August:
		fmt.Println("🧥 Invierno")
	case time.September, time.October, time.November:
		fmt.Println("🌸 Primavera")
	}
	// SWITCH SIN EXPRESIÓN (actúa como if-else-if)
	hora := time.Now().Hour()
	temperatura := 22.0
	switch {
	case hora < 6:
		fmt.Println("🌃 Madrugada")
	case hora < 12 && temperatura > 20:
		fmt.Println("🌞 Mañana agradable")
	case hora < 12:
		fmt.Println("🌅 Mañana fresca")
	case hora < 18 && temperatura > 25:
		fmt.Println("☀️ Tarde calurosa")
	case hora < 18:
		fmt.Println("🌤 Tarde normal")
	default:
		fmt.Println("🌙 Noche")
	}
	// SWITCH CON FALLTHROUGH (poco común)
	numero := 3
	switch numero {
	case 1:
		fmt.Print("uno")
		fallthrough
	case 2:
		fmt.Print("dos")
		fallthrough
	case 3:
		fmt.Print("tres")
		fallthrough
	case 4:
		fmt.Print("cuatro")
	}
	fmt.Println() // Nueva línea
	// SWITCH CON TYPE ASSERTION
	var interfaz interface{} = "texto"
	switch valor := interfaz.(type) {
	case string:
		fmt.Printf("Es string: '%s' (longitud: %d)\n", valor, len(valor))
	case int:
		fmt.Printf("Es entero: %d\n", valor)
	case float64:
		fmt.Printf("Es float: %.2f\n", valor)
	case bool:
		fmt.Printf("Es booleano: %t\n", valor)
	case nil:
		fmt.Println("Es nil")
	default:
		fmt.Printf("Tipo desconocido: %T\n", valor)
	}
	// CASOS PRÁCTICOS CON SWITCH
	demonstrarCasosPracticosSwitch()
}
func demonstrarCasosPracticosSwitch() {
	fmt.Println("\n--- Casos prácticos con switch ---")
	// 1. Procesamiento de códigos de estado HTTP
	statusCode := 404
	switch statusCode {
	case 200:
		fmt.Println("✅ OK")
	case 201:
		fmt.Println("✅ Creado")
	case 400:
		fmt.Println("❌ Petición incorrecta")
	case 401:
		fmt.Println("🔐 No autorizado")
	case 403:
		fmt.Println("🚫 Prohibido")
	case 404:
		fmt.Println("🔍 No encontrado")
	case 500:
		fmt.Println("💥 Error interno del servidor")
	default:
		if statusCode >= 200 && statusCode < 300 {
			fmt.Println("✅ Éxito")
		} else if statusCode >= 400 && statusCode < 500 {
			fmt.Println("❌ Error del cliente")
		} else if statusCode >= 500 {
			fmt.Println("💥 Error del servidor")
		} else {
			fmt.Printf("🤔 Código desconocido: %d\n", statusCode)
		}
	}
	// 2. Categorización de archivos por extensión
	filename := "documento.pdf"
	extension := filename[len(filename)-3:]
	switch extension {
	case "pdf":
		fmt.Println("📄 Documento PDF")
	case "doc", "docx":
		fmt.Println("📝 Documento de Word")
	case "xls", "xlsx":
		fmt.Println("📊 Hoja de cálculo")
	case "jpg", "png", "gif":
		fmt.Println("🖼 Imagen")
	case "mp4", "avi", "mov":
		fmt.Println("🎬 Video")
	case "mp3", "wav", "flac":
		fmt.Println("🎵 Audio")
	default:
		fmt.Printf("📁 Archivo de tipo: %s\n", extension)
	}
	// 3. Lógica de permisos por rol
	rol := "admin"
	accion := "delete_user"
	switch rol {
	case "super_admin":
		fmt.Println("🔑 Acceso total - Todas las acciones permitidas")
	case "admin":
		switch accion {
		case "create_user", "edit_user", "view_user":
			fmt.Println("✅ Acción permitida para admin")
		case "delete_user":
			fmt.Println("⚠️ Acción sensible - Requiere confirmación")
		default:
			fmt.Println("❌ Acción no permitida para admin")
		}
	case "moderator":
		switch accion {
		case "view_user", "edit_user":
			fmt.Println("✅ Acción permitida para moderador")
		default:
			fmt.Println("❌ Acción no permitida para moderador")
		}
	case "user":
		switch accion {
		case "view_user":
			fmt.Println("✅ Solo visualización permitida")
		default:
			fmt.Println("❌ Acción no permitida para usuario regular")
		}
	default:
		fmt.Println("❌ Rol no reconocido")
	}
	// 4. Procesamiento por sistema operativo
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("🐧 Configuración para Linux")
		configurarLinux()
	case "darwin":
		fmt.Println("🍎 Configuración para macOS")
		configurarMacOS()
	case "windows":
		fmt.Println("🪟 Configuración para Windows")
		configurarWindows()
	default:
		fmt.Printf("🤔 Sistema operativo no soportado: %s\n", os)
	}
	// 5. State machine simple
	estado := "inicio"
	evento := "login_exitoso"
	nuevoEstado := procesarEstado(estado, evento)
	fmt.Printf("Estado: %s -> Evento: %s -> Nuevo Estado: %s\n", estado,
		evento, nuevoEstado)
}
func configurarLinux() {
	fmt.Println(" - Configurando paths de Linux")
	fmt.Println(" - Estableciendo permisos UNIX")
}
func configurarMacOS() {
	fmt.Println(" - Configurando paths de macOS")
	fmt.Println(" - Configurando Keychain")
}
func configurarWindows() {
	fmt.Println(" - Configurando paths de Windows")
	fmt.Println(" - Configurando Registry")
}
func procesarEstado(estadoActual, evento string) string {
	switch estadoActual {
	case "inicio":
		switch evento {
		case "login_exitoso":
			return "autenticado"
		case "registro":
			return "registrando"
		default:
			return "inicio"
		}
	case "autenticado":
		switch evento {
		case "logout":
			return "inicio"
		case "timeout":
			return "sesion_expirada"
		default:
			return "autenticado"
		}
	case "sesion_expirada":
		switch evento {
		case "relogin":
			return "autenticado"
		case "timeout_final":
			return "inicio"
		default:
			return "sesion_expirada"
		}
	default:
		return "inicio"
	}
}







// ESTRUCTURA FOR









func main() {
	fmt.Println("=== ESTRUCTURAS FOR ===")
	// FOR CLÁSICO (C-style)
	fmt.Println("--- For clásico ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteración %d\n", i)
	}
	// FOR COMO WHILE
	fmt.Println("\n--- For como while ---")
	contador := 0
	for contador < 3 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}
	// FOR INFINITO
	fmt.Println("\n--- For infinito con break ---")
	i := 0
	for {
		if i >= 3 {
			break

		}
		fmt.Printf("Bucle infinito - iteración: %d\n", i)
		i++
	}
	// FOR CON MÚLTIPLES VARIABLES
	fmt.Println("\n--- For con múltiples variables ---")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d, suma=%d\n", i, j, i+j)
	}
	// FOR CON CONDICIONES COMPLEJAS
	fmt.Println("\n--- For con condiciones complejas ---")
	x, y := 1, 1
	for x < 100 && y < 100 {
		fmt.Printf("Fibonacci: x=%d, y=%d\n", x, y)
		x, y = y, x+y
	}
	// RANGE CON SLICES
	fmt.Println("\n--- Range con slices ---")
	frutas := []string{"manzana", "banana", "naranja", "uva"}
	// Con índice y valor
	for indice, fruta := range frutas {

		fmt.Printf("%d: %s\n", indice, fruta)
	}
	// Solo valores
	fmt.Println("Solo valores:")
	for _, fruta := range frutas {
		fmt.Printf("- %s\n", fruta)
	}
	// Solo índices
	fmt.Println("Solo índices:")
	for indice := range frutas {
		fmt.Printf("Índice: %d\n", indice)
	}
	// RANGE CON MAPS
	fmt.Println("\n--- Range con maps ---")
	edades := map[string]int{
		"Ana":   25,
		"Luis":  30,
		"María": 28,
	}
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)

	}
	// RANGE CON STRINGS
	fmt.Println("\n--- Range con strings ---")
	texto := "Hola 世界"
	// Por runes (caracteres Unicode)
	for i, caracter := range texto {
		fmt.Printf("Posición %d: %c (U+%04X)\n", i, caracter, caracter)
	}
	// CASOS PRÁCTICOS
	demonstrarCasosPracticosFor()
}
func demonstrarCasosPracticosFor() {
	fmt.Println("\n--- Casos prácticos con for ---")
	// 1. Procesamiento de lotes de datos
	fmt.Println("1. Procesamiento en lotes:")
	datos := make([]int, 100)
	for i := range datos {
		datos[i] = i + 1
	}

	tamañoLote := 10
	for i := 0; i < len(datos); i += tamañoLote {
		fin := i + tamañoLote
		if fin > len(datos) {
			fin = len(datos)
		}
		lote := datos[i:fin]
		fmt.Printf(" Procesando lote %d: %d elementos\n", i/tamañoLote+1,
			len(lote))
		// Simular procesamiento
		time.Sleep(50 * time.Millisecond)
	}
	// 2. Búsqueda con múltiples criterios
	fmt.Println("\n2. Búsqueda de usuarios:")
	usuarios := []struct {
		ID     int
		Nombre string
		Edad   int
		Activo bool
		Ciudad string
	}{
		{1, "Ana García", 25, true, "Lima"},
		{2, "Luis Martín", 30, false, "Cusco"},
		{3, "María López", 28, true, "Lima"},

		{4, "Carlos Ruiz", 35, true, "Arequipa"},
		{5, "Elena Torres", 29, true, "Lima"},
	}
	// Buscar usuarios activos de Lima mayores de 25
	fmt.Println("Usuarios activos de Lima > 25 años:")
	for _, usuario := range usuarios {
		if usuario.Activo && usuario.Ciudad == "Lima" && usuario.Edad > 25 {
			fmt.Printf(" - %s (%d años)\n", usuario.Nombre,
				usuario.Edad)
		}
	}
	// 3. Validación de datos con acumuladores
	fmt.Println("\n3. Validación de formulario:")
	campos := map[string]string{
		"nombre":   "Bryan Soberon",
		"email":    "bryansoberon@email.com",
		"telefono": "123456789",
		"edad":     "22",
		"ciudad":   "cix",
	}
	errores := make([]string, 0)
	camposValidos := 0

	for campo, valor := range campos {
		if valor == "" {
			errores = append(errores, fmt.Sprintf("Campo '%s' es requerido", campo))
		} else {
			camposValidos++
			fmt.Printf(" ✅ %s: %s\n", campo, valor)
		}
	}
	if len(errores) > 0 {
		fmt.Println(" Errores encontrados:")
		for _, error := range errores {
			fmt.Printf(" ❌ %s\n", error)
		}
	}
	fmt.Printf(" Campos válidos: %d/%d\n", camposValidos, len(campos))
	// 4. Generación de reportes con agrupación
	fmt.Println("\n4. Reporte de ventas por región:")
	ventas := []struct {
		Producto string
		Region   string
		Monto    float64
	}{
		{"Laptop", "Norte", 2500.00},
		{"Mouse", "Norte", 45.50},
		{"Laptop", "Sur", 2500.00},
		{"Teclado", "Centro", 120.00},
		{"Mouse", "Sur", 45.50},
		{"Laptop", "Centro", 2500.00},
	}
	ventasPorRegion := make(map[string]float64)
	contadorPorRegion := make(map[string]int)
	for _, venta := range ventas {
		ventasPorRegion[venta.Region] += venta.Monto
		contadorPorRegion[venta.Region]++
	}
	for region, total := range ventasPorRegion {
		promedio := total / float64(contadorPorRegion[region])
		fmt.Printf(" %s: $%.2f total (%d ventas, promedio: $%.2f)\n",
			region, total, contadorPorRegion[region], promedio)
	}
	// 5. Algoritmo de retry con backoff
	fmt.Println("\n5. Simulación de retry con backoff:")

	maxIntentos := 5
	for intento := 1; intento <= maxIntentos; intento++ {
		fmt.Printf(" Intento %d/%d", intento, maxIntentos)
		// Simular operación que puede fallar
		if rand.Float32() < 0.7 { // 70% probabilidad de fallo
			fmt.Println(" - ❌ Falló")
			if intento < maxIntentos {
				// Backoff exponencial
				delay := time.Duration(intento*intento) * 100 *
					time.Millisecond
				fmt.Printf(" Esperando %v antes del siguiente intento...\n", delay)
				time.Sleep(delay)
			}
		} else {
			fmt.Println(" - ✅ Éxito")
			break
		}
	}
	// 6. Algoritmo de ordenamiento burbuja
	fmt.Println("\n6. Ordenamiento burbuja:")
	numeros := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Printf(" Array original: %v\n", numeros)
	n := len(numeros)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
			}
		}
	}
	fmt.Printf(" Array ordenado: %v\n", numeros)
}





// Range Loop in Golang

func main() {
	opiniones := []string{
		"El servicio fue bueno y rápido",
		"El sistema es muy lento y malo",
		"Buen producto pero entrega lenta",
		"Rápido, eficiente y bueno",
		"Malo servicio, lento y sin soporte",
	}

	palabrasClave := []string{"bueno", "malo", "rápido", "lento"}
	// Inicializar el mapa de conteo
	conteo := make(map[string]int)
	for _, clave := range palabrasClave {
		conteo[clave] = 0
	}
	// Procesar opiniones
	for _, opinion := range opiniones {
		// Convertimos a minúsculas para uniformidad
		palabras := strings.Fields(strings.ToLower(opinion))
		for _, palabra := range palabras {
			// Limpiar comas o puntos si los hubiera (básico)
			palabra = strings.Trim(palabra, ".,;")
			// Verificamos si es una palabra clave
			if _, existe := conteo[palabra]; existe {
				conteo[palabra]++
			}
		}
	}
	// Mostrar resultados
	fmt.Println("📊 Conteo de palabras clave:")
	for palabra, cantidad := range conteo {
		fmt.Printf("- %s: %d veces\n", palabra, cantidad)
	}
}










//Control del Flujo

func main() {
	rand.Seed(time.Now().UnixNano())
	secreto := rand.Intn(10) + 1
	var input string
	jugadores := []string{"Jugador 1", "Jugador 2"}
	intentosMax := 3
	intentos := map[string]int{"Jugador 1": 0, "Jugador 2": 0}
	fmt.Println("🎮 ¡Bienvenidos al juego de adivinanza!")
	fmt.Println("🔢 Adivina el número secreto entre 1 y 10. Escribe 'salir' para salir.")

JUEGO:
	for {
		for _, jugador := range jugadores {
			if intentos[jugador] >= intentosMax {
				continue
			}
			fmt.Printf("👉 %s, intento %d: ", jugador,
				intentos[jugador]+1)
			fmt.Scanln(&input)
			if strings.ToLower(input) == "salir" {
				fmt.Println("🚪 El juego ha sido cancelado por el usuario.")
				goto FIN
			}
			var guess int
			_, err := fmt.Sscanf(input, "%d", &guess)
			if err != nil {
				fmt.Println("❌ Entrada no válida. Escribe un número.")
				continue
			}
			if guess%2 == 0 {
				fmt.Println("⚠️ Los pares no traen suerte. Intenta con otro número impar.")

				continue
			}
			intentos[jugador]++
			if guess == secreto {
				fmt.Printf("🎉 ¡%s adivinó el número secreto! Era %d🎯\n", jugador, secreto)
				break JUEGO
			} else {
				fmt.Println("❌ Incorrecto. Sigue intentando.")
			}
		}
		// Verifica si ambos jugadores agotaron sus intentos
		if intentos["Jugador 1"] >= intentosMax && intentos["Jugador 2"] >= intentosMax {
			fmt.Println("😢 Ambos jugadores agotaron sus intentos.")
			break
		}
	}
FIN:
	fmt.Println("🎯 Fin del juego. El número secreto era:", secreto)
}










// Defer


func main() {
	fmt.Println("=== DEFER ===")

	// DEFER BÁSICO
	demonstrarDeferBasico()

	// MÚLTIPLES DEFERS
	demonstrarMultiplesDefers()

	// DEFER CON VALORES
	demonstrarDeferConValores()

	// CASOS PRÁCTICOS
	demonstrarCasosPracticosDefer()
}
func demonstrarDeferBasico() {
	fmt.Println("--- Defer básico ---")

	fmt.Println("1. Inicio de función")

	defer fmt.Println("4. Este mensaje se ejecuta al final (defer)")

	fmt.Println("2. En medio de función")
	fmt.Println("3. Antes del return")
	// El defer se ejecuta aquí automáticamente
}
func demonstrarMultiplesDefers() {
	fmt.Println("\n--- Múltiples defers (LIFO - Last In, First Out) ---")

	defer fmt.Println("🥉 Tercer defer (se ejecuta primero)")
	defer fmt.Println("🥈 Segundo defer (se ejecuta segundo)")
	defer fmt.Println("🥇 Primer defer (se ejecuta último)")

	fmt.Println("Código normal ejecutándose...")
}
func demonstrarDeferConValores() {
	fmt.Println("\n--- Defer con valores capturados ---")

	x := 10
	defer fmt.Printf("Valor de x en defer: %d (capturado al definir defer)\n",
		x)

	x = 20
	fmt.Printf("Valor actual de x: %d\n", x)

	// El defer usará x=10

	// Para usar el valor actual, usar función anónima
	defer func() {
		fmt.Printf("Valor actual de x en defer con closure: %d\n", x)
	}()

	x = 30
	fmt.Printf("Valor final de x: %d\n", x)
}
func demonstrarCasosPracticosDefer() {
	fmt.Println("\n--- Casos prácticos con defer ---")

	// 1. Manejo de archivos
	fmt.Println("1. Manejo de archivos:")
	manejarArchivo()

	// 2. Medición de tiempo
	fmt.Println("\n2. Medición de tiempo:")
	medirTiempoEjecucion()

	// 3. Cleanup de recursos
	fmt.Println("\n3. Cleanup de recursos:")
	simularConexionDB()

	// 4. Logging de entrada y salida

	fmt.Println("\n4. Logging:")
	funcionConLogging("parametro_importante")

	// 5. Mutex unlocking
	fmt.Println("\n5. Manejo de mutex:")
	simularMutex()
}
func manejarArchivo() {
	// Simular apertura de archivo
	fmt.Println(" 📂 Abriendo archivo...")

	// defer se ejecuta incluso si hay error
	defer fmt.Println(" 🔒 Cerrando archivo (defer)")

	// Simular trabajo con archivo
	fmt.Println(" 📝 Escribiendo datos...")
	fmt.Println(" 📖 Leyendo datos...")

	// Si hubiera un error aquí, defer aún se ejecutaría
	// return // El defer se ejecuta antes del return
}
func medirTiempoEjecucion() {
	inicio := time.Now()
	defer func() {
		duracion := time.Since(inicio)
		fmt.Printf(" ⏱️ Función tardó: %v\n", duracion)
	}()

	fmt.Println(" 🔄 Iniciando operación costosa...")
	time.Sleep(100 * time.Millisecond) // Simular trabajo
	fmt.Println(" ✅ Operación completada")
}
func simularConexionDB() {
	fmt.Println(" 🔌 Conectando a base de datos...")

	defer fmt.Println(" 🔌 Desconectando de base de datos (defer)")

	// Simular múltiples operaciones
	fmt.Println(" 📊 Ejecutando query 1...")
	fmt.Println(" 📊 Ejecutando query 2...")
	fmt.Println(" 📊 Ejecutando query 3...")
}
func funcionConLogging(parametro string) {
	fmt.Printf(" 📥 ENTRADA: funcionConLogging(%s)\n", parametro)
	defer fmt.Println(" 📤 SALIDA: funcionConLogging")

	// Lógica de la función
	fmt.Println(" ⚙️ Procesando lógica de negocio...")

	if parametro == "error" {
		fmt.Println(" ❌ Error simulado")
		return // defer aún se ejecuta
	}

	fmt.Println(" ✅ Procesamiento exitoso")
}
func simularMutex() {
	fmt.Println(" 🔐 Adquiriendo lock...")

	defer fmt.Println(" 🔓 Liberando lock (defer)")

	// Simular trabajo en sección crítica
	fmt.Println(" ⚙️ Trabajando en sección crítica...")
	time.Sleep(50 * time.Millisecond)
}





//panic Y RECOVER

func causaPanico() {
	fmt.Println("Antes del pánico 😱")
	panic("¡Algo salió mal!")
	fmt.Println("Después del pánico ❌") // no se ejecuta
}
func main() {
	causaPanico()
	fmt.Println("Esto nunca se ejecuta ❌")
}

// RECOVER

func protegido() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🧯 Se recuperó del pánico:", r)
		}
	}()
	fmt.Println("Ejecutando función protegida")
	panic("🔥 Error inesperado")
}
func main() {
	protegido()
	fmt.Println("✅ El programa continúa después del recover")
}


//otro ejercicio para lograr entender mejor el uso de panic y recover















*/
