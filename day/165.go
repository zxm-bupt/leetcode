package day

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1String := strings.Split(version1, ".")
	version2String := strings.Split(version2, ".")
	length := max(len(version2String), len(version1String))
	for i := 0; i < length; i++ {
		var version1Num int
		var version2Num int
		if i >= len(version1String) {
			version1Num = 0
		} else {
			version1Num, _ = strconv.Atoi(version1String[i])
		}

		if i >= len(version2String) {
			version2Num = 0
		} else {
			version2Num, _ = strconv.Atoi(version2String[i])
		}

		if version1Num > version2Num {
			return 1
		}
		if version1Num < version2Num {
			return -1
		}
	}
	return 0
}
