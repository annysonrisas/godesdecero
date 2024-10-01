package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func TabladeMultiplicar() {
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
					fmt.Println(numero1, " X ", i, " =", numero1*i)
				}
				exit = 5
			}
		}
	}
}
