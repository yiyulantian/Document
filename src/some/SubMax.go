package some

import (
	"fmt"
)

//SubArrMax 求数组中最大的子串和
func SubArrMax() {
	a := []int{1, 2, -1, -3, 4 - 2, 5, 1, -3, 7}
	max := subArrMax(a)
	fmt.Println(max)
}

//O(n)
func subArrMax(a []int) int {
	curMax := a[0]
	Max := a[0]
	for i := 1; i < len(a); i++ {
		if curMax+a[i] < a[i] {
			curMax = a[i]
		} else {
			curMax = curMax + a[i]
		}
		if curMax > Max {
			Max = curMax
		}
	}
	return Max
}

//SubMatrixMax 求矩阵中子矩阵和最大
func SubMatrixMax() {
	a := [][]int{
		{1, 2, -3, 2},
		{-1, 2, 4, 1},
		{3, 4, 3, -4},
	}
	Max := a[0][0]
	for i := 0; i < len(a); i++ {
		subArr := []int{0, 0, 0, 0}
		for j := i; j < len(a); j++ {
			for k, v := range a[j] {
				subArr[k] += v
			}
			thisMax := subArrMax(subArr)
			if thisMax > Max {
				Max = thisMax
			}
		}
	}
	fmt.Println(Max)
}
