package some

import (
	"fmt"
)

func TopK() {
	a := []int{2, 3, 1, 4, 6, 2, 1, 4, 7, 4, 6, 5, 8}
	k := 7
	heap := createHeap(a[:k])
	//fmt.Println(heap)
	for i := k; i < len(a); i++ {
		if a[i] < heap[0] {
			heap[0] = a[i]
			developHeap(&heap)
		}
	}
	for _, v := range heap {
		fmt.Println(v)
	}
}

func createHeap(a []int) []int {
	heap := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		j := i
		heap[j] = a[j]
		for j > 0 {
			if heap[(j-1)/2] < heap[j] {
				heap[(j-1)/2], heap[j] = heap[j], heap[(j-1)/2]
				j = (j - 1) / 2
			} else {
				break
			}
		}
	}
	return heap
}

//大顶堆 求最小
func developHeap(heap *[]int) {
	tmp := *heap
	i := 0
	for i*2+1 < len(tmp) {
		if (i*2+2) >= len(tmp) || tmp[i*2+1] >= tmp[i*2+2] {
			if tmp[i] >= tmp[i*2+1] {
				return
			} else {
				htmp := tmp[i*2+1]
				tmp[i*2+1] = tmp[i]
				tmp[i] = htmp
				i = i*2 + 1
			}
		} else {
			if tmp[i] >= tmp[i*2+2] {
				return
			} else {
				htmp := tmp[i*2+2]
				tmp[i*2+2] = tmp[i]
				tmp[i] = htmp
				i = i*2 + 2
			}
		}
		//fmt.Println(heap)
	}
	return
}

func TopKQuick() {
	a := []int{2, 3, 1, 4, 6, 2, 1, 4, 7, 4, 6, 5, 8}
	k := 10
	offset := quickFind(&a, 0, len(a)-1)

	for offset != k-1 {
		if offset > k-1 {
			offset = quickFind(&a, 0, offset-1)
		} else {
			offset = quickFind(&a, offset+1, len(a)-1)
		}
	}
	//fmt.Println(a)
	for i := 0; i < k; i++ {
		fmt.Println(a[i])
	}
}
func quickFind(a *[]int, l, h int) int {
	//fmt.Println("quickFind")
	tmp := *a
	k := tmp[l]
	offset := l
	l++
	for l <= h {
		for l <= h && k <= tmp[h] {
			h--
		}
		if l <= h {
			tmp[offset] = tmp[h]
			offset = h
			h--
		}

		for l <= h && k >= tmp[l] {
			l++
		}
		if l <= h {
			tmp[offset] = tmp[l]
			offset = l
			l++
		}

	}
	tmp[offset] = k
	return offset
}
