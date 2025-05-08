package main

import (
	"fmt"
	"math"
)

func main() {

	p := fmt.Println

	// Declaración de constantes
	const name string = "Gopher"
	const pi = math.Pi

	// Declaración de variables
	apellido := "Golang"
	var experiencia int = 10
	var sueldo float64

	p(name, " ", apellido, pi, " ", experiencia, " ", sueldo)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	p(a, b, c, d)

	// Calculo de propiedades geometricas
	const baseCuadrado = 25
	areaCuadrado := math.Pow(baseCuadrado, 2)
	p("Area del cuadrado: ", areaCuadrado, " con base: ", baseCuadrado)

	// Operaciones aritmeticas

	// Suma
	x := 10
	y := 50
	resultado := x + y
	p("Dados dos numeros x", x, " y ", y)

	p("Resultado de la suma: ", resultado)

	// Resta
	result := x - y
	p("Resultado de la resta: ", result)

	// Multiplicacion
	resultado = x * y
	p("Resultado de la multiplicacion: ", resultado)

	// Division
	resultado = y / x
	p("Resultado de la division: ", resultado)

	// Modulo
	resultado = y % x
	p("Resultado del modulo: ", resultado)

	// Incrmento
	x++
	p("Resultado del incremento: ", x)

	// Decremento
	x--
	p("Resultado del decremento: ", x)
}
