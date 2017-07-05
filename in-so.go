package main
import "fmt"

func main() {
	un_sort := []int{2,45,76,23,657,988,543,10,1}
	fmt.Println(un_sort)
	sorted_a := insertion_sort(un_sort) 
	fmt.Println(sorted_a)
}

func insertion_sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			swap(j-1,j,&array);
			j--;
		}
		
	}
	return array
}

func swap(previus int, current int,array *[]int) {
	
	arr := *array
	s_copy := arr[current]
	arr[current] = arr[previus]
	arr[previus] = s_copy
}