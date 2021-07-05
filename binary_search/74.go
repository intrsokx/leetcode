package main

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104
//
// Related Topics 数组 二分查找 矩阵
// 👍 451 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

//思路：将二维数组转换成1维数组去理解，再二分
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])

	start, end := 0, row*col-1

	//start+1 == end
	for start+1 < end {
		mid := (end-start)/2 + start
		val := matrix[mid/col][mid%col]
		if val == target {
			return true
		} else if val < target {
			start = mid
		} else {
			end = mid
		}
	}

	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
