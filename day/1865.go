package day

type FindSumPairs struct {
	nums1 []int
	nums2 []int

	countNums1 map[int]int // Optional: to optimize counting pairs
	countNums2 map[int]int // Counts of elements in nums1
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {

	countNums1 := make(map[int]int)
	for _, num := range nums1 {
		countNums1[num]++
	}
	countNums2 := make(map[int]int)
	for _, num := range nums2 {
		countNums2[num]++
	}
	return FindSumPairs{
		nums1:      nums1,
		nums2:      nums2,
		countNums1: countNums1,
		countNums2: countNums2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.countNums2[this.nums2[index]]--
	this.nums2[index] += val
	this.countNums2[this.nums2[index]]++
	// Note: The problem does not specify what to do with the nums1 array,
	// so we assume it remains unchanged.
	// If nums1 needs to be updated, you can implement that logic here.
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0
	for num, count1 := range this.countNums1 {
		target := tot - num
		if count2, exists := this.countNums2[target]; exists {
			count += count1 * count2
		}
	}
	return count
}
