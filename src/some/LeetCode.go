package some

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

func LeetCodeMain() {
	t := canMeasureWater(2, 6, 5)
	fmt.Println(t)
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	tmp := []byte(s)
	b := make([]byte, len(tmp)*2+1)
	b[0] = '#'

	for k, v := range tmp {
		b[2*k+1] = v
		b[2*k+2] = '#'
	}

	dp := make([]int, len(b))
	dp[0] = 1
	maxmidd := 0
	maxright := 0
	for i := 1; i < len(b); i++ {
		//i在最长回文串内
		if i < maxright {
			j := 2*maxmidd - i
			if i+dp[j]-1 < maxright {
				dp[i] = dp[j]
			} else {
				ti := maxright
				for ; ti < len(b) && (2*i-ti) >= 0; ti++ {
					if b[2*i-ti] != b[ti] {
						break
					}

				}

				ti--
				dp[i] = ti - i + 1
			}
		} else {
			ti := i + 1
			for ; ti < len(b) && (2*i-ti) >= 0; ti++ {
				if b[2*i-ti] != b[ti] {
					break
				}
			}
			ti--
			dp[i] = ti - i + 1
		}
		if dp[i] > dp[maxmidd] {
			maxmidd = i
			maxright = i + dp[i] - 1
		}
	}
	// fmt.Println(string(b))
	// fmt.Println(dp)
	t := ""
	for i := 2*maxmidd - maxright; i <= maxright; i++ {
		if b[i] == '#' {
			continue
		}
		t = t + string(b[i])
	}
	return t
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := []byte(s)

	l := len(b)
	eb := make([]byte, l)
	of := -1
	for i := 1; i <= numRows; i++ {
		if i == 1 || i == numRows {
			for j := i - 1; j < l; {
				of++
				eb[of] = b[j]
				j = j + 2*numRows - 2
			}
		} else {
			for j := i - 1; j < l; {
				of++
				eb[of] = b[j]
				tmpj := j + 2*(numRows-i)
				//fmt.Println(tmpj, j, numRows, i)
				if tmpj < l {
					of++
					eb[of] = b[tmpj]
				}
				j = j + 2*numRows - 2
			}
		}
	}

	t := string(eb)
	return t
}

//滑动窗口
func maxArea(height []int) int {
	max, i, j := 0, 0, len(height)-1
	for i < j {
		min := height[i]
		l := j - i
		if min > height[j] {
			min = height[j]
			j--
		} else {
			i++
		}
		if max < min*l {
			max = min * l
		}
	}
	return max
}

func intToRoman(num int) string {
	t := ""
	tm := num / 1000
	for i := 0; i < tm; i++ {
		t = t + "M"
	}
	num = num % 1000

	tm = num / 100
	if tm == 9 {
		t = t + "CM"
	} else if tm >= 5 {
		t = t + "D"
		for i := 0; i < tm-5; i++ {
			t = t + "C"
		}
	} else if tm == 4 {
		t = t + "CD"
	} else {
		for i := 0; i < tm; i++ {
			t = t + "C"
		}
	}
	num = num % 100

	tm = num / 10
	if tm == 9 {
		t = t + "XC"
	} else if tm >= 5 {
		t = t + "L"
		for i := 0; i < tm-5; i++ {
			t = t + "X"
		}
	} else if tm == 4 {
		t = t + "XL"
	} else {
		for i := 0; i < tm; i++ {
			t = t + "X"
		}
	}
	num = num % 10

	tm = num / 1
	if tm == 9 {
		t = t + "IX"
	} else if tm >= 5 {
		t = t + "V"
		for i := 0; i < tm-5; i++ {
			t = t + "I"
		}
	} else if tm == 4 {
		t = t + "IV"
	} else {
		for i := 0; i < tm; i++ {
			t = t + "I"
		}
	}
	return t
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var i, j *ListNode
	i = head
	t := 0
	for i.Next != nil && t < n {
		i = i.Next
		t++
	}
	if t < n {
		return head.Next
	}

	j = head
	for i.Next != nil {
		i = i.Next
		j = j.Next
	}
	j.Next = j.Next.Next
	return head
}

