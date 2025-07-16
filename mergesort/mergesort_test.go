package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALMergeSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)

	arr = MergeSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, 0, len(arr))
}

func TestINTERNALMergekSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}

	arr = MergeSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}

func TestINTERNALMergeSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}

	arr = MergeSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}

func TestINTERNALMergeSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}

	arr = MergeSort(arr, func(a, b int) int { return a - b })

	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestINTERNALMergeSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}

	arr = MergeSort(arr, func(a, b int) int { return b - a })

	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
