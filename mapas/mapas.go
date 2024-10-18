package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])
}
