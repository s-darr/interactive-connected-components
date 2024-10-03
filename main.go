package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		
	}

	// Create a new Fyne app
	myApp := app.New()
	myWindow := myApp.NewWindow("Interactive Grid Visualization")

	
	windowSize := 600
	myWindow.Resize(fyne.NewSize(float32(windowSize), float32(windowSize)))

	// Create a label to show the result of BFSConnectedComponents
	resultLabel := widget.NewLabel(fmt.Sprintf("Largest Connected Component: %d", 0))

	rows := len(grid)
	cellSize := windowSize / rows

	gridContainer := container.NewGridWithRows(rows)

	for rowIndex, row := range grid {
		for colIndex := range row {
			cellValue := &grid[rowIndex][colIndex]
			rect := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 255}) // // All blue to begin with
			rect.SetMinSize(fyne.NewSize(float32(cellSize), float32(cellSize)))

			clickableRectangle := &ClickableRectangle{
				rect: rect,
				cellValue: cellValue,
				resultLabel: resultLabel,
				grid: grid,
				rowIndex: rowIndex,
				colIndex: colIndex,
			}
			clickableRectangle.ExtendBaseWidget(clickableRectangle)
			gridContainer.Add(clickableRectangle)
		}
	}


	content := container.NewVBox(resultLabel, gridContainer)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

type ClickableRectangle struct {
	widget.BaseWidget
	rect        *canvas.Rectangle
	cellValue   *int
	resultLabel *widget.Label
	grid        [][]int
	rowIndex int
	colIndex int
}

func (c *ClickableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return &ClickableRectangleRenderer{
		clickableRectangle: c,
		objects:            []fyne.CanvasObject{c.rect},
	}
}

func (c *ClickableRectangle) Tapped(_ *fyne.PointEvent) {
	fmt.Println(*c.cellValue)
    if *c.cellValue == 1 {
        *c.cellValue = 0
        c.grid[c.rowIndex][c.colIndex] = 0
        c.rect.FillColor = color.NRGBA{R: 0, G: 0, B: 255, A: 255} // Blue for water
    } else {
        *c.cellValue = 1
		
        c.grid[c.rowIndex][c.colIndex] = 1
        c.rect.FillColor = color.NRGBA{R: 0, G: 128, B: 0, A: 255} // Green for land
    }
	

	c.Refresh()
	

 
}


type ClickableRectangleRenderer struct {
	clickableRectangle *ClickableRectangle
	objects            []fyne.CanvasObject
}

func (r *ClickableRectangleRenderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *ClickableRectangleRenderer) MinSize() fyne.Size {
	return r.objects[0].MinSize()
}

func (r *ClickableRectangleRenderer) Refresh() {
	canvas.Refresh(r.clickableRectangle)
}

func (r *ClickableRectangleRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *ClickableRectangleRenderer) Destroy() {
	// Optional
}

