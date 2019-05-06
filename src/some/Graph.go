package some

import "fmt"

//图

//给定一个数字串 (均为正整数)，现在需要从第一个数跳跃到最后一个，所在位置的数字表示可以跳跃的最大步数。求出从第一个位置跳跃到最后位置所需的最少步数。
/*
无权图，最短路径
广度搜索就行了
*/

type Queue struct {
	len    int64
	header *Qelem
}

type Qelem struct {
	key   int
	value int64
	next  *Qelem
	pre   *Qelem
}

func (q *Queue) Push(key int, value int64) {
	qe := &Qelem{
		key:   key,
		value: value,
	}
	q.len++
	if q.len == 1 {
		q.header = qe
		qe.next = qe
		qe.pre = qe
		return
	}
	end := q.header.pre
	qe.next = q.header
	end.next = qe
	q.header.pre = qe
	qe.pre = end

}

func (q *Queue) Pop() (key int, value int64) {

	if q.len <= 0 {
		return -1, -1
	}
	qe := q.header
	q.len--
	if q.len == 0 {
		q.header = nil
	} else {
		end := q.header.pre
		q.header = q.header.next
		q.header.pre = end
		end.next = q.header
	}
	return qe.key, qe.value
}

func Unweighted() {
	a := []int{2, 3, 2, 3, 2, 3, 2, 3, 2, 3}
	q := Queue{}
	path := map[int]int64{}
	q.Push(0, 0)
	v := len(a) - 1
	find := false
	for q.len >= 0 {
		if find {
			break
		}
		key, value := q.Pop()
		//fmt.Println(key, value)
		path[key] = value
		for i := 1; i <= a[key]; i++ {
			if key+1 >= len(a) {
				break
			}
			if key+i == v {
				fmt.Println(value + 1)
				find = true
				break
			}
			if _, ok := path[key+i]; !ok {
				q.Push(key+i, value+1)
			}
		}
	}
}
