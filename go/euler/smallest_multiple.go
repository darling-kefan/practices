package main

import (
	"fmt"
	"math"
	_ "os"
	_ "sort"
	_ "strconv"
)

// 求1,2,3,4,...,20的最小公倍数
func main() {
	n := 20
	args := make([]int, n)
	for i := 1; i <= n; i++ {
		args[i-1] = i
	}
	fmt.Println(SmallestMultiple(args...))
}

// 求最小公倍数
//
// 参考：https://baike.baidu.com/item/%E6%9C%80%E5%A4%A7%E5%85%AC%E7%BA%A6%E6%95%B0
//
// “质因数分解法”：
// 把每个数分别分解质因数，再把各数中的全部公有质因数提取出来连乘，所得的积就是这几个数的最大公约数。
// 把每个数分别分解质因数，再把各数中的全部公有的质因数和独有的质因数提取出来连乘，所得的积就是这几个数的最小公倍数。
//
// 例如：
// 求24和60的最大公约数，先分解质因数，得24=2×2×2×3, 60=2×2×3×5, 24与60的全部公有的质因数是2,2,3；
// 最大公约数：(24,60) = 2×2×3 = 12
// 最小公倍数：[24,60] = 2×2×3×2×5 = 120
//
func SmallestMultiple(args ...int) int {
	mutiple := 1
	for _, v := range args {
		mutiple = smallestMultiple(mutiple, v)
	}
	return mutiple
}

// 计算两个数的最小公倍数
// [2, 2, 3, 4, 5] 和 [2, 2, 4, 6] 的并集应该为 [2, 2, 4, 3, 5, 6]
func smallestMultiple(x int, y int) int {
	// x,y分解质因数
	xs, ys := decomposeFactors(x), decomposeFactors(y)

	// 存储ys里等于xs元素的索引值
	var xKeys []int
	for _, vx := range xs {
		for k, vy := range ys {
			isX := false
			for _, mk := range xKeys {
				if mk == k {
					isX = true
					break
				}
			}
			if !isX && vx == vy {
				xKeys = append(xKeys, k)
				break
			}
		}
	}

	xys := xs
	for k, v := range ys {
		isX := false
		for _, xk := range xKeys {
			if xk == k {
				isX = true
				break
			}
		}
		if !isX {
			xys = append(xys, v)
		}
	}
	//sort.Ints(xys)

	multiple := 1
	for _, v := range xys {
		multiple = multiple * v
	}

	return multiple
}

// 分解质因数
func decomposeFactors(n int) (factors []int) {
	if n < 2 {
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			nextFacts := decomposeFactors(n / i)
			factors = append(factors, nextFacts...)
			break
		}
	}
	if len(factors) == 0 {
		factors = append(factors, n)
	}

	return
}

// 此方法不够通用
//func smallestMultiple20() int64 {
//
//	fmt.Println(decomposeFactors(19))
//
//	primeNums := []int{2, 3, 5, 7, 11, 13, 17, 19}
//	nums12 := []int{2, 2, 3}
//	nums14 := []int{2, 7}
//	nums16 := []int{2, 2, 2, 2}
//	nums18 := []int{2, 3, 3}
//	nums20 := []int{2, 2, 5}
//
//	allSet := smallestMultiSet(primeNums, nums12)
//	allSet = smallestMultiSet(allSet, nums14)
//	allSet = smallestMultiSet(allSet, nums16)
//	allSet = smallestMultiSet(allSet, nums18)
//	allSet = smallestMultiSet(allSet, nums20)
//	fmt.Println(allSet)
//
//	var multi int64 = 1
//	for _, v := range allSet {
//		multi = multi * int64(v)
//	}
//	return multi
//}
//
//// 计算两个数的所有共有质因数和独有质因数的集合，用于计算
//func smallestMultiSet(setA []int, setB []int) (set []int) {
//	var delBKeys []int
//	for _, va := range setA {
//		for k, vb := range setB {
//			isDel := false
//			for _, kv := range delBKeys {
//				if kv == k {
//					isDel = true
//				}
//			}
//			if !isDel && va == vb {
//				delBKeys = append(delBKeys, k)
//				break
//			}
//		}
//	}
//	// fmt.Println(delBKeys)
//
//	set = setA
//	for k, v := range setB {
//		isDel := false
//		for _, dk := range delBKeys {
//			if dk == k {
//				isDel = true
//			}
//		}
//		if !isDel {
//			set = append(set, v)
//		}
//	}
//	sort.Ints(set)
//
//	return
//}

// 方式一：
// 逻辑简单，但是效率太低～
//func smallestMultiple(base int64) int64 {
//	var i, smallest int64
//	var isStop bool
//
//	smallest = 1
//	for {
//		isStop = true
//		for i = base; i >= 2; i-- {
//			if smallest%i != 0 {
//				isStop = false
//				break
//			}
//		}
//		if isStop {
//			break
//		}
//		smallest++
//
//		fmt.Println(smallest)
//	}
//	return smallest
//}
