package main

import (
	"fmt"
	"github/annysonrisas/godesdecero/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
