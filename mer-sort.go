package main 

import "fmt"

func main() {
	un_sort := []int{2,45,76,23,657,988,543,10,1}
	fmt.Println(un_sort)
	sorted_a := merge_sort(un_sort) 
	fmt.Println(sorted_a)
}

func merge_sort(array []int)[]int {

    if len(array) == 1 {
    	
    	return array
    }

	array_size := len(array)
	half := array_size/2
	left := array[0:half]
	right := array[half:array_size]

	left = merge_sort(left)
	right = merge_sort(right)

	return merge(left,right)

}

func merge(array_left, array_right []int)[]int {

	last_left := array_left[len(array_left) -1]
	first_right := array_right[0]
	if last_left <= first_right {
		return append(array_left,array_right...)
	}

	resultado := make([]int,0)
	var menor int

	for len(array_left) > 0 && len(array_right) > 0{
			if array_left[0] <= array_right[0]{
				menor = array_left[0]
				array_left = array_left[1:]
			}else{
				menor = array_right[0]
				array_right = array_right[1:]
			}
			resultado = append(resultado,menor)

		}
		if len(array_left) > 0{
			resultado = append(resultado,array_left...)
		}

		if len(array_right) > 0{
			resultado = append(resultado,array_right...)
		}
	return resultado
}