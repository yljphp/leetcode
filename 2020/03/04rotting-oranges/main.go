package main

import (
	"fmt"
)

/**


994. 腐烂的橘子
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
@see https://leetcode-cn.com/problems/rotting-oranges/
*/
func main() {

	//grid := [][]int{
	//	[]int{2, 1, 1},
	//	[]int{0, 1, 1},
	//	[]int{1, 0, 1},
	//}

	grid := [][]int{
		[]int{0},
	}

	num := orangesRotting(grid)

	fmt.Println(num)
}

func handleGridValue(tmp *int) bool {

	if *tmp == 1 {
		*tmp = 2

		return true
	}

	return false
}

type BadCoordinate struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {

	//x y轴最大值
	maxX := len(grid) - 1
	maxY := len(grid[0]) - 1

	minute := 0

	badSlice := make([]BadCoordinate, 0)

	for x, tmpItem := range grid {

		for y, item := range tmpItem {

			if item == 2 {

				//fmt.Println(x, y)

				badSlice = append(badSlice, BadCoordinate{x, y})
			}
		}
	}

	if len(badSlice) <= 0 {
		if checkExistNormal(grid) {
			return -1
		}
		return 0
	}

	for {

		tmpBadSlice := badSlice

		for _, item := range badSlice {

			if item.x >= 1 {

				if handleGridValue(&grid[item.x-1][item.y]) {
					tmpBadSlice = append(tmpBadSlice, BadCoordinate{item.x - 1, item.y})
				}

			}

			if item.x < maxX {
				if handleGridValue(&grid[item.x+1][item.y]) {
					tmpBadSlice = append(tmpBadSlice, BadCoordinate{item.x + 1, item.y})
				}
			}

			if item.y >= 1 {

				if handleGridValue(&grid[item.x][item.y-1]) {
					tmpBadSlice = append(tmpBadSlice, BadCoordinate{item.x, item.y - 1})
				}
			}

			if item.y < maxY {
				if handleGridValue(&grid[item.x][item.y+1]) {
					tmpBadSlice = append(tmpBadSlice, BadCoordinate{item.x, item.y + 1})
				}
			}
		}

		if len(tmpBadSlice) == len(badSlice) {

			//验证是否存在正常的
			if checkExistNormal(grid) {
				minute = -1
			}

			fmt.Println(tmpBadSlice)
			break
		}

		//计数加一
		minute++

		badSlice = tmpBadSlice
	}

	//判断是否存在永远不会腐烂的情况

	//fmt.Println(badSlice)
	return minute
}

//检测处理后的数据中是否存在正常的
func checkExistNormal(grid [][]int) bool {

	for _, tmpItem := range grid {

		for _, item := range tmpItem {

			if item == 1 {

				return true
			}
		}
	}

	return false
}
