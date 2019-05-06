package some

import (
	"fmt"
)

func KMP() {
	t := []byte{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 1, 2, 3, 8, 9, 10}
	p := []byte{1, 2, 3, 4, 5, 6, 1, 2, 3}

	move := getMove(p)
	k := 0
	i := 0
	for i < len(t) {
		if t[i] == p[k] {
			k++
			i++
			if k == len(p) {
				break
			}
		} else if k == 0 {
			i++
		} else {
			k = move[k-1]
		}
	}

	fmt.Println(i - len(p))
}
func getMove(p []byte) []int {
	move := make([]int, len(p))
	move[0] = 0
	i := 1
	for i < len(p) {
		k := move[i-1]
		if p[k] == p[i] {
			move[i] = k + 1
		} else {
			move[i] = 0
		}
		i++
	}
	return move
}