func generateParenthesis(n int) []string {
	ln, rn := n, n

	return sgenerateParenthesisPartB(ln-1, rn, "(")
}
func sgenerateParenthesisPartB(ln, rn int, s string) []string {

	if ln == 0 && rn == 1 {
		return []string{s + ")"}
	}

	ts := []string{}
	if ln > 0 {
		ts = sgenerateParenthesisPartB(ln-1, rn, s+"(")
	}
	if rn > 0 && rn > ln {
		ts = append(ts, sgenerateParenthesisPartB(ln, rn-1, s+")")...)
	}
	return ts
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			j := i + 1

			for ; j < len(nums); j++ {
				if nums[j] <= nums[i-1] {
					break
				}
			}
			nums[j-1], nums[i-1] = nums[i-1], nums[j-1]
			break
		}
	}

	for j := len(nums) - 1; j >= i; {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
	//fmt.Println(nums)

}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := 0
	for i := 0; i <= max; i++ {
		tmax := i + nums[i]
		if max < tmax {
			if tmax >= len(nums)-1 {
				return true
			}
			max = tmax
		}
	}
	return false
}

//Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	mergePartB(intervals[:], 0, len(intervals)-1)
	t := make([]Interval, 0)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= tmp.End {
			if intervals[i].End > tmp.End {
				tmp.End = intervals[i].End
			}
		} else {
			t = append(t, tmp)
			tmp = intervals[i]
		}
	}
	t = append(t, tmp)
	return t

}

