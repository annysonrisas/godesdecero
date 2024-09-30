package main

import (
	"fmt"
	"github/annysonrisas/godesdecero/variables/ejercicios"
)

func main() {
	/*	estado, texto := variables.ConviertoaTexto(1588)
		fmt.Println(estado)
		fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es", os)
	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
	numero, texto := ejercicios.ConviertoaNumerico("40")
	fmt.Println(numero)
	fmt.Println(texto)
}
