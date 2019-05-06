package some

import (
	"fmt"
	"time"
)

type IntStack []int

func (is IntStack) Len() int {
	return len(is)
}

func (is IntStack) IsEmpty() bool {
	return len(is) == 0
}
func (is IntStack) Top() int {
	if is.IsEmpty() {
		return -1
	}
	return is[is.Len()-1]
}
func (is *IntStack) Pop() int {
	if is.IsEmpty() {
		return -1
	}
	tmp := *is
	this := tmp[is.Len()-1]
	*is = tmp[:is.Len()-1]
	return this
}
func (is *IntStack) Push(i int) {
	*is = append(*is, i)
}

func MiddleTree() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	this := 0
	is := &IntStack{}
	for this != -1 || !is.IsEmpty() {
		//fmt.Println(this)
		for this != -1 && this < len(a) {
			is.Push(this)
			this = this*2 + 1
		}
		this = -1
		if !is.IsEmpty() {
			this = is.Pop()
			fmt.Println(a[this])
			if this*2+2 < len(a) {
				this = this*2 + 2
			} else {
				this = -1
			}
		}
	}
}

func PreTree() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	this := 0
	is := &IntStack{}
	for this != -1 || !is.IsEmpty() {
		for this != -1 {
			fmt.Println(a[this])
			is.Push(this)
			if this*2+1 < len(a) {
				this = this*2 + 1
			} else {
				this = -1
			}
		}
		if !is.IsEmpty() {
			this = is.Pop()
			//fmt.Println("-----", this)
			if this*2+2 < len(a) {
				this = this*2 + 2
			} else {
				this = -1
			}
		}
	}
}

func SufTree() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	this := 0
	is := &IntStack{}

	for this != -1 || !is.IsEmpty() {
		for this != -1 {
			is.Push(this)
			if this*2+1 < len(a) {
				this = this*2 + 1
			} else {
				this = -1
			}
		}
		this = is.Top()
		if this*2+2 < len(a) {
			this = this*2 + 2
			continue
		}

		if !is.IsEmpty() {
			this = is.Pop()
			fmt.Println(a[this])
			//fmt.Println("-----", this)
			for !is.IsEmpty() && (this)/2 == is.Top()+1 {
				this = is.Pop()
				fmt.Println(a[this])
			}
			time.Sleep(1 * time.Second)
			this = -1
		}
	}

}
