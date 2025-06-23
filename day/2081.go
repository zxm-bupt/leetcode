package day

// Day 2081: Palindromic Numbers
// https://adventofcode.com/2023/day/2081
func kMirror(k int, n int) int64 {
	left, count, ans := 1, 0, int64(0)
	for count < n {
		right := left * 10
		// op = 0 表示枚举奇数长度回文，op = 1 表示枚举偶数长度回文
		for op := 0; op < 2; op++ {
			// 枚举 i'
			for i := left; i < right && count < n; i++ {
				combined := int64(i)
				x := i
				if op == 0 {
					x = i / 10
				}
				for x > 0 {
					combined = combined*10 + int64(x%10)
					x /= 10
				}
				if isPalindrome(combined, k) {
					count++
					ans += combined
				}
			}
		}
		left = right
	}
	return ans
}

// Check if a number is a palindrome in a given base
func isPalindrome(x int64, k int) bool {
	length := -1
	digit := make([]int, 64) // Sufficient for base up to 64
	for x > 0 {
		length++
		digit[length] = int(x % int64(k))
		x /= int64(k)
	}
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		if digit[i] != digit[j] {
			return false
		}
	}
	return true
}
