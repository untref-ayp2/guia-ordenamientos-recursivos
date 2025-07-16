package quicksort

func QuickSort[T any](arr []T, compare func(a, b T) int) {
	quickSort(arr, compare, 0, len(arr)-1)
}

func quickSort[T any](arr []T, compare func(a, b T) int, low, high int) {
	if low < high {
		pi := partition(arr, compare, low, high)
		quickSort(arr, compare, low, pi-1)
		quickSort(arr, compare, pi+1, high)
	}
}

func partition[T any](arr []T, compare func(a, b T) int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if compare(arr[j], pivot) < 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
