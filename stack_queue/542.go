package main

//给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
//
//
// 两个相邻元素间的距离为 1 。
//
//
//
// 示例 1：
//
//
//输入：
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
//
//输出：
//[[0,0,0],
// [0,1,0],
// [0,0,0]]
//
//
// 示例 2：
//
//
//输入：
//[[0,0,0],
// [0,1,0],
// [1,1,1]]
//
//输出：
//[[0,0,0],
// [0,1,0],
// [1,2,1]]
//
//
//
//
// 提示：
//
//
// 给定矩阵的元素个数不超过 10000。
// 给定矩阵中至少有一个元素是 0。
// 矩阵中的元素只在四个方向上相邻: 上、下、左、右。
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵
// 👍 444 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//求每个1到0的曼哈顿距离
//即曼哈顿距离，曼哈顿距离就是只能沿着横着、竖着路径到达另一个节点的距离。
//思路：BFS 广度优先遍历，从0开始进入队列，弹出之后计算上下左右的距离，将上下左右重新入队进行下一层操作
func updateMatrix(mat [][]int) [][]int {
	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[i]

			var x, y int
			for _, v := range directions {
				x = point[0] + v[0]
				y = point[1] + v[1]

				if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == -1 {
					mat[x][y] = mat[point[0]][point[1]] + 1
					queue = append(queue, []int{x, y})
				}
			}

		}
		queue = queue[size:]
	}

	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
