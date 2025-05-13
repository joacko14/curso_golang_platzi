package main

import (
	"fmt"
	"math"
)

func printMessage(message string) {
	fmt.Println(message)
}

func greeting(name string, today string) (greeting string) {
	// Esta es una funcion que recibe un nombre y devuelve un saludo
	// y la fecha de hoy
	return fmt.Sprintf("Hola %s, hoy es %s", name, today)
}

func interpreter(args []string) (message string, count int) {
	// Esta es una funcion que recibe un slice de strings
	// y devuelve un string y la cantidad de elementos en el slice
	return fmt.Sprintf("%s %s %s", args[0], args[1], args[2]), len(args)
}

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
	// area de un cuadrado, rectangulo, circulo y trapecio
	const baseCuadrado float64 = 25
	areaCuadrado := math.Pow(baseCuadrado, 2)
	p("Area del cuadrado: ", areaCuadrado, " con base: ", baseCuadrado)

	const base int = 15
	const altura int = 10
	areaRectangulo := base * altura
	p("Area del rectangulo: ", areaRectangulo, " con base: ", base, " y altura: ", altura)

	const radio float64 = 5
	areaCirculo := math.Pi * math.Pow(radio, 2)
	p("Area del circulo: ", areaCirculo, " con radio: ", radio)

	const baseMayor int = 20
	const baseMenor int = 10
	const alturaTrapecio int = 5
	areaTrapecio := ((baseMayor + baseMenor) * alturaTrapecio) / 2
	p("Area del trapecio: ", areaTrapecio, " con base mayor: ", baseMayor, " base menor: ", baseMenor, " y altura: ", alturaTrapecio)

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

	// Pruebas del paquete fmt

	bytes, err := fmt.Print("Hola mundo\n")

	fmt.Println("Bytes: ", bytes, " Error: ", err)

	fmt.Printf("Hola mundo %s\n", "Gopher")
	fmt.Printf("bytes is type %T\n", bytes)
	fmt.Printf("err is type %T\n", err)
	fmt.Printf("Pi is %f\n", pi)

	// Pruebas de funciones
	printMessage("Hola mundo")
	p(greeting("Gopher", "hoy es lunes"))
	message, count := interpreter([]string{"Hola", "mundo", "Gopher"})
	p(message, count)
	// Si necesito ignorar uno de los valores devueltos
	_, count = interpreter([]string{"Hello", "world", "Gopher"})
	p("Cantidad de elementos: ", count)
}
