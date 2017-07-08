package main

import "fmt"

func main() {

	un_sort := []int{2,45,76,23,657,988,543,10,1}
	fmt.Println(un_sort)
	sorted_a := bubble_sort(un_sort) 
	fmt.Println(sorted_a)
	
}

func bubble_sort(array []int)[]int {
	intercambio := true
	numero_ordenados := 0
	for intercambio{
		intercambio = false
		for i := 1; i < (len(array)-numero_ordenados); i++ {
			if array[i-1] > array[i]{
				swap(&array,i);
				intercambio = true
			}
			
		}
		numero_ordenados++
	}

	return array
}

func swap(puntero_array *[]int, index_derecha int){
	index_izquierda := index_derecha-1
	array := *puntero_array
	copia := array[index_izquierda]
	array[index_izquierda] = array[index_derecha]
	array[index_derecha] = copia
}