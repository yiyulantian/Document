package some

import "fmt"

// 排队问题：（h，k）表示身高h，前面有k个比h大的元素；给出一组数据对，整理出合适序列
// 思路：先将数据组按身高排序，然后依次遍历排好的数据组，取出元素的k，将最终序列前方k个空位留空，将元素放在第k+1个空位上，依次类推，序列就被填满了
func KOrder() {
	//7 0 4 4 7 1 5 0 6 1 5 2
	h := []int{7, 4, 7, 5, 6, 5}
	k := []int{0, 4, 1, 0, 1, 2}
	//5 7 5 6 4 7
	for i := len(h) - 1; i > 0; i-- {
		tv := h[0]
		tk := 0
		for j := 0; j <= i; j++ {
			if h[j] > tv {
				tk = j
				tv = h[j]
			}
		}
		tmp := h[tk]
		h[tk] = h[i]
		h[i] = tmp
		tmp = k[tk]
		k[tk] = k[i]
		k[i] = tmp
	}
	fmt.Println(h, k)
	d := make([]int, len(h))
	for i := range h {
		j := k[i]
		for s := 0; s < len(d); s++ {
			if d[s] == 0 {
				if j > 0 {
					j--
				} else {
					d[s] = h[i]
					break
				}
			}
		}
	}
	fmt.Println(d)
}
