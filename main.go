package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	output := findDiagonalOrder(input)
	fmt.Println(output)
}

func findDiagonalOrder(mat [][]int) []int {
	var recursiveFn func(*[][]int, int, int, string) []int
	recursiveFn = func(mat *[][]int, x, y int, flow string) []int {
		if len(*mat) == 0 {
			return nil
		}
		if len(*mat)-1 == x && len((*mat)[0])-1 == y {
			return []int{(*mat)[x][y]}
		}

		nextX, nextY := x, y
		if flow == "topRight" {
			nextX, nextY = x-1, y+1
			if nextX < 0 || nextY >= len((*mat)[0]) {
				nextX, nextY = x, y+1
				if nextY >= len((*mat)[0]) {
					nextX, nextY = x+1, y
				}
				flow = "downLeft"
			}
		} else if flow == "downLeft" {
			nextX, nextY = x+1, y-1
			if nextX >= len(*mat) || nextY < 0 {
				nextX, nextY = x+1, y
				if nextX >= len(*mat) {
					nextX, nextY = x, y+1
				}
				flow = "topRight"
			}
		}
		result := append([]int{(*mat)[x][y]}, recursiveFn(mat, nextX, nextY, flow)...)
		return result
	}

	return recursiveFn(&mat, 0, 0, "topRight")
}
