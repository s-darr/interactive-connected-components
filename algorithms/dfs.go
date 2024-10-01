package algorithms

//  Finds the largest connected component using DFS.
func DFSConnectedComponents(matrix [][]int) int {
    maxRegion := 0
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[row]); col++ {
            if matrix[row][col] == 1 {
                size := getRegionSize(matrix, row, col)
                if size > maxRegion {
                    maxRegion = size
                }
            }
        }
    }
    return maxRegion
}

// Calculates the size of the connected region using DFS.
func getRegionSize(matrix [][]int, row int, column int) int {
    if row < 0 || column < 0 || row >= len(matrix) || column >= len(matrix[0]) || matrix[row][column] == 0 {
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

    return size
}
