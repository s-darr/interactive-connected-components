package main

import "fmt"
func main(){
	maxRegion := 0
	var grid = [5][5]int{ // [rows][cols]
		{0,1,1,0,1},
	    {1,1,0,1,0},
		{0,0,1,0,0},
		{1,0,1,1,0},
		{0,0,1,1,1},
	}

	for row := 0; row < len(grid); row++ {
		for col:=0; col< len(grid); col++{
			if  grid[row][col] == 1{
				size := getRegionSize(&grid, row, col)
				if size > maxRegion{
					maxRegion = size
				}
			}
		}
	}
	fmt.Println(maxRegion)
}
func getRegionSize(matrix *[5][5]int, row int, column int) int {
	if row < 0 || column < 0 || row >= len(matrix) || column >= len(matrix){
		return 0
	}
	if matrix[row][column] == 0 {
		return 0
	}
	matrix[row][column] = 0
	size := 1
	directions := [][]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
	}
	
	
	for _, dir := range directions {
		r := row + dir[0]
		c := column + dir[1]
		size += getRegionSize(matrix, r, c)
	}
	return size;
}


