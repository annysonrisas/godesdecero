package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMidlleware() {
	fmt.Println("inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(12, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(3, 5)
	fmt.Println(result)

}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
