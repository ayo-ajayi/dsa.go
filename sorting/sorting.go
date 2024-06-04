package sorting

func BubbleSort(arrPtr *[]int) {
	arr := *arrPtr
	for range arr {
		swapped := false
		for i, j := 0, 1; j < len(arr); i, j = i+1, j+1 {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// find the smallest element using a linear scan
// and move it to the front swapping it with the front element
func SelectionSort(arrPtr *[]int) {
	arr := *arrPtr
	for i := range arr {
		index := i + minIndex(arr[i:])
		arr[index], arr[i] = arr[i], arr[index]
	}
}

func minIndex(arr []int) int {
	min := arr[0]
	index := 0
	for i, num := range arr {
		if num < min {
			min = num
			index = i
		}
	}
	return index
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result

}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	left := MergeSort(arr[:n/2])
	right := MergeSort(arr[n/2:])
	return merge(left, right)

}

func QuickSort(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)
}
func quickSort(arr *[]int, left, right int) {
	var index int = partition(*arr, left, right)
	if left < index-1 {
		quickSort(arr, left, index-1)
	}
	if index < right {
		quickSort(arr, index, right)

	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}
