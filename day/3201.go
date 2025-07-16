package day

func maximumLength(nums []int) int {
	len1 := 0
	len2 := 0
	len3 := 0
	len4 := 0
	for _, num := range nums {
		if num%2 == 1 {
			len4++
			if len1%2 == 0 {
				len1++
			}
			if len2%2 == 1 {
				len2++
			}
		} else {
			len3++
			if len1%2 == 1 {
				len1++
			}
			if len2%2 == 0 {
				len2++
			}
		}

	}

	return max(len1, max(len2, max(len3, len4)))
}

func MaximumLength(nums []int) int {
	return maximumLength(nums)
}
