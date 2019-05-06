package some

import "fmt"

//BitMap 简单bitmap 求一组数中，唯一出现一(单)次的数，数范围256 ,256/8=32 byte
func BitMap() {
	data := []int{10, 10, 12, 13, 14, 14, 13, 12, 11, 7, 7, 11, 13}
	a := make([]byte, 32)
	for _, d := range data {
		n := d / 8
		m := uint(d % 8)
		k := 1 << m
		if int(a[n])&k == 0 {
			a[n] = byte(int(a[n]) + k)
		} else {
			a[n] = byte(int(a[n]) - k)
		}
	}

	for n, md := range a {
		i := -1
		for ; md > 0; md = md >> 1 {
			i++
		}
		if i >= 0 {
			fmt.Println((n*8 + i))
			return
		}
	}
}
