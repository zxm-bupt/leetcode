package day

import "strings"

func makeFancyString(s string) string {
	if len(s) <= 2 {
		return s
	}

	var builder strings.Builder
	builder.WriteByte(s[0])
	secondByte := s[0]
	var lastByte byte
	if len(s) > 1 {
		builder.WriteByte(s[1])
		lastByte = s[1]
	}

	for i := 2; i < len(s); i++ {
		if s[i] != lastByte || s[i] != secondByte {
			secondByte = lastByte
			lastByte = s[i]
			builder.WriteByte(s[i])
		}
	}

	return builder.String()
}

func MakeFancyString(s string) string {
	return makeFancyString(s)
}
