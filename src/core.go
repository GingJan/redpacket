package src

import (
	"math/rand"
	"sort"
)

// 二倍均值法，假设当前人均值是x，那么随机数取[1,2x)（单位是分）
func Avg2Times(totalAmount int, totalNum int, seed int64) []int {
	ret := make([]int, 0, totalNum)
	if totalAmount <= 0 || totalNum <= 0 {
		return ret
	}

	r := rand.New(rand.NewSource(seed))
	for i := totalNum; i >= 2; i-- {
		x := (totalAmount << 1) / i
		random := r.Intn(x)
		if random == 0 {
			random = 1
		}
		ret = append(ret, random)
		totalAmount -= random
	}
	ret = append(ret, totalAmount)

	return ret
}

//线段分割法。其思想是把红包抽象成一根在数轴上的线段，线段的起点是原点，终点是红包的总金额 totalAmount，然后把线段分成 totalNum 个部分，只需要生成 totalNum-1 个随机数，这些随机数的范围是 [ 1 , totalAmount )
func SplitLine(totalAmount int, totalNum int, seed int64) []int {
	ret := make([]int, 0, totalNum)
	if totalAmount <= 0 || totalNum <= 0 {
		return ret
	}

	src := rand.NewSource(seed)
	r := rand.New(src)
	m := make(map[int]struct{})
	for len(m) < totalNum - 1 {
		random := r.Intn(totalAmount)
		if random == 0 {
			random = 1
		}
		m[random] = struct{}{}
	}


	amounts := make([]int, 0, len(m))
	for a := range m {
		amounts = append(amounts, a)
	}

	sort.Ints(amounts)
	ret = append(ret, amounts[0])
	for i := 1; i < len(amounts); i++ {
		ret = append(ret, amounts[i] - amounts[i - 1])
	}
	ret = append(ret, totalAmount - amounts[len(amounts) - 1])

	return ret
}

// 每次抢到的金额=随机范围（0，M/N *2） M表示剩余红包金额，N表示剩余个数。这个公式保证了每次随机金额的平均值是相等的。
func AvgSplit(totalAmount int, totalNum int, seed int64) []int {
	ret := make([]int ,0 ,totalNum)
	if totalAmount <= 0 || totalNum <= 0 {
		return ret
	}

	restNum := totalNum
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < totalNum - 1; i++ {
		//保证金额范围是[1，剩余金额2倍），左闭右开
		random := r.Intn(totalAmount / restNum * 2 - 1) + 1
		if random == 0 {
			random = 1
		}
		totalAmount -= random
		restNum--
		ret = append(ret, random)
	}
	ret = append(ret, totalAmount)
	return ret
}