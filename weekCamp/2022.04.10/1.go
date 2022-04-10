package weekCamp

import "sort"

func largestInteger(num int) int {
    array1 := make([]int, 0)
	array2 := make([]int, 0)
	array := make([]bool, 0)
    
    for num > 0 {
		x := num % 10
		if x%2 == 0 {
			array1 = append(array1, x)
			array = append(array, true)
		} else {
			array2 = append(array2, x)
			array = append(array, false)
		}
		num /= 10
    }

	sort.Ints(array1)
	sort.Ints(array2)

	ret := 0
	for i:=len(array)-1; i>=0; i-- {
		ret *= 10
		if array[i] {
			ret += array1[len(array1)-1]
			array1 = array1[:len(array1)-1]
		} else {
			ret += array2[len(array2)-1]
			array2 = array2[:len(array2)-1]
		}
	}

	return ret
}