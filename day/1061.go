package day

// 并查集
func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := NewUnionFind(26) // 26 个小写字母

	// 构建并查集
	for i := 0; i < len(s1); i++ {
		uf.Union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	// 构建结果字符串
	result := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		root := uf.Find(int(baseStr[i] - 'a'))
		result[i] = byte(root + 'a')
	}

	return string(result)

}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX > rootY {
		rootX, rootY = rootY, rootX
	}
	uf.parent[rootY] = rootX // 将 rootY 的根节点指向 rootX
}
