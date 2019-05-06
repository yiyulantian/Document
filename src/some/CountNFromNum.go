package some

import "fmt"

//给定一个奇数n，可得到一个由从1到n的所有奇数所组成的数列，求这一数列中数字k所出现的总次数。例如当n=3，k=3时，可得到奇数列：1,3，其中有一个数字3，故可得1
/*思路：数位DP，dfs(offset,limit,dp)
深度优先搜索，使用dp记录已经计算过的数据（pos位，limit是false时，符合条件的数总个数match（奇数）和k出现的总次数count）
有一套模版，大致意思就是用dfs递归计算各位符合要求的数，其中都需要特别判断的是该位是否有限制；
例如21这个数，十位是0、1时个位可以取到0-9没有限制，十位是2时个位就有限制了，只能取0-1，这就是limit的作用；
*/
func CountNFromNum() {
	k := 3
	num := 35
	numArr := []int{}
	dpMap := make(map[int][]int)
	for num > 0 {
		numArr = append(numArr, num%10)
		num = num / 10
	}
	count, _ := dfs(len(numArr)-1, k, true, numArr, dpMap)

	fmt.Println(count)
}

func dfs(pos, t int, limit bool, numArr []int, dpMap map[int][]int) (count, match int) {
	if pos == -1 {
		return 0, 1
	}
	if !limit && dpMap[pos] != nil {
		return dpMap[pos][0], dpMap[pos][1]
	}
	max := 9
	if limit {
		max = numArr[pos]
	}

	for i := 0; i <= max; i++ {
		if pos == 0 && i%2 == 0 {
			continue
		}
		tcount, tmatch := dfs(pos-1, t, limit && i == max, numArr, dpMap)
		count += tcount
		match += tmatch
		if i == t {
			count += tmatch
		}
	}

	if !limit {
		dpMap[pos] = make([]int, 2)
		dpMap[pos][0] = count
		dpMap[pos][1] = match
	}
	return
}
