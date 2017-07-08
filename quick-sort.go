package main 

import "fmt"

func main() {
	
	original_array := []int {3312,435,2131,767,998,446,9,69,1,4,3,7,8}
	fmt.Println(original_array)
	quick_sort(&original_array,0,len(original_array)-1)
	fmt.Println(original_array)
}

func quick_sort(array_ptr *[]int, left, right int){
	if left < right{
		array := *array_ptr

		limit,origin := right,left

		pivot := array[right]
		right--;

		for left <= right{

			for left < len(array) && array[left] < pivot{
				left++
			}


			for right >= 0 && array[right] > pivot{
				right--
			}


			if left <= right {
			   swap(array_ptr,left,right)

			   left++

			   right--
		}

		}

		swap(array_ptr,left,limit)
		quick_sort(array_ptr,origin,right)
		quick_sort(array_ptr,left,limit)

	}
}



func swap(array_ptr *[]int,left,right int){

	array := *array_ptr
	tmp := array[left]
	array[left] = array[right]
	array[right] = tmp
}




