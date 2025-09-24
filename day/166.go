package day

import (
	"strconv"
	"strings"
)

func FractionToDecimal(numerator int, denominator int) string {
	return fractionToDecimal(numerator, denominator)
}

func fractionToDecimal(numerator int, denominator int) string {
	var res strings.Builder
	var remainder int
	var integer int
	// 符号部分
	if numerator*denominator < 0 {
		res.WriteRune('-')
	}

	if numerator < 0 {
		numerator = -1 * numerator
	}

	if denominator < 0 {
		denominator = -1 * denominator
	}
	// 整数部分
	if denominator > numerator {
		integer = 0
		remainder = numerator
	} else {
		integer = numerator / denominator
		remainder = numerator % denominator
	}

	res.WriteString(strconv.Itoa(integer))
	if remainder != 0 {
		res.WriteRune('.')
	}

	// 小数部分
	hash := make(map[int]int)
	stack := []int{}
	index := 0
	for remainder != 0 {
		if _, ok := hash[remainder]; !ok {
			hash[remainder] = index
		} else {
			index = hash[remainder]
			break
		}
		remainder = remainder * 10
		quotient := remainder / denominator
		stack = append(stack, quotient)
		remainder = remainder % denominator
		index++
	}
	for i, num := range stack {
		if i == index {
			res.WriteRune('(')
		}
		res.WriteString(strconv.Itoa(num))
	}
	if remainder != 0 {
		res.WriteRune(')')
	}

	return res.String()
}
