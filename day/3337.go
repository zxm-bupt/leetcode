package day

func LengthAfterTransformations2(s string, t int, nums []int) int {
	mat := newMatrix(26)
	for i := 0; i < 26; i++ {
		for j := 1; j <= nums[i]; j++ {
			index := (i + j) % 26
			mat.a[index][i]++
		}
	}
	mat = mat.matrixQuickPow(t)
	res := 0
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			res += mat.a[i][j] * count[j]
			res %= mod
		}
	}
	return res
}

func change2(hash []int, nums []int) []int {
	res := make([]int, 26)
	for i := 0; i < 26; i++ {
		for j := 1; j <= nums[i]; j++ {
			index := (i + j) % 26
			res[index] += hash[i]
			res[index] %= mod
		}
	}

	return res
}

// 传统快速幂
func quickPow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
}

// 矩阵快速幂
type Matrix struct {
	n int
	a [][]int
}

func newMatrix(n int) *Matrix {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	return &Matrix{
		n: n,
		a: mat,
	}
}

func multiply(a, b *Matrix) *Matrix {
	res := newMatrix(a.n)
	for i := 0; i < a.n; i++ {
		for j := 0; j < b.n; j++ {
			for k := 0; k < a.n; k++ {
				res.a[i][j] = (res.a[i][j] + a.a[i][k]*b.a[k][j]) % mod
			}
		}
	}
	return res
}

// 自身相乘
func (m *Matrix) multiplySelf() {
	res := multiply(m, m)
	m.a = res.a
}

// 单位矩阵
func newIdentityMatrix(n int) *Matrix {
	identity := newMatrix(n)
	for i := 0; i < n; i++ {
		identity.a[i][i] = 1
	}
	return identity
}

// 矩阵快速幂
func (m *Matrix) matrixQuickPow(n int) *Matrix {
	res := newIdentityMatrix(m.n)
	for n > 0 {
		if n&1 == 1 {
			res = multiply(res, m)
		}
		m.multiplySelf()
		n >>= 1
	}
	return res
}
