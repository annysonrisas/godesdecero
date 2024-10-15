package arreglos_slices

import "fmt"

//var tabla [10]int
var tabla [10]int = [10]int{0, 23, 0, 39}
var matriz [5][10]int

func MuestroArrelos() {
	tabla[7] = 33
	tabla[2] = 54
	tabla[0] = 11

	tabla2 := [10]string{"Antonia", "Chris", "", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		for tabla2[i] != "" {
			fmt.Println(tabla2[i])
			break
		}
	}
	matriz[3][8] = 25
	fmt.Println(matriz)
}
