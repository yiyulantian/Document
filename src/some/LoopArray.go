package some

import (
	"fmt"
)

func LoopArray() {
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	h, l := 0, 0
	b := 2
	r := 4
	for {
		ringArray(arr, h, l, b, r)
		h++
		l++
		b--
		r--
		if h > b || l > r {
			break
		}
	}

}

//begin point(h,l) | end point(b,r)
func ringArray(arr [][]int, h, l, b, r int) {

	//hl => hr
	for i := l; i <= r; i++ {
		fmt.Print(arr[h][i], " ")
		fmt.Println("")
	}

	//hr => br
	for i := h + 1; i <= b; i++ {
		fmt.Print(arr[i][r], " ")
		fmt.Println("")
	}

	//br => bl
	for i := r - 1; b > h && i >= l; i-- {
		fmt.Print(arr[b][i], " ")
		fmt.Println("")
	}

	//bl => hl
	for i := b - 1; r > l && i > h; i-- {
		fmt.Print(arr[i][l], " ")
		fmt.Println("")
	}

}
