package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	exit := 1

	for exit == 1 {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("numero ingresado incorrecto.")
			} else {
				for i := 1; i <= 10; i++ {
					texto += fmt.Sprintln(numero1, " X ", i, " =", numero1*i)
				}
				exit = 5
			}
		}
	}

	return texto
}
