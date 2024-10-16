package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{3, 45, 12, 3, 44, 54, 6, 12, 34, 4}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // slice creado desde posicion 3
	porcion2 := arreglo[:5]  // slice creado desde 0 hasta 5
	porcion3 := arreglo[6:7] //slice creado desde posicion 6 a posicion 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}
func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
