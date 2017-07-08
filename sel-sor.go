package main
import "fmt"

func main() {
	
	original_array := []int {3312,435,2131,767,998,446,9,69,1,4,3,7,8}
	fmt.Println(original_array)
	sorted_array := selection_sort(original_array)
	fmt.Println(sorted_array)
}

func selection_sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		minimo_observado,posicion := array[i],i
		valor_original := array[i]
		for j := i + 1 ; j < len(array); j++ {
			valor_comparacion := array[j]

			if valor_comparacion < minimo_observado{
				minimo_observado,posicion = valor_comparacion,j
			}

		}

		if minimo_observado != valor_original {
			array[i],array[posicion] = minimo_observado, valor_original
		}
		
	}
	return array
	
}