package main

import (
	"connected-components/algorithms"
	"fmt"

	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
    grid := [][]int{
        {0, 1, 1, 0, 1},
        {1, 1, 0, 1, 0},
        {0, 0, 1, 0, 0},
        {1, 0, 1, 1, 0},
        {0, 0, 1, 1, 1},
    }

    fmt.Println(algorithms.DFSConnectedComponents(grid))

	go func(){
		w := new(app.Window)
		w.Option(app.Title("Visualisation of Largest Connected Component on a grid "))
		w.Option(app.Size(unit.Dp(1024), unit.Dp(1024)))
		for{
			w.Event()
		}
	}()
	app.Main()
}

