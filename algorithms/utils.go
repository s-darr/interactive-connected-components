package algorithms


var Directions = [][]int{
    {-1, 0},  // Up
    {1, 0},   // Down
    {0, -1},  // Left
    {0, 1},   // Right
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}