func mergePartB(intervals []Interval, b, e int) {
	if b >= e {
		return
	}
	t := intervals[b]
	tb, te := b, e
	for b < e {
		for b < e {
			tmp := intervals[e]
			if tmp.Start < t.Start {
				intervals[b] = intervals[e]

				break
			}
			e--
		}

		for b < e {
			tmp := intervals[b]
			if tmp.Start > t.Start {
				intervals[e] = intervals[b]
				break
			}
			b++
		}
	}

	intervals[e] = t
	mergePartB(intervals[:], tb, e-1)

	mergePartB(intervals[:], e+1, te)
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, m)
		dp[k][0] = 1
		if k == 0 {
			for j := range dp[0] {
				dp[0][j] = 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[n-1][m-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n <= 0 {
		return 0
	}
	m := len(obstacleGrid[0])

	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, m)

		if (k == 0 && obstacleGrid[k][0] == 0) || (k > 0 && obstacleGrid[k][0] == 0 && dp[k-1][0] == 1) {
			dp[k][0] = 1
		} else {
			dp[k][0] = 0
		}

		if k == 0 {
			for j := range dp[0] {
				if (j == 0 && obstacleGrid[0][j] == 0) || (j > 0 && obstacleGrid[0][j] == 0 && dp[0][j-1] == 1) {
					dp[0][j] = 1
				} else {
					dp[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = 0
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if obstacleGrid[i][j-1] == 0 {
				dp[i][j] = dp[i][j] + dp[i][j-1]
			}
			if obstacleGrid[i-1][j] == 0 {
				dp[i][j] = dp[i][j] + dp[i-1][j]
			}
		}
	}
	return dp[n-1][m-1]
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	t := searchMatrixPartB(matrix, 0, 0, len(matrix)-1, -1, 0, 0, target)
	if t < 0 || t > len(matrix)-1 {
		return false
	}
	if matrix[t][0] == target {
		return true
	} else if matrix[t][0] > target {
		if t-1 >= 0 && matrix[t-1][0] < target {
			t--
		} else {
			return false
		}
	}
	//fmt.Println(t)
	et := searchMatrixPartB(matrix, -1, 0, 0, t, 0, len(matrix[0])-1, target)
	if et < 0 || et > len(matrix[0])-1 {
		return false
	}
	//fmt.Println(et)
	if matrix[t][et] == target {
		return true
	}
	return false
}

func searchMatrixPartB(d [][]int, y, xb, xe, x, yb, ye, target int) int {
	//fmt.Println(y, xb, xe, x, yb, ye)
	if x != -1 {
		if yb >= ye {
			return ye
		}
		midd := (yb + ye) / 2
		if target == d[x][midd] {
			return midd
		} else if target < d[x][midd] {
			t := searchMatrixPartB(d, -1, 0, 0, x, yb, midd-1, target)
			return t
		} else {
			t := searchMatrixPartB(d, -1, 0, 0, x, midd+1, ye, target)
			return t
		}
	} else {
		if xb >= xe {
			return xe
		}
		midd := (xb + xe) / 2
		if target == d[midd][y] {
			return midd
		} else if target < d[midd][y] {
			t := searchMatrixPartB(d, y, xb, midd-1, -1, 0, 0, target)
			return t
		} else {
			t := searchMatrixPartB(d, y, midd+1, xe, -1, 0, 0, target)
			return t
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	d := []byte(s)
	if len(d) == 0 {
		return 0
	}
	max := 1
	b, e := 0, 1
	findmap := map[byte]int{}
	findmap[d[0]] = 0
	for b < len(d) && e < len(d) {
		find, ok := findmap[d[e]]
		if !ok {
			findmap[d[e]] = e
			if (e - b + 1) > max {
				max = e - b + 1
			}
			e++
		} else {
			for ; b < find+1; b++ {
				delete(findmap, d[b])
			}
		}
	}
	return max
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	t := []int{}
	sl := 0
	this := root
	for this != nil || sl > 0 {
		if this == nil {
			this = stack[sl-1]
			sl--
			t = append(t, this.Val)
			//fmt.Println(t, sl)
			this = this.Right
			continue
		}

		for this != nil {
			if len(stack) <= sl {
				stack = append(stack, this)
				sl++
				//fmt.Println(sl)
			} else {
				stack[sl] = this
				sl++
				//fmt.Println("---", sl)
			}
			this = this.Left
		}

	}
	return t
}

func isValidBST(root *TreeNode) bool {
	sortTree := inorderTraversal(root)

	left := ^(int(^uint(0) >> 1))
	for _, v := range sortTree {
		if left >= v {
			return false
		}
		left = v
	}
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]int{}
	for _, v := range wordList {
		wordMap[v] = 1
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	if beginWord == endWord {
		return 1
	}

	// fmt.Println(wordMap)

	matchMap := map[string]int{}
	stack := list.New()
	stack.PushBack(beginWord)
	delete(wordMap, beginWord)
	matchMap[beginWord] = 1

	for i := stack.Front(); i != nil; i = i.Next() {
		this := []byte(i.Value.(string))
		thiscount := matchMap[i.Value.(string)]
		for word := range wordMap {
			w := []byte(word)

			for k, v := range this {
				if w[k] == v {
					continue
				}
				this[k] = w[k]
				if _, ok := wordMap[string(this)]; ok {
					if string(this) == endWord {
						return thiscount + 1
					}
					stack.PushBack(string(this))
					matchMap[string(this)] = thiscount + 1
					delete(wordMap, string(this))
				}
				this[k] = v
			}

		}

	}
	return 0
}

func partition(s string) [][]string {
	all := make([][]string, 0)
	tmp := make([]string, 0)
	all = partitionPartB(s, 0, all, tmp)
	return all
}

func partitionPartB(s string, l int, all [][]string, tmp []string) [][]string {
	if len(s) == l {
		all = append(all, tmp)
		return all
	}

	for i := l; i <= len(s)-1; i++ {
		b := l
		e := i
		for b <= e {
			if s[b] != s[e] {
				break
			}
			b++
			e--
		}
		if b < e {
			continue
		}
		tmp = append(tmp, s[l:i+1])

		all = partitionPartB(s, i+1, all, tmp)
		tmp = append([]string{}, tmp[:len(tmp)-1]...)
	}
	return all
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]int{}
	for _, v := range wordDict {
		wordMap[v] = 1
	}
	matchMap := map[int]bool{}
	return wordBreakPartB([]byte(s), 0, wordMap, matchMap)
}

func wordBreakPartB(sb []byte, b int, wordMap map[string]int, matchMap map[int]bool) bool {
	if b == len(sb) {
		return true
	}
	res := false
	for i := b; i < len(sb); i++ {
		if _, ok := wordMap[string(sb[b:i+1])]; ok {
			if v, ok := matchMap[i+1]; ok {
				res = v
			} else {
				res = wordBreakPartB(sb, i+1, wordMap, matchMap)
				matchMap[i+1] = res
			}
			if res {
				return res
			}
		}
	}
	return res
}

func reverseWords(s string) string {
	sb := []byte(s)
	if len(sb) == 0 {
		return ""
	}
	ko := false
	l := len(sb)
	for i := 0; i < l; {
		if i == 0 && sb[i] == ' ' {
			sb = sb[1:]
			l--
		} else if sb[i] == ' ' && ko {
			sb = append(sb[0:i], sb[i+1:]...)
			l--
		} else if sb[i] == ' ' && !ko {
			i++
			ko = true
		} else {
			if ko {
				ko = false
			}
			i++
		}
	}
	if len(sb) == 0 {
		return ""
	}
	if sb[len(sb)-1] == ' ' {
		sb = sb[:len(sb)-1]
	}
	if len(sb) == 0 {
		return ""
	}
	//fmt.Println(string(sb))

	i, j := 0, len(sb)-1
	for i < j {
		sb[i], sb[j] = sb[j], sb[i]
		i++
		j--
	}
	//fmt.Println(string(sb))
	block := 0
	for of := 0; of < len(sb); of++ {
		if sb[of] == ' ' || of == len(sb)-1 {
			i := block
			j := of - 1
			if of == len(sb)-1 {
				j = of
			}
			for i < j {
				sb[i], sb[j] = sb[j], sb[i]
				i++
				j--
			}
			block = of + 1
		}
	}
	return string(sb)
}

func findMin(nums []int) int {
	i, j := 0, len(nums)-1

	for i < j-1 {
		midd := (i + j) / 2
		//fmt.Println(nums[midd])
		if nums[midd] > nums[i] && nums[midd] < nums[j] {
			return nums[i]
		} else if nums[midd] < nums[i] && nums[midd] > nums[j] {
			return nums[j]
		} else if nums[midd] > nums[i] {
			i = midd
		} else if nums[midd] < nums[j] {
			j = midd
		}
		//fmt.Println(i, j)
	}
	if nums[j] < nums[i] {
		return nums[j]
	}
	return nums[i]
}

func compareVersion(version1 string, version2 string) int {
	v1 := []byte(version1)
	v2 := []byte(version2)
	i, j := 0, 0

	for i < len(v1) || j < len(v2) {
		tmp1 := ""
		tmp2 := ""
		f1 := true
		f2 := true
		for ; i < len(v1); i++ {
			if f1 && v1[i] == '0' {
				continue
			} else if v1[i] != '.' {
				f1 = false
				tmp1 = tmp1 + string(v1[i])
			} else {
				i++
				break
			}
		}
		for ; j < len(v2); j++ {
			if f2 && v2[j] == '0' {
				continue
			} else if v2[j] != '.' {
				f2 = false
				tmp2 = tmp2 + string(v2[j])
			} else {
				j++
				break
			}
		}
		fnum1 := 0
		fnum2 := 0
		if tmp1 != "" {
			fnum1, _ = strconv.Atoi(tmp1)
		}
		if tmp2 != "" {
			fnum2, _ = strconv.Atoi(tmp2)
		}
		if fnum1 > fnum2 {
			return 1
		} else if fnum1 < fnum2 {
			return -1
		}
	}
	return 0
}

func largestNumber(nums []int) string {
	ns := make([]string, len(nums))
	for k, v := range nums {
		ns[k] = strconv.Itoa(v)
	}
	//fmt.Println(ns)
	ns = ladderLengthPartB(ns, 0, len(ns)-1)
	//fmt.Println(ns)
	s := ""
	if len(ns) > 0 && ns[0] == "0" {
		return "0"
	}
	for _, v := range ns {
		s += v
	}

	return s
}

func ladderLengthPartB(d []string, b, e int) []string {
	if b >= e {
		return d
	}
	ob, oe := b, e
	kv := d[b]
	for b < e {
		for b < e {
			if largestNumberPartC(d[e], kv) {
				d[b] = d[e]
				break
			} else {
				e--
			}
		}
		for b < e {
			if largestNumberPartC(kv, d[b]) {
				d[e] = d[b]
				break
			} else {
				b++
			}
		}
	}

	d[b] = kv
	d = ladderLengthPartB(d, ob, b-1)
	d = ladderLengthPartB(d, b+1, oe)
	return d
}

func largestNumberPartC(a, b string) bool {
	if a+b > b+a {
		return true
	}
	return false
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 0 || k == 0 {
		return false
	}
	matchMap := map[int]int{nums[0]: 0}

	for i := 1; i < len(nums); i++ {
		for d := range matchMap {
			diff := d - nums[i]
			if diff < 0 {
				diff = 0 - diff
			}
			if diff <= t {
				return true
			}
		}

		matchMap[nums[i]] = i

		if i >= k {
			if matchMap[nums[i-k]] == i-k {
				delete(matchMap, nums[i-k])
			}
		}

	}
	return false
}

func nthUglyNumber(n int) int {

	q := []int{0, 0, 0}
	d := make([]int, n)
	d[0] = 1
	for i := 1; i < n; i++ {
		a := d[q[0]] * 2
		b := d[q[1]] * 3
		c := d[q[2]] * 5
		min := a
		if min > b {
			min = b
		}
		if min > c {
			min = c
		}
		if min == a {
			q[0]++
		}
		if min == b {
			q[1]++
		}
		if min == c {
			q[2]++
		}
		d[i] = min
	}
	return d[n-1]
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {
		thismax := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && thismax < dp[j] {
				thismax = dp[j]
			}
		}
		dp[i] = thismax + 1
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Println(dp)
	return max
}

func bulbSwitch(n int) int {
	x := math.Sqrt(float64(n))
	return int(x)
}

func increasingTriplet(nums []int) bool {

	if len(nums) < 3 {
		return false
	}
	pre := nums[0]
	min1 := -1
	min2 := -1
	frist := true
	for i := 1; i < len(nums); i++ {
		if nums[i] <= pre {
			pre = nums[i]
			continue
		} else {

			if !frist && ((min1 < min2 && min2 < nums[i]) || (min1 == min2 && min1 < pre && min1 < nums[i])) {
				return true
			} else {
				if frist {
					frist = false
					min1 = pre
					min2 = nums[i]
					continue
				}
				if min1 > pre {
					min1 = pre
				} else if min2 > pre {
					min2 = pre
				}

				if min2 > nums[i] {
					min2 = nums[i]
				}

				pre = nums[i]

			}
		}
	}
	return false
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	d := [][]int{}
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return d
	}
	if len(nums1)*len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	of := make([]int, len(nums1))
	for len(d) < k {
		min := 2 << 31
		minof := 0
		for i := 0; i < len(nums1); i++ {
			if of[i] >= len(nums2) {
				continue
			}
			if nums1[i]+nums2[of[i]] < min {
				min = nums1[i] + nums2[of[i]]
				minof = i
			}
		}
		//fmt.Println(minof, len(of))
		d = append(d, []int{nums1[minof], nums2[of[minof]]})
		of[minof]++
	}
	return d
}

func lexicalOrder(n int) []int {
	d := make([]int, n)
	of := 0
	tmp := 1
	for {
		for tmp <= n {
			d[of] = tmp
			of++
			if of == n {
				return d
			}
			tmp = tmp * 10
		}
		tmp = tmp / 10
		//fmt.Println("----", tmp)
		for j := 1; j <= 9-(tmp%10); j++ {
			if tmp+j <= n {
				d[of] = tmp + j
				of++
				if of == n {
					return d
				}
			} else {
				break
			}
		}
		//fmt.Println(tmp)
		if tmp < 10 {
			tmp++
		} else {
			tmp = tmp / 10
			for tmp%10 == 9 {
				tmp = tmp / 10
			}
			tmp++
		}
	}
	return d
}

func isSubsequence(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	if len(sb) > len(tb) || len(tb) == 0 {
		return false
	}
	of := 0
	for _, v := range sb {
		j := of
		for ; j < len(tb); j++ {
			if tb[j] == v {
				of = j + 1
				break
			}
		}

		if j == len(tb) {
			return false
		}
	}
	return true
}

func numMatchingSubseq(S string, words []string) int {
	sb := []byte(S)
	d := map[byte][]int{}
	for k, v := range sb {
		if _, ok := d[v]; ok {
			d[v] = append(d[v], k)
		} else {
			d[v] = []int{k}
		}
	}

	n := 0
	for _, s := range words {
		sb := []byte(s)
		t := 0
		for _, v := range sb {
			if _, ok := d[v]; ok {
				//fmt.Println(t, d[v])
				t = numMatchingSubseqPartB(t, d[v])
				//fmt.Println(t)
				if t == -1 {
					break
				}
				t++
			} else {
				t = -1
				break
			}
		}
		if t != -1 {
			n++
		}

	}
	return n
}

func numMatchingSubseqPartB(t int, d []int) int {
	b, e := 0, len(d)-1
	for b < e {
		midd := (b + e) / 2
		if d[midd] < t {
			b = midd + 1
		} else if d[midd] > t {
			e = midd - 1
		} else {
			return d[midd]
		}
	}

	if d[b] >= t {
		return d[b]
	} else if d[b] < t {
		if b+1 == len(d) {
			return -1
		}
		return d[b+1]
	}
	return -1
}

func removeKdigits(num string, k int) string {
	sm := []byte(num)
	rm := []byte{}

	thisk := k + 1
	for i := 0; i < len(sm); {
		minof := i
		min := sm[i]
		thisb := i
		if len(sm[i:]) < thisk {
			break
		}
		for j := 1; j <= thisk && i < len(sm); j++ {
			if sm[i] < min {
				minof = i
				min = sm[i]
			}
			i++
		}

		thisk = thisk - (minof - thisb)
		i = minof + 1
		rm = append(rm, sm[minof])
		if thisk == 0 {
			rm = append(rm, sm[i:]...)
			break
		}

	}
	b := 0
	for ; b < len(rm); b++ {
		if rm[b] != '0' {
			break
		}
	}
	if b == len(rm) {
		return "0"
	} else {
		return string(rm[b:])
	}

}

func longestSubstring(s string, k int) int {
	return longestSubstringPartB([]byte(s), k)
}
func longestSubstringPartB(d []byte, k int) int {

	max := 0
	cmap := map[byte]int{}
	for _, db := range d {
		cmap[db] = cmap[db] + 1
	}

	match := [][]byte{}
	offset := 0
	for i, db := range d {
		if cmap[db] < k {
			match = append(match, d[offset:i])
			offset = i + 1
		}
	}
	if offset != 0 && offset < len(d) {
		match = append(match, d[offset:len(d)])
	}
	if len(match) > 0 {
		for _, matchd := range match {
			tmp := longestSubstringPartB(matchd, k)
			if max < tmp {
				max = tmp
			}
		}
	} else {
		if max < len(d) {
			max = len(d)
		}
	}
	return max
}

func findDuplicates(nums []int) []int {
	match := []int{}
	for _, v := range nums {
		if v < 0 {
			v = 0 - v
		}
		if nums[v-1] < 0 {
			match = append(match, v)
		} else {
			nums[v] = 0 - nums[v-1]
		}
	}
	return match
}

func nextGreaterElements(nums []int) []int {
	target := make([]int, len(nums))
	taroffset := make([]int, len(nums))
	b := 0
	e := len(nums)
	ob := -1
	for b < e && ob < b {
		ob = b
		b = nextGreaterElementsPartB(nums, ob, target, taroffset)
	}

	return target
}

func nextGreaterElementsPartB(nums []int, k int, target, taroffset []int) int {
	nl := len(nums)
	offset := (k + 1) % nl
	target[k] = -1
	taroffset[k] = -1
	for {
		if offset == -1 {
			return -1
		}
		if nums[k] < nums[offset] {
			target[k] = nums[offset]
			taroffset[k] = offset
			return offset
		} else if taroffset[offset] > 0 {
			offset = taroffset[offset]
		} else if taroffset[offset] == 0 {
			target[offset] = -1
			taroffset[offset] = -1
			next := nextGreaterElementsPartB(nums, offset, target, taroffset)
			offset = next
		} else if taroffset[offset] == -1 {
			return -1
		}
	}
}

func findLUSlength(strs []string) int {
	mStrs := [11]map[string]int{}
	for i := 0; i <= 10; i++ {
		mStrs[i] = make(map[string]int, 0)
	}
	for _, s := range strs {
		mStrs[len(s)][s]++
	}
	for i := 10; i > 0; i-- {
		for str, count := range mStrs[i] {
			if count == 1 {
				return len(str)
			}
			for j := 0; j < len(str); j++ {
				mStrs[i-1][str[:j]+str[j+1:]] += 2
			}
		}
	}
	return -1
}

func numFriendRequests(ages []int) int {
	ageCount := make([]int, 121)
	for _, age := range ages {
		ageCount[age]++
	}
	allcount := 0
	for i, iage := range ageCount {
		if iage == 0 {
			continue
		}
		for j, jage := range ageCount {
			if jage == 0 {
				continue
			}
			if ((i/2 + 7) >= j) || (j > i) || (j > 100 && j < 100) {
				continue
			}
			if i == j {
				allcount += (iage * (jage - 1))
			} else {
				allcount += (iage * jage)
			}
		}
	}
	return allcount
}
