package main

import (
	"flag"
	"fmt"
	"math"
)

var num int64

func init() {
	flag.Int64Var(&num, "n", 0, "number")
	flag.Parse()
}

func main() {
	max := maxPrime(num)
	fmt.Println(max)
}

// 方式三：
// 递归实现：n和最大因子maxfactor含有相同的最大素数
func maxPrime(n int64) (max int64) {
	var i int64
	max = n
	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			max = maxPrime(n / i)
			break
		}
	}
	// fmt.Println(n, max)
	return
}

// 方式二：
// | o | | | | o | | | | o | | | | | | | | | | o | | |
// |表示不可被除尽的数字；
// o表示可被除尽的数字；
// |和o之间是成对存在的；
//
//func maxPrime(n int64) (max int64) {
//	if n <= 2 {
//		return
//	}
//
//	var i int64
//	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
//		if n%i == 0 {
//			if isPrimeFactors(n / i) {
//				max = n / i
//				break
//			}
//			if isPrimeFactors(i) {
//				max = i
//			}
//		}
//	}
//
//	return
//}

// 方式一：
// 从小到大循环，挨个判断数字是否是素数，如果是素数存入slice，最后取slice最大值
// 缺点：该方式如果数字特别大的话，执行非常慢
//func primeFactors(n int64) (pf []int64) {
//	if n <= 2 {
//		return
//	}
//	var i int64
//	for i = 2; i <= n/2; i++ {
//		if n%i == 0 && isPrimeFactors(i) {
//			pf = append(pf, i)
//		}
//	}
//	return
//}

// 计算n是否是素数
//func isPrimeFactors(n int64) bool {
//	if n <= 0 {
//		return false
//	} else if n <= 2 {
//		return true
//	}
//
//	var i int64
//	isPrime := true
//	for i = 2; i <= n/2; i++ {
//		if n%i == 0 {
//			isPrime = false
//			break
//		}
//	}
//	return isPrime
//}
