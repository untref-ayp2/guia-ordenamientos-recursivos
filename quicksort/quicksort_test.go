package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALQuickSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)

	QuickSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, 0, len(arr))
}

func TestINTERNALQuickSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}

	QuickSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}

func TestINTERNALQuickSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}

	QuickSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}

func TestINTERNALQuickSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}

	QuickSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestINTERNALQuickSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}

	QuickSort(arr, func(a, b int) int { return b - a })

	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
