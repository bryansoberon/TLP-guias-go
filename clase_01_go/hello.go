package main

import (
	"fmt"

	//"math"
	//"strconv"
	//"unicode/utf8"
	//"unsafe"
)


/*"""

go build -o hello


# Formatear un archivo específico
go fmt main.go
# Formatear todos los archivos .go en el directorio actual
go fmt .
# Formatear recursivamente (todos los subdirectorios)

go fmt ./...
# Formatear un paquete específico
go fmt mipackage
# Formatear múltiples archivos
go fmt archivo1.go archivo2.go archivo3.go




package main
import "fmt"
func main() {
if true {
fmt.Println("Mal formateado")
}
for i := 0; i < 5; i++ {
fmt.Printf("Número: %d\n", i)
}
}



package main
import (
"fmt"
"time"
)
type Persona struct {
Nombre string
Edad int
Email string
}
func (p *Persona) Saludar() string {
return fmt.Sprintf("Hola, soy %s", p.Nombre)
}
func main() {
p := &Persona{
Nombre: "Juan",
Edad: 30,
Email: "juan@ejemplo.com",
}
fmt.Println(p.Saludar())
}





package main
import "fmt"
func main() {
numeros := []int{1, 2, 3, 4, 5}
mapa := map[string]int{"uno": 1, "dos": 2, "tres": 3}
matriz := [][]int{{1, 2}, {3, 4}, {5, 6}}
for _, num := range numeros {
if num%2 == 0 {
fmt.Printf("Par: %d\n", num)
} else {
fmt.Printf("Impar: %d\n", num)
}
}
}


go vet [archivos/paquetes]

# Analizar un archivo específico
go vet main.go
# Analizar el paquete actual
go vet .
# Analizar recursivamente todos los subpaquetes
go vet ./...
# Analizar un paquete específico
go vet mipackage

# Analizar con verbose (más información)
go vet -v ./...
# Mostrar todas las verificaciones disponibles
go vet help

func main() {
	// Go infiere el tipo basándose en el valor
	var x = 42      // int
	var y = 3.14    // float64
	var z = "hello" // string
	var w = true    // bool
	var v = 'A'     // rune (int32)

	fmt.Printf("x: %T = %v\n", x, x)
	fmt.Printf("y: %T = %v\n", y, y)
	fmt.Printf("z: %T = %v\n", z, z)
	fmt.Printf("w: %T = %v\n", w, w)
	fmt.Printf("v: %T = %v\n", v, v)
}

// ASCII es American Standar Code For information intechange

func main() {
	// Go infiere el tipo basándose en el valor
	var x = 42      // int
	var y = 3.14    // float64
	var z = "hello" // string
	var w = true    // bool
	var v = 'A'     // rune (int32)

	fmt.Printf("x: %T = %v\n", x, x)
	fmt.Printf("y: %T = %v\n", y, y)
	fmt.Printf("z: %T = %v\n", z, z)
	fmt.Printf("w: %T = %v\n", w, w)
	fmt.Printf("v: %T = %v\n", v, v)
	// Bloque anidado
if true {
// Variable local al bloque
var blockVar = "bloque"
// Puede acceder a variables de niveles superiores
fmt.Printf("Desde bloque - Global: %d, Local: %s, Bloque: %s\n",
globalCounter, localVar, blockVar)
// Shadowing (sombreado)
var localVar = "sombreada" // Nueva variable que sombrea la exterior fmt.Printf("Variable sombreada: %s\n", localVar)
}
// blockVar no existe aquí
// fmt.Println(blockVar) // ERROR: undefined
fmt.Printf("Después del bloque - Local: %s\n", localVar) //Valor original// Loop scope
for i := 0; i < 3; i++ {
var loopVar = fmt.Sprintf("iteración_%d", i)
fmt.Printf("Loop: i=%d, loopVar=%s\n", i, loopVar)
}
// i y loopVar no existen aquí
// fmt.Println(i) // ERROR: undefined
}

func otherFunction() {
// Puede acceder a variables globales
globalCounter++
fmt.Printf("Desde otra función - Global: %d\n",
globalCounter)
// No puede acceder a variables locales de main
// fmt.Println(localVar) // ERROR: undefined
}



//patrones de inicialización

func main() {
// Inicialización con valores calculados
var tiempo = time.Now()
var timestamp = tiempo.Unix()
var año = tiempo.Year()
// Inicialización con llamadas a funciones
var raizCuadrada = math.Sqrt(16)
var aleatorio = time.Now().Nanosecond() % 100
// Inicialización condicional
var mensaje string
if año%2 == 0 {
mensaje = "Año par"
} else {
mensaje = "Año impar"
}
// Inicialización lazy (cuando se necesite)
var configuracion map[string]string
if configuracion == nil {
configuracion = make(map[string]string)
configuracion["env"] = "development"
}
fmt.Printf("Tiempo: %v\n", tiempo)
fmt.Printf("Timestamp: %d, Año: %d\n", timestamp, año)
fmt.Printf("Raíz cuadrada: %.2f, Aleatorio: %d\n",
raizCuadrada, aleatorio)
fmt.Printf("Mensaje: %s\n", mensaje)
fmt.Printf("Configuración: %v\n", configuracion)
}



//Constantes
// Constantes a nivel de paquete
const CompanyName = "TechCorp" // Exportada
const version = "1.0.0" // No exportada
func main() {
// Constantes locales
const pi = 3.14159
const mensaje = "¡Hola, Go!"
const activo = true
// Go infiere el tipo de las constantes
fmt.Printf("Pi: %T = %v\n", pi, pi)
fmt.Printf("Mensaje: %T = %v\n", mensaje, mensaje)
fmt.Printf("Activo: %T = %v\n", activo, activo)
// Las constantes deben ser evaluables en tiempo de
compilación

const tiempoCompilacion = "Compilado el 2024"
// const tiempoEjecucion = time.Now() // ERROR: no es
constante
fmt.Printf("Constantes de empresa: %s v%s\n", CompanyName,
version)
fmt.Printf("Tiempo: %s\n", tiempoCompilacion)
}

//constantes Tipadas y no tipadas
func main() {
// Constantes no tipadas (untyped constants)
const a = 42 // Constante numérica no tipada
const b = 3.14 // Constante flotante no tipada
const c = "hello" // Constante string no tipada
// Constantes tipadas (typed constants)
const d int = 42
const e float64 = 3.14
const f string = "hello"
// Las constantes no tipadas pueden usarse con diferentes tipos compatibles
var x1 int = a
var x2 int64 = a
var x3 float32 = a
var x4 float64 = a
fmt.Printf("Constante no tipada 'a' usada como:\n")
fmt.Printf(" int: %T = %v\n", x1, x1)
fmt.Printf(" int64: %T = %v\n", x2, x2)
fmt.Printf(" float32: %T = %v\n", x3, x3)
fmt.Printf(" float64: %T = %v\n", x4, x4)
// Las constantes tipadas solo pueden usarse con su tipo
exacto
var y1 int = d
// var y2 int64 = d // ERROR: cannot use d (int) as int64
fmt.Printf("Constante tipada 'd': %T = %v\n", y1, y1)
// Precisión extendida en constantes no tipadas
const huge = 1e100
const tiny = 1e-100
fmt.Printf("Huge: %g, Tiny: %g\n", huge, tiny)
}

//Enumeraciones con iota
func main() {
// Enumeración básica con iota
const (
Lunes = iota // 0
Martes // 1
Miercoles // 2
Jueves // 3
Viernes // 4
Sabado // 5
Domingo // 6
)
fmt.Printf("Días de la semana:\n")
fmt.Printf("Lunes: %d, Martes: %d, Miércoles: %d\n", Lunes,
Martes, Miercoles)
// Enumeración con valores específicos
const (
StatusInactivo = iota + 1 // 1
StatusActivo // 2
StatusSuspendido // 3
StatusBloqueado // 4
)
fmt.Printf("Estados:\n")
fmt.Printf("Inactivo: %d, Activo: %d, Suspendido: %d, Bloqueado: %d\n", StatusInactivo, StatusActivo, StatusSuspendido,
StatusBloqueado)
// Enumeración con potencias de 2 (flags)
const (
Read = 1 << iota // 1 << 0 = 1
Write // 1 << 1 = 2
Execute // 1 << 2 = 4
)
fmt.Printf("Permisos (flags):\n")
fmt.Printf("Read: %d, Write: %d, Execute: %d\n", Read, Write,
Execute)
// Combinación de permisos
const ReadWrite = Read | Write // 3
const FullAccess = Read | Write | Execute // 7
fmt.Printf("Permisos combinados:\n")
fmt.Printf("ReadWrite: %d, FullAccess: %d\n", ReadWrite,
FullAccess)
// Enumeración con saltos
const (
Small = iota + 1 // 1
_ // 2 (skip)
_ // 3 (skip)
Medium // 4
_ // 5 (skip)
Large // 6
)
fmt.Printf("Tamaños:\n")
fmt.Printf("Small: %d, Medium: %d, Large: %d\n", Small,
Medium, Large)
// Enumeración con tipos personalizados
type Prioridad int
const (
Baja Prioridad = iota + 1
Media
Alta
Critica
)
var taskPriority Prioridad = Alta
fmt.Printf("Prioridad de tarea: %d\n", taskPriority)
}
// CONSTANTES COMPLEJAS
func main() {
	// Constantes matemáticas
	const (
		Pi    = 3.14159265358979323846
		E     = 2.71828182845904523536
		Phi   = 1.61803398874989484820 // Golden ratio
		Sqrt2 = 1.41421356237309504880
	)
	// Constantes derivadas
	const (
		CircleArea   = Pi * 5 * 5 // Área de círculo
		radio        = 5
		Tau          = 2 * Pi // Tau = 2π
		HalfPi       = Pi / 2
		DegreesToRad = Pi / 180
	)
	// Constantes con expresiones complejas
	const (
		KiB = 1024
		MiB = KiB * 1024
		GiB = MiB * 1024
		TiB = GiB * 1024
	)
	// Constantes de configuración
	const (
		MaxUsers       = 1000
		SessionTimeout = 30 * 60 // 30 minutos en segundos
		RetryAttempts  = 3
		BufferSize     = 8 * KiB
	)
	fmt.Printf("Constantes matemáticas:\n")
	fmt.Printf("π = %.10f, e = %.10f, φ = %.02f, √2 = %.4f\n", Pi, E, Phi, Sqrt2)
	fmt.Printf("Área círculo de radior=5: %.2f, τ = %.6f\n", CircleArea, Tau)
	fmt.Printf("\nConstantes de almacenamiento:\n")
	fmt.Printf("1 KiB = %d bytes\n", KiB)
	fmt.Printf("1 MiB = %d bytes\n", MiB)
	fmt.Printf("1 GiB = %d bytes\n", GiB)
	fmt.Printf("1 TiB = %d bytes\n", TiB)
	fmt.Printf("\nConstantes de configuración:\n")
	fmt.Printf("Max usuarios: %d, Timeout: %ds, Buffer: %d bytes\n",
		MaxUsers, SessionTimeout, BufferSize)
	// Uso en cálculos
	angulo := 45.0
	radianes := angulo * DegreesToRad
	seno := math.Sin(radianes)
	fmt.Printf("\nCálculo: %g° = %g radianes, sin(%g°) = %.4f\n",
		angulo, radianes, angulo, seno)
}



func main() {
	// Go NO permite conversiones implícitas automáticas
	var i int = 42
	var f float64 = 3.14
	// var result = i + f // ERROR: mismatched types int and float64
	// Conversiones explícitas requeridas
	var result1 = float64(i) + f // Convertir int a float64
	var result2 = i + int(f)     // Convertir float64 a int (trunca)
	fmt.Printf("Conversiones numéricas:\n")
	fmt.Printf("int %d + float64 %.2f = %.2f\n", i, f, result1)
	fmt.Printf("int %d + int(%.2f) = %d\n", i, f, result2)
	// Conversiones entre tipos enteros
	var a int8 = 100
	var b int16 = 200
	var c int32 = 300
	var d int64 = 400
	// Cada conversión debe ser explícita
	var suma16 = int16(a) + b
	var suma32 = int32(suma16) + c
	var suma64 = int64(suma32) + d
	fmt.Printf("Conversiones entre enteros:\n")
	fmt.Printf("int8(%d) + int16(%d) = %d\n", a, b, suma16)
	fmt.Printf("Resultado anterior + int32(%d) = %d\n", c,
		suma32)
	fmt.Printf("Resultado anterior + int64(%d) = %d\n", d,
		suma64)
	// Conversiones que pueden perder datos
	var grande int64 = 1000000
	var pequeño int8 = int8(grande) // Overflow! Solo mantiene los 8 bits menos significativos
	fmt.Printf("Overflow: int64(%d) -> int8(%d)\n", grande,
		pequeño)
	// Conversiones entre signed y unsigned
	var negativo int = -10
	var positivo uint = uint(negativo) // Comportamiento indefinido con números negativos
	fmt.Printf("Signed a unsigned: int(%d) -> uint(%d)\n",
		negativo, positivo)
}


//Conversiones String ↔ Tipos Numéricos
func main() {
	// Conversiones usando strconv
	// String a números
	strNumero := "123"
	strFloat := "3.14159"
	strBool := "true"
	// ParseInt(string, base, bitSize)
	numero, err1 := strconv.ParseInt(strNumero, 10, 64)
	if err1 != nil {
		fmt.Printf("Error convertir '%s' a int: %v\n", strNumero,
			err1)
	} else {
		fmt.Printf("String '%s' -> int64: %d\n", strNumero,
			numero)
	}
	// Atoi es equivalente a ParseInt(s, 10, 0) pero retorna int
	numeroInt, err2 := strconv.Atoi(strNumero)
	if err2 != nil {
		fmt.Printf("Error en Atoi: %v\n", err2)
	} else {
		fmt.Printf("String '%s' -> int: %d\n", strNumero,
			numeroInt)
	}
	// ParseFloat
	flotante, err3 := strconv.ParseFloat(strFloat, 64)
	if err3 != nil {
		fmt.Printf("Error convertir '%s' a float: %v\n",
			strFloat, err3)
	} else {
		fmt.Printf("String '%s' -> float64: %.5f\n", strFloat,
			flotante)
	}
	// ParseBool
	booleano, err4 := strconv.ParseBool(strBool)
	if err4 != nil {
		fmt.Printf("Error convertir '%s' a bool: %v\n", strBool,
			err4)
	} else {
		fmt.Printf("String '%s' -> bool: %t\n", strBool,
			booleano)
	}
	// Números a string
	var entero int = 456
	var flotante64 float64 = 2.71828
	var booleano2 bool = false
	// Itoa es equivalente a FormatInt(int64(i), 10)
	strDesdeInt := strconv.Itoa(entero)
	fmt.Printf("int %d -> string: '%s'\n", entero, strDesdeInt)
	// FormatFloat(f, fmt, prec, bitSize)
	strDesdeFloat := strconv.FormatFloat(flotante64, 'f', 3, 64)
	fmt.Printf("float64 %.5f -> string: '%s'\n", flotante64,
		strDesdeFloat)
	// FormatBool
	strDesdeBool := strconv.FormatBool(booleano2)
	fmt.Printf("bool %t -> string: '%s'\n", booleano2,
		strDesdeBool)
	// Manejo de errores en conversiones
	strInvalido := "abc123"
	_, err := strconv.Atoi(strInvalido)
	if err != nil {
		fmt.Printf("Error esperado al convertir '%s': %v\n",
			strInvalido, err)
	}
	// Conversiones con diferentes bases
	strBinario := "1010"
	strOctal := "755"
	strHex := "FF"
	binario, _ := strconv.ParseInt(strBinario, 2, 64)  // Base 2
	octal, _ := strconv.ParseInt(strOctal, 8, 64)      // Base 8
	hexadecimal, _ := strconv.ParseInt(strHex, 16, 64) // Base 16
	fmt.Printf("Conversiones con bases:\n")
	fmt.Printf("Binario '%s' (base 2) -> %d\n", strBinario,
		binario)
	fmt.Printf("Octal '%s' (base 8) -> %d\n", strOctal, octal)
	fmt.Printf("Hexadecimal '%s' (base 16) -> %d\n", strHex,
		hexadecimal)
}

//Conversiones String ↔ Bytes ↔ Runes


func main() {
	texto := "Hola 世界 🌍"
	// String a []byte
	bytes := []byte(texto)
	fmt.Printf("String: '%s'\n", texto)
	fmt.Printf("Bytes: %v\n", bytes)
	fmt.Printf("Longitud en bytes: %d\n", len(bytes))
	// []byte a string
	textoRecuperado := string(bytes)
	fmt.Printf("Bytes de vuelta a string: '%s'\n",
		textoRecuperado)
	// String a []rune
	runes := []rune(texto)
	fmt.Printf("Runes: %v\n", runes)
	fmt.Printf("Longitud en runes: %d\n", len(runes))
	// []rune a string
	textoDesdeRunes := string(runes)
	fmt.Printf("Runes de vuelta a string: '%s'\n",
		textoDesdeRunes)
	// Conversión individual rune a string
	var r rune = '🚀'
	strDesdeRune := string(r)
	fmt.Printf("Rune %c -> string: '%s'\n", r, strDesdeRune)
	// Conversión byte a string (cuidado con UTF-8)
	var b byte = 65 // ASCII 'A'
	strDesdeByte := string(b)
	fmt.Printf("Byte %d -> string: '%s'\n", b, strDesdeByte)
	// Análisis detallado de UTF-8
	fmt.Printf("\nAnálisis UTF-8 de '%s':\n", texto)
	for i, r := range texto {
		fmt.Printf("Posición %d: rune %c (U+%04X)\n", i, r, r)
	}
	// Validación UTF-8
	textoValido := "Texto válido"
	textoInvalido := string([]byte{0xff, 0xfe, 0xfd}) // Bytes inválidos UTF-8
	fmt.Printf("'%s' es UTF-8 válido: %t\n", textoValido,
		utf8.ValidString(textoValido))
	fmt.Printf("Texto inválido es UTF-8 válido: %t\n",
		utf8.ValidString(textoInvalido))
	// Contar runes sin conversión completa
	cantidadRunes := utf8.RuneCountInString(texto)
	fmt.Printf("Cantidad de runes en '%s': %d\n", texto,
		cantidadRunes)
}

//Conversiones Unsafe (Avanzado)
func main() {
	// ADVERTENCIA: unsafe es peligroso y rompe garantías de Go
	// Solo usar cuando sea absolutamente necesario y se entienda completamente
	// // Conversión de punteros
	var i int64 = 0x0123456789ABCDEF
	// Obtener puntero al int64
	ptr := unsafe.Pointer(&i)
	// Convertir a puntero de array de bytes
	bytePtr := (*[8]byte)(ptr)
	fmt.Printf("int64: 0x%016X\n", i)
	fmt.Printf("Como bytes: %v\n", *bytePtr)
	// Conversión string a []byte sin copia (PELIGROSO)
	s := "Hello, World!"
	// Estructura interna de string
	type StringHeader struct {
		Data uintptr
		Len  int
	}
	// Estructura interna de slice
	type SliceHeader struct {
		Data uintptr
		Len  int
		Cap  int
	}
	// Obtener header del string
	sHeader := (*StringHeader)(unsafe.Pointer(&s))
	// Crear slice header con los mismos datos
	var b []byte
	bHeader := (*SliceHeader)(unsafe.Pointer(&b))
	bHeader.Data = sHeader.Data
	bHeader.Len = sHeader.Len
	bHeader.Cap = sHeader.Len
	fmt.Printf("String: '%s'\n", s)
	fmt.Printf("Slice (sin copia): %v\n", b)
	fmt.Printf("Slice como string: '%s'\n", string(b))
	// PELIGRO: Modificar el slice afectaría el string (inmutable)
	// NO HACER: b[0] = 'h' // Esto causaría un crash
	// Tamaños de tipos
	fmt.Printf("\nTamaños de tipos (en bytes):\n")
	fmt.Printf("bool: %d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int: %d\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int8: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("float32: %d\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64: %d\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("string: %d\n", unsafe.Sizeof(string("")))
	fmt.Printf("[]byte: %d\n", unsafe.Sizeof([]byte{}))
	// Alineación de memoria
	type Estructura struct {
		a bool  // 1 byte
		b int64 // 8 bytes
		c bool  // 1 byte
	}
	var est Estructura
	fmt.Printf("\nAlineación de estructura:\n")
	fmt.Printf("Tamaño total: %d bytes\n", unsafe.Sizeof(est))
	fmt.Printf("Offset de 'a': %d\n", unsafe.Offsetof(est.a))
	fmt.Printf("Offset de 'b': %d\n", unsafe.Offsetof(est.b))
	fmt.Printf("Offset de 'c': %d\n", unsafe.Offsetof(est.c))
}

"""*/
