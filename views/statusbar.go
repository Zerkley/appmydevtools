package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	customToolsLayouts "github.com/zerkley/appmydevtools/layouts"
)

var dbCheck *canvas.Circle

func StatusBar() *fyne.Container {
	customLayout := &customToolsLayouts.StatusBarLayout{}
	dbCheck = canvas.NewCircle(color.RGBA{255, 0, 0, 200})

	StatusItems := container.New(customLayout, dbCheck, canvas.NewText("MongoDB", color.White))
	mainView := container.NewBorder(
		canvas.NewLine(color.White), nil, nil, nil, StatusItems,
	)

	return mainView
}

func SetColor(status bool) {
	if dbCheck == nil {
		return // Ensure the circle is initialized
	}
	if status {
		dbCheck.StrokeWidth = 1
		dbCheck.FillColor = color.RGBA{166, 227, 161, 255} // Green color for "true" status
	} else {
		dbCheck.FillColor = color.RGBA{255, 0, 0, 200} // Red color for "false" status
	}
	dbCheck.Refresh() // Ensure the circle gets redrawn with the new color
}
