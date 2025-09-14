package day

import "container/heap"

type ByteHeap []byte

func (h ByteHeap) Len() int           { return len(h) }
func (h ByteHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h ByteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ByteHeap) Push(x interface{}) {
	*h = append(*h, x.(byte))
}

func (h *ByteHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortVowels(s string) string {
	// 判断是否为元音字母
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	// 收集所有元音字母并排序
	var vowels []byte
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	// 使用堆对元音字母按ASCII序排序
	h := (*ByteHeap)(&vowels)
	heap.Init(h)

	// 重新构造字符串
	result := []byte(s)
	for i := 0; i < len(result); i++ {
		if isVowel(result[i]) {
			result[i] = heap.Pop(h).(byte)
		}
	}

	return string(result)
}
