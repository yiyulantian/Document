package some

import "fmt"

//向量操作

//给出平面坐标系上四个点的坐标，判断这四个点能不能连成正方形。
/*
只需要依次判断两条边是否等长，且为直角即可
两点（x1,y1）(x2,y2)边长的平方：(x1-x2)平方+(y1-y2)平方
判断是否直角两个方法：1.求斜率；2.勾股定理
*/
type point struct {
	x int64
	y int64
}

func FormSquare() {
	a := []point{{0, 0}, {1, 2}, {0, 1}, {1, 0}}
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	a01 := (a0.x-a1.x)*(a0.x-a1.x) + (a0.y-a1.y)*(a0.y-a1.y)
	a02 := (a0.x-a2.x)*(a0.x-a2.x) + (a0.y-a2.y)*(a0.y-a2.y)
	a03 := (a0.x-a3.x)*(a0.x-a3.x) + (a0.y-a3.y)*(a0.y-a3.y)
	var la, sa, rla, rsa int64

	if a01 == a02 {
		la = a03
		sa = a01
		rla = (a2.x-a1.x)*(a2.x-a1.x) + (a2.y-a1.y)*(a2.y-a1.y)
		rsa = (a3.x-a1.x)*(a3.x-a1.x) + (a3.y-a1.y)*(a3.y-a1.y)
	} else if a01 == a03 {
		la = a02
		sa = a01
		rla = (a3.x-a1.x)*(a3.x-a1.x) + (a3.y-a1.y)*(a3.y-a1.y)
		rsa = (a2.x-a1.x)*(a2.x-a1.x) + (a2.y-a1.y)*(a2.y-a1.y)
	} else if a02 == a03 {
		la = a01
		sa = a02
		rla = (a2.x-a3.x)*(a2.x-a3.x) + (a2.y-a3.y)*(a2.y-a3.y)
		rsa = (a3.x-a1.x)*(a3.x-a1.x) + (a3.y-a1.y)*(a3.y-a1.y)
	} else {
		fmt.Println("0")
		return
	}
	if la != 2*sa || la != rla || rsa != sa {
		fmt.Println("0")
		return
	}
	fmt.Println("1")
}
