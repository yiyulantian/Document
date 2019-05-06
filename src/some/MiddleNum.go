package some

import (
	"fmt"
	"sort"
)

//MiDianNum .
func MiDianNum() {
	base()
	highGrade()
}
func base() {
	a := []int{2, 3, 4, 5, 6, 7}
	b := []int{7, 6, 5, 4, 3, 2, 1}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))
	count := len(a) + len(b)
	middle := count / 2

	this, an, bn := 0, 0, 0
	for i := 0; i < middle; i++ {
		if a[an] <= b[bn] {
			this = a[an]
			an++
		} else {
			this = b[bn]
			bn++
		}
	}
	fmt.Println(this)
}

func highGrade() {
	a := []int{2, 3, 4, 5, 6, 7}
	b := []int{7, 6, 5, 4, 3, 2, 1}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))
	la := len(a)
	lb := len(b)
	middle := (la + lb) / 2
	this := recursion(a, b, middle)
	fmt.Println(this)
}

//废柴
func recursion1(a, b []int, n int) (this int) {
	al := (len(a) - 1) / 2
	bl := (len(b) - 1) / 2

	ad := a[al]
	bd := b[bl]
	if ad <= bd {
		if al+1 >= n {
			return a[n-1]
		} else {
			n = n - (al + 1)
			if al == 1 {
				return b[n-2]
			}
			return recursion(a[al+1:], b, n)
		}
	} else {
		if bl+1 >= n {
			return b[n-1]
		} else {
			n = n - (bl + 1)
			if bl == 1 {
				return a[n-2]
			}
			return recursion(a, b[bl+1:], n)
		}
	}
}

func recursion(a, b []int, n int) (this int) {
	if len(a) == 0 {
		return b[n-1]
	}
	if len(b) == 0 {
		return a[n-1]
	}

	if n == 1 {
		if a[0] <= b[0] {
			return a[0]
		}
		return b[0]
	}

	as, bs := n/2, n/2
	if len(a) < as {
		as = len(a)
	}
	if len(b) < bs {
		bs = len(b)
	}

	if a[as] <= b[bs] {
		return recursion(a[as:], b, n-as)
	} else {
		return recursion(a, b[bs:], n-bs)
	}

}
