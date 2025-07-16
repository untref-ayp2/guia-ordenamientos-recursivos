package heapsort

func HeapSort[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	heapify(arr, compare)
	end := size - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		downHeap(arr, 0, end-1, compare)
		end--
	}
}

func heapify[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	start := (size - 2) / 2
	for start >= 0 {
		downHeap(arr, start, size-1, compare)
		start--
	}
}

func downHeap[T any](arr []T, start, end int, compare func(T, T) int) {
	parent := start
	left := parent*2 + 1
	right := left + 1
	for left <= end {
		if right <= end && compare(arr[right], arr[left]) > 0 {
			left = right
		}
		if compare(arr[left], arr[parent]) > 0 {
			arr[left], arr[parent] = arr[parent], arr[left]
			parent = left
			left = parent*2 + 1
			right = left + 1
		} else {
			return
		}
	}
}
