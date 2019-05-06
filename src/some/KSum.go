package some

import (
	"fmt"
	"sort"
)

//KSum .
func KSum() {
	fmt.Println("begin")
	k := 3
	t := 9
	g := []int{1, 2, 3, 3, 4, 5, 6, 7, 8}
	sort.Sort(sort.IntSlice(g))
	ksum(k, t, g, "")
}

func ksum(k, t int, g []int, pres string) {
	//fmt.Println(k, t)
	if k == 2 {
		b := 0
		e := len(g) - 1
		for e > b {
			if g[b]+g[e] > t {
				//fmt.Println(b, e)
				e--
			} else {
				//fmt.Println("---", b, e, g[b], g[e])
				if g[b]+g[e] == t {

					fmt.Printf("%s,%d,%d\n", pres, g[b], g[e])
				}
				b++
			}
		}
		return
	}

	for i := 0; i <= len(g)-k+1; i++ {
		ksum(k-1, t-g[i], g[i+1:], fmt.Sprintf("%s,%d", pres, g[i]))
	}

	return
}
