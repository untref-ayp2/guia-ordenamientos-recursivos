package heapsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALHeapSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)
	HeapSort(arr, func(a, b int) int { return a - b })
	assert.Equal(t, 0, len(arr))
}

func TestINTERNALHeapSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}
	HeapSort(arr, func(a, b int) int { return a - b })
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}

func TestINTERNALHeapSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}
	HeapSort(arr, func(a, b int) int { return a - b })
	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}

func TestINTERNALHeapSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	HeapSort(arr, func(a, b int) int { return a - b })
	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestINTERNALHeapSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	HeapSort(arr, func(a, b int) int { return b - a })
	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
