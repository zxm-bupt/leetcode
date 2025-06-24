package day

func findKDistantIndices(nums []int, key int, k int) []int {
	result := []int{}
	visted := make([]bool, len(nums))
	for i, num := range nums {
		if num == key {
			for j := max(0, i-k); j <= min(len(nums)-1, i+k); j++ {
				if !visted[j] {
					visted[j] = true
					result = append(result, j)
				}
			}
		}
	}
	return result
}
