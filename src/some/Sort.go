package some

import "fmt"

//冒泡
func Sort1() {
	a := []int{4, 5, 1, 6, 8, 3, 5, 9}
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	fmt.Printf("%v", a)
}

//选择
func Sort2() {
	a := []int{4, 5, 1, 6, 8, 3, 5, 9}
	for i := 0; i < len(a)-1; i++ {
		minv := a[i]
		mink := i
		for j := i + 1; j <= len(a)-1; j++ {
			if a[j] < minv {
				mink = j
				minv = a[j]
			}
		}
		a[i], a[mink] = a[mink], a[i]
	}
	fmt.Printf("%v", a)
}

//快排
func Sort3() {
	a := []int{4, 5, 1, 6, 8, 3, 5, 9}
	a = sort3_b(0, len(a)-1, a)
	fmt.Printf("%v", a)
}
func sort3_b(l, r int, a []int) []int {
	if l >= r {
		return a
	}

	kv := a[l]
	b := l
	e := r
	for b < e {
		for b < e {
			if kv <= a[e] {
				e--
			} else {
				a[b] = a[e]
				break
			}
		}

		for b < e {
			if kv >= a[b] {
				b++
			} else {
				a[e] = a[b]
				break
			}
		}
	}
	a[b] = kv

	a = sort3_b(l, b-1, a)
	a = sort3_b(b+1, r, a)

	return a
}

//堆排
func Sort4() {
	a := []int{4, 5, 1, 6, 8, 3, 5, 9}
	for i := 0; i < len(a); i++ {
		a = sort4_b(a, i)
	}

	for i := len(a) - 1; i > 0; i-- {
		a = sort4_c(a, i)
	}
	fmt.Println(a)
}

//生成堆
func sort4_b(a []int, k int) []int {
	for (k-1)/2 >= 0 {
		if a[(k-1)/2] < a[k] {
			a[(k-1)/2], a[k] = a[k], a[(k-1)/2]
		} else {
			break
		}
		k = (k - 1) / 2
	}
	return a
}

//平衡堆
func sort4_c(a []int, e int) []int {
	a[0], a[e] = a[e], a[0]
	e--

	for i := 0; e > 0 && i <= (e-1)/2; {
		r := (i + 1) * 2
		l := (i+1)*2 - 1
		max := l
		if r <= e && a[l] < a[r] {
			max = r
		}
		if a[i] < a[max] {
			a[i], a[max] = a[max], a[i]
			i = max
		} else {
			break
		}
	}
	return a
}
