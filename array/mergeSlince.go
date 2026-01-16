package main

import (
	"fmt"
	"sort"
)

/*题目：合并区间

题目描述：
给定一个区间的集合，请合并所有重叠的区间。

示例：
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]*/

// 按照左区间排序【升序】
func sortSlice(slices [][]int) [][]int {
	sort.Slice(slices, func(i, j int) bool {
		if slices[i][0] == slices[j][0] {
			return slices[i][1] < slices[j][1]
		}
		return slices[i][0] < slices[j][0]
	})
	return slices
}
func mergeSlice(slices [][]int) [][]int {
	slices = sortSlice(slices)

	resultSlices := make([][]int, 0)
	start := slices[0][0]
	end := slices[0][1]
	for i := 1; i < len(slices); i++ {
		s, e := slices[i][0], slices[i][1]
		// 若有重合，则将end滞后
		if s <= end {
			if e > end {
				end = e
			}
		} else {
			resultSlices = append(resultSlices, []int{start, end})
			start, end = s, e
		}
	}
	// 记着最后的这个数组【最后一个数组处理不上】
	resultSlices = append(resultSlices, []int{start, end})
	return resultSlices
}

func main() {
	var sourceArrays = [][]int{{4, 8}, {11, 14}, {1, 3}, {2, 6}, {8, 10}, {15, 18}}
	resultArray := mergeSlice(sourceArrays)
	fmt.Println()
	fmt.Println(resultArray)
}
