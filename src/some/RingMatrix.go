package some

import "fmt"

//循环向外写数,1位置视为向量中心，则f(0,0)=1,f(1,1)=9,f(1,2)=24,求f(x,y)
/*
  21 22 23 24 25 26
  20  7  8  9 10 27 ....
  19  6  1  2 11
  18  5  4  3 12
  17 16 15 14 13
*/
//视x,x点为一次循环的结束，则xx这一圈的数字个数是8*x个，则f(x,x)=f(x-1,x-1)+8x
func RingMatrix() {
	x := 2
	y := 2
	abX := (x ^ x>>31) - x>>31
	abY := (y ^ y>>31) - y>>31
	t := abX
	if abX < abY {
		t = abY
	}

	fi := 1
	for i := 1; i <= t; i++ {
		fi = fi + 8*i
	}
	if x == y && x == t {
		fmt.Println(fi)
	} else if t == x {
		fmt.Println(fi - 8*t + (t - y))
	} else if t == abX {
		fmt.Println(fi - 4*t + (t + y))
	} else if t == y {
		fmt.Println(fi - 2*t + (t + x))
	} else {
		fmt.Println(fi - 6*t + (t - x))
	}
}
