package day

func TriangleType(nums []int) string {
	if !isTriangle(nums) {
		return "none"
	}
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	}
	if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		return "isosceles"
	}
	return "scalene"

}

func isTriangle(nums []int) bool {
	maxV := max(nums[0], max(nums[1], nums[2]))
	others := 0
	visied := 0
	for _, v := range nums {
		if v == maxV && visied == 0 {
			visied = 1
			continue
		}
		if v <= maxV {
			others += v
		}
	}
	return maxV < others

}
