package algorithms

//  Finds the largest connected component using DFS
func DFSConnectedComponents(matrix [][]int) int {
    maxRegion := 0
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[row]); col++ {
            if matrix[row][col] == 1 {
                size := getRegionSizeDFS(matrix, row, col)
                if size > maxRegion {
                    maxRegion = size
                }
            }
        }
    }
    return maxRegion
}

// Calculates the size of the connected region using DFS
func getRegionSizeDFS(matrix [][]int, row int, column int) int {
    if row < 0 || column < 0 || row >= len(matrix) || column >= len(matrix[0]) || matrix[row][column] == 0 {
        return 0
    }

    matrix[row][column] = 0
    size := 1

    for _, dir := range Directions {
        r := row + dir[0]
        c := column + dir[1]
        size += getRegionSizeDFS(matrix, r, c)
    }
    return size
}
