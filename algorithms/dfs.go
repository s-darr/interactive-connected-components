package algorithms

//  Finds the largest connected component using DFS
func DFSConnectedComponents(matrix [][]int) int {
    visited := make([][]bool, len(matrix))
    for i := range visited {
        visited[i] = make([]bool, len(matrix[0]))
    }
    maxRegion := 0
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[row]); col++ {
            if matrix[row][col] == 1 && !visited[row][col] {
                size := getRegionSizeDFS(matrix, visited, row, col)
                if size > maxRegion {
                    maxRegion = size
                }
            }
        }
    }
    return maxRegion
}

func getRegionSizeDFS(matrix [][]int, visited [][]bool, row int, column int) int {
    if row < 0 || column < 0 || row >= len(matrix) || column >= len(matrix[0]) || matrix[row][column] == 0 || visited[row][column] {
        return 0
    }

    visited[row][column] = true
    size := 1

    for _, dir := range Directions {
        r := row + dir[0]
        c := column + dir[1]
        size += getRegionSizeDFS(matrix, visited, r, c)
    }
    return size
}