package day

func numWaterBottles(numBottles int, numExchange int) int {
	res := 0
	for numBottles >= numExchange {
		remainder := numBottles % numExchange
		numBottles = numBottles / numExchange
		res += numBottles * numExchange
		numBottles += remainder
	}
	res += numBottles
	return res
}
