package main

import (
	"fmt"
)

type TreeNode struct {
	leftChild, rightChild int
	data                  int
}

type Tree []*TreeNode

// 根节点到叶子节点的所有路径
var paths [][]*TreeNode

func findPaths(tn *TreeNode, list []*TreeNode) {
	list = append(list, tn)

	if tn.leftChild == 0 && tn.rightChild == 0 {
		copyList := make([]*TreeNode, len(list))
		copy(copyList, list)
		paths = append(paths, copyList)

		list = list[:len(list)-1]
		return
	}

	if tn.leftChild != 0 {
		findPaths(tree[tn.leftChild], list)
	}
	if tn.rightChild != 0 {
		findPaths(tree[tn.rightChild], list)
	}

	list = list[:len(list)-1]
}

// http://ccc013.github.io/2016/08/18/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%9F%BA%E6%9C%AC%E6%A6%82%E5%BF%B5%E5%92%8C%E5%AE%9E%E7%8E%B0/
// https://projecteuler.net/problem=18
// 示例数据是二叉树，描述成链表
var tree Tree = []*TreeNode{
	&TreeNode{1, 2, 75},
	&TreeNode{3, 4, 95},
	&TreeNode{4, 5, 64},
	&TreeNode{6, 7, 17},
	&TreeNode{7, 8, 47},
	&TreeNode{8, 9, 82},
	&TreeNode{10, 11, 18},
	&TreeNode{11, 12, 35},
	&TreeNode{12, 13, 87},
	&TreeNode{13, 14, 10},
	&TreeNode{15, 16, 20},
	&TreeNode{16, 17, 4},
	&TreeNode{17, 18, 82},
	&TreeNode{18, 19, 47},
	&TreeNode{19, 20, 65},
	&TreeNode{21, 22, 19},
	&TreeNode{22, 23, 1},
	&TreeNode{23, 24, 23},
	&TreeNode{24, 25, 75},
	&TreeNode{25, 26, 3},
	&TreeNode{26, 27, 34},
	&TreeNode{28, 29, 88},
	&TreeNode{29, 30, 2},
	&TreeNode{30, 31, 77},
	&TreeNode{31, 32, 73},
	&TreeNode{32, 33, 7},
	&TreeNode{33, 34, 63},
	&TreeNode{34, 35, 67},
	&TreeNode{36, 37, 99},
	&TreeNode{37, 38, 65},
	&TreeNode{38, 39, 4},
	&TreeNode{39, 40, 28},
	&TreeNode{40, 41, 6},
	&TreeNode{41, 42, 16},
	&TreeNode{42, 43, 70},
	&TreeNode{43, 44, 92},
	&TreeNode{45, 46, 41},
	&TreeNode{46, 47, 41},
	&TreeNode{47, 48, 26},
	&TreeNode{48, 49, 56},
	&TreeNode{49, 50, 83},
	&TreeNode{50, 51, 40},
	&TreeNode{51, 52, 80},
	&TreeNode{52, 53, 70},
	&TreeNode{53, 54, 33},
	&TreeNode{55, 56, 41},
	&TreeNode{56, 57, 48},
	&TreeNode{57, 58, 72},
	&TreeNode{58, 59, 33},
	&TreeNode{59, 60, 47},
	&TreeNode{60, 61, 32},
	&TreeNode{61, 62, 37},
	&TreeNode{62, 63, 16},
	&TreeNode{63, 64, 94},
	&TreeNode{64, 65, 29},
	&TreeNode{66, 67, 53},
	&TreeNode{67, 68, 71},
	&TreeNode{68, 69, 44},
	&TreeNode{69, 70, 65},
	&TreeNode{70, 71, 25},
	&TreeNode{71, 72, 43},
	&TreeNode{72, 73, 91},
	&TreeNode{73, 74, 52},
	&TreeNode{74, 75, 97},
	&TreeNode{75, 76, 51},
	&TreeNode{76, 77, 14},
	&TreeNode{78, 79, 70},
	&TreeNode{79, 80, 11},
	&TreeNode{80, 81, 33},
	&TreeNode{81, 82, 28},
	&TreeNode{82, 83, 77},
	&TreeNode{83, 84, 73},
	&TreeNode{84, 85, 17},
	&TreeNode{85, 86, 78},
	&TreeNode{86, 87, 39},
	&TreeNode{87, 88, 68},
	&TreeNode{88, 89, 17},
	&TreeNode{89, 90, 57},
	&TreeNode{91, 92, 91},
	&TreeNode{92, 93, 71},
	&TreeNode{93, 94, 52},
	&TreeNode{94, 95, 38},
	&TreeNode{95, 96, 17},
	&TreeNode{96, 97, 14},
	&TreeNode{97, 98, 91},
	&TreeNode{98, 99, 43},
	&TreeNode{99, 100, 58},
	&TreeNode{100, 101, 50},
	&TreeNode{101, 102, 27},
	&TreeNode{102, 103, 29},
	&TreeNode{103, 104, 48},
	&TreeNode{105, 106, 63},
	&TreeNode{106, 107, 66},
	&TreeNode{107, 108, 4},
	&TreeNode{108, 109, 68},
	&TreeNode{109, 110, 89},
	&TreeNode{110, 111, 53},
	&TreeNode{111, 112, 67},
	&TreeNode{112, 113, 30},
	&TreeNode{113, 114, 73},
	&TreeNode{114, 115, 16},
	&TreeNode{115, 116, 69},
	&TreeNode{116, 117, 87},
	&TreeNode{117, 118, 40},
	&TreeNode{118, 119, 31},
	&TreeNode{0, 0, 4},
	&TreeNode{0, 0, 62},
	&TreeNode{0, 0, 98},
	&TreeNode{0, 0, 27},
	&TreeNode{0, 0, 23},
	&TreeNode{0, 0, 9},
	&TreeNode{0, 0, 70},
	&TreeNode{0, 0, 98},
	&TreeNode{0, 0, 73},
	&TreeNode{0, 0, 93},
	&TreeNode{0, 0, 38},
	&TreeNode{0, 0, 53},
	&TreeNode{0, 0, 60},
	&TreeNode{0, 0, 4},
	&TreeNode{0, 0, 23},
}

func main() {
	fmt.Println(len(tree))

	list := make([]*TreeNode, 0)
	findPaths(tree[0], list)

	fmt.Println(len(paths))

	maxSum := 0
	maxKey := 0
	for i, path := range paths {
		sum := 0
		for _, tn := range path {
			sum += tn.data
		}
		if sum > maxSum {
			maxKey = i
			maxSum = sum
		}
	}
	fmt.Println(maxKey, maxSum)

	// 打印各个节点
	for k, v := range paths[maxKey] {
		fmt.Println(k, *v)
	}
}
