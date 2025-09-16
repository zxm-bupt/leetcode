package day

func ReplaceNonCoprimes(nums []int) []int {
	return replaceNonCoprimes(nums)
}
func replaceNonCoprimes(nums []int) []int {
	stack := []int{}
	for _, num := range nums {
		for len(stack) != 0 && gcd(num, stack[len(stack)-1]) != 1 {
			top := stack[len(stack)-1]
			num = lcm(num, top)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	return stack
}

func gcd(x, y int) int {
	maxV := max(x, y)
	minV := min(x, y)

	remainder := maxV % minV

	if remainder != 0 {
		return gcd(minV, remainder)
	} else {
		return minV
	}

}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
