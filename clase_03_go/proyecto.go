package main

import "fmt"

// Definición básica de un struct
type Usuario struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

// Struct con tags (útil para JSON, validaciones, etc.)
type Producto struct {
	ID     int     `json:"id" db:"product_id"`
	Nombre string  `json:"nombre" validate:"required"`
	Precio float64 `json:"precio" validate:"min=0"`
}

// Struct anidado
type Empresa struct {
	Nombre    string
	Direccion Direccion
	Empleados []Usuario
}
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

//PARA RECEPTORES

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int { // Value receiver
	return r.Width * r.Height
}
func (r *Rectangle) Scale(factor int) { // Pointer receiver
	r.Width *= factor
	r.Height *= factor
}

// comportamiento en memoria
type DatosGrandes struct {
	buffer [1000000]byte
	nombre string
	id     int
}

// ✅ EFICIENTE - Solo pasa referencia (8 bytes en x64)
func (d *DatosGrandes) ProcesarDatosPtr() {
	fmt.Printf("Procesando: %s\n", d.nombre)
}

// Mutabilidad y Side effects
type Contador struct {
	valor int
	logs  []string
}

// Receptor de VALOR - NO modifica el original
func (c Contador) IncrementarCopia() {
	c.valor++
	c.logs = append(c.logs, "incrementado")
}

// Receptor de PUNTERO - SÍ modifica el original
func (c *Contador) IncrementarOriginal() {
	c.valor++
	c.logs = append(c.logs, "incrementado")
}

// Métodos faltantes
func (u Usuario) MetodoValor() {
	fmt.Println("Método de valor:", u.Nombre)
}

func (u *Usuario) MetodoPuntero() {
	fmt.Println("Método de puntero:", u.Nombre)
}

func main() {

	//funciones para inicializar

	// Inicialización con valores cero
	var u1 Usuario

	// Inicialización con valores específicos
	u2 := Usuario{
		ID:     1,
		Nombre: "Bryan",
		Email:  "bryan@email.com",
		Activo: true,
	}
	// Inicialización parcial (otros campos toman valor cero)
	u3 := Usuario{
		Nombre: "Daniela",
		Email:  "daniela@email.com",
	}
	// Usando punteros (más eficiente para structs grandes)
	u4 := &Usuario{
		ID:     2,
		Nombre: "Claudia",
	}

	//structs ANONIMOS

	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "milot"
	person.age = 50
	person.pet = "dog"
	pet := struct {
		name string
		kind string
	}{
		name: "doki",
		kind: "dog",
	}

	fmt.Println(u1, u2, u3, u4, person, pet)

	//Receptores go

	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area() // Calling the Area method
	fmt.Println("Area:", area)
	rect.Scale(2) // Calling the Scale method (modifies the original rect)
	fmt.Println("Scaled rect:", rect)

	//ANATOMIA DE UN RECEPTOR
	/*
		// Sintaxis básica
		func (receptor TipoReceptor) NombreMetodo(parametros) retorno {
		// implementación
		}
		// Ejemplos
		func (u Usuario) Leer() string // Receptor de valor
		func (u *Usuario) Escribir() error // Receptor de puntero
	*/

	// Mutabilidad y Side effects
	contador := Contador{valor: 0, logs: []string{}}
	contador.IncrementarCopia()
	fmt.Println("Después de copia:", contador.valor)
	contador.IncrementarOriginal()
	fmt.Println("Después de original:", contador.valor)

	// Métodos con receptor valor y puntero
	usuario := Usuario{Nombre: "Carlos"}
	usuario.MetodoValor()
	usuario.MetodoPuntero()

	ptr := &Usuario{Nombre: "Carlos"}
	ptr.MetodoValor()
	ptr.MetodoPuntero()

}
