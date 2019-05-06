package some

import "fmt"

//最长公共子序列(可以不连续)
//a，b两个字符串中的最长的公共子序列
//思路：用dp，设dp[i][j]表示以a[0:i],b[0:j]子字串中最长的公共子序列长度，那么：
//dp[i][j]:{a[i]==b[j]: dp[i-1][j-1]+1 ; a[i]!=b[j]:max(dp[i,j-1],dp[i-1,j])}

//如果简单点，求最长公共子串，即连续的串;则设dp[i][j]表示分别以a[i],b[j]为结尾的最长子串的长度，那么：
//dp[i][j]:{a[i]==b[j]: dp[i-1][j-1]+1 ; a[i]!=b[j]:0}

func LCS() {
	a := "adegftssaefclkif"
	b := "faafemoalshfsec"

	abyte := []byte(a)
	bbyte := []byte(b)

	dp := make([][]int, len(abyte)+1)
	for i := range dp {
		dp[i] = make([]int, len(bbyte)+1)
		dp[i][0] = 0
	}

	for i := range dp[0] {
		dp[0][i] = 0
	}

	max := 0
	for i := 0; i < len(abyte); i++ {
		for j := 0; j < len(bbyte); j++ {
			if abyte[i] == bbyte[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if max < dp[i+1][j+1] {
					max = dp[i+1][j+1]
				}
			} else {
				if dp[i+1][j] >= dp[i][j+1] {
					dp[i+1][j+1] = dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}

	fmt.Println(dp, max)
}

//最小字符串编辑长度
/*
对字符串中字符进行增、删、替换操作的次数称为编辑长度
a,b两个字符串，求使得a==b的最小编辑长度
思路：和上面求最大公共子序列差不多；dp[i][j]表示a[0:i]到b[0:j]的最小编辑距离，那么：
dp[i][j] = {a[i]==b[j]:dp[i-1][j-1] ; a[i]!=b[j]:min(dp[i][j-1],dp[i-i][j],dp[i-1][j-1])+1 }
*/
func EditStr() {
	a := "aabcd"
	b := "abe"
	abyte := []byte(a)
	bbyte := []byte(b)

	dp := make([][]int, len(abyte)+1)
	for i := range dp {
		dp[i] = make([]int, len(bbyte)+1)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 0; i < len(abyte); i++ {
		for j := 0; j < len(bbyte); j++ {
			if abyte[i] == bbyte[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				min := dp[i][j]
				if dp[i][j+1] < min {
					min = dp[i][j+1]
				} else if dp[i+1][j] < min {
					min = dp[i+1][j]
				}
				dp[i+1][j+1] = min + 1
			}
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[len(abyte)][len(bbyte)])
}
