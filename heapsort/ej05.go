package heapsort

func RecursiveHeapSort[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	recuriveHeapify(arr, compare)
	end := size - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		// recursiveDownHeap(...)
		end--
	}
}

func recuriveHeapify[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	start := (size - 2) / 2
	for start >= 0 {
		// recursiveDownHeap(...)
		start--
	}
}

func recursiveDownHeap[T any](array []int, start, end int, compare func(T, T) int) {
	// Implementar
}
