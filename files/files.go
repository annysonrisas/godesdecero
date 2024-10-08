package files

import (
	"bufio"
	"fmt"
	"github/annysonrisas/godesdecero/ejercicios"
	"os"
)

var filename string = "./files/archivostxt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return

	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writeString " + err.Error())
		return false
	}
	arch.Close()
	return true
}

/* Esto es usando "io/ioutil" que esta deprecado func LeoArchivo() {
archivo, err := ioutil.ReadFile(filename)
if err != nil {
	fmt.Println("Error al leer el archivo " + err.Error())
	return
}*/

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(string("> " + registro))
	}

	archivo.Close()

}
