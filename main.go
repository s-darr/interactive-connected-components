package main

import (
	"connected-components/algorithms"
	"fmt"
)

func main() {
    grid := [][]int{
        {0, 1, 1, 0, 1},
        {1, 1, 0, 1, 0},
        {0, 0, 1, 0, 0},
        {1, 0, 1, 1, 0},
        {0, 0, 1, 1, 1},
    }

    fmt.Println(algorithms.BFSConnectedComponents(grid))

}

