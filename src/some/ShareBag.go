package some

import "fmt"

// 一包中有很多不同重量的金子，问能不能被按重量平分
//思路：其实是个01背包问题，只不过重量和价值是等价的（可以认为重量=价值）；如果是两人平分，其实就是求是否存在组合背包内元素使得组合重量的和 == 总重量/2；
//	   换句话说，就是一个可容纳总重量/2的背包，可携带的最大价值，如果可以完美装满，则等价于可以平分金子
func ShareBag() {
	//1,6,8,3,5,9
	g := []int{1, 6, 8, 3, 5, 9}
	sum := 0
	for _, v := range g {
		sum += v
	}

	if sum%2 != 0 { //元素都是整数，如果不是偶数，肯定不能平分
		fmt.Println("false")
		return
	}
	//背包大小
	max := sum / 2
	//fmt.Println(max)
	data := make([][]int, len(g))
	for i := range data {
		data[i] = make([]int, max+1)
		for j := range data[i] {
			if i == 0 {
				if g[i] > j {
					data[i][j] = 0
				} else {
					data[i][j] = g[i]
				}
			} else {
				if g[i] > j {
					data[i][j] = data[i-1][j]
				} else {
					zore := data[i-1][j]
					one := data[i-1][j-g[i]] + g[i] //这里重量即价值，所以都用g[i]
					if zore > one {
						data[i][j] = zore
					} else {
						data[i][j] = one
					}
				}
			}
		}
	}
	//fmt.Println(data)
	for i := range data {
		if max == data[i][max] {
			fmt.Println("ture")
			return
		}
	}
}
