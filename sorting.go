package main

import "fmt"

func partition(arr []Data, low, high int) ([]Data, int) {
	pivot := arr[high].score
	i := low
	for j := low; j < high; j++ {
		if arr[j].score < pivot {			// Change this line if you want 'asc' or 'desc'
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []Data, low, high int) []Data {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []Data) []Data {
	return quickSort(arr, 0, len(arr)-1)
}

func Sorting(collection []Data) {
	fmt.Println("Result: ", quickSortStart(collection))
}
