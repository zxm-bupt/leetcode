package day

func makeTheIntegerZero(num1 int, num2 int) int {
	for k := 0; k < 64; k++ {
		var value int64 = int64(num1 - k*num2)
		if int(value) < k {
			continue
		} else if k < bitCount(value) {
			continue
		}
		return k
	}
	return -1
}

// 计算一个数字有多少位二进制1的算法
// Brian Kernighan算法
func bitCount(n int64) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
