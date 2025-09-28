package day

import "math"

func area(pointA, pointB, pointC []int) float64 {

	return 0.5 * math.Abs(float64((pointA[0]*pointB[1]+pointB[0]*pointC[1]+pointC[0]*pointA[1])-(pointB[0]*pointA[1]+pointC[0]*pointB[1]+pointA[0]*pointC[1])))

}

func LargestTriangleArea(points [][]int) float64 {
	return largestTriangleArea(points)
}

func largestTriangleArea(points [][]int) float64 {
	res := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {

				res = max(res, area(points[i], points[j], points[k]))

			}
		}
	}
	return res
}
