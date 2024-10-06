package algorithms

import (
	"container/list"
)

// Finds the largest connected component using BFS

func BFSConnectedComponents(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	
	visited := make([][]bool,rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
    maxRegion := 0
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if matrix[row][col] == 1 {
                size := getRegionSizeBFS(matrix, row, col, visited)
				maxRegion = Max(size, maxRegion)
            }
        }
    }
	
    return maxRegion
}

// Calculates the size of the connected region using BFS
func getRegionSizeBFS(matrix [][]int, row int, column int, visited[][] bool) int {
    
    queue := list.New()
    queue.PushBack([2]int{row, column})  
    visited[row][column] = true
    size := 0

    for queue.Len() > 0 {
    
        current := queue.Remove(queue.Front()).([2]int)
        r, c := current[0], current[1]
        size++

        for _, dir := range Directions {
            newRow := r + dir[0]
            newCol := c + dir[1]

            if newRow >= 0 && newCol >= 0 && newRow < len(matrix) && newCol < len(matrix[0]) && matrix[newRow][newCol] == 1 && !visited[newRow][newCol] {
                visited[newRow][newCol] = true
                queue.PushBack([2]int{newRow, newCol})
            }
        }
    }

    return size
}
