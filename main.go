package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	customLayouts "github.com/zerkley/appmydevtools/layouts"
	"github.com/zerkley/appmydevtools/views"
)

func RenderElement(update *fyne.Container, component *fyne.Container) {
	update.Objects = []fyne.CanvasObject{component}
	update.Refresh()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("AppmyDevtools")
	customLayout := &customLayouts.MainWindowLayout{}
	updatingComponent := views.Home()
	sideBar := views.Sidebar(updatingComponent)
	statusBar := container.NewBorder(

		canvas.NewLine(color.White), nil, nil, nil,
		canvas.NewText("Status bar", color.White),
	)

	statusBar.Resize(fyne.NewSize(700, 30))

	mainContent := container.New(layout.NewHBoxLayout(), sideBar, updatingComponent)

	fullWindow := container.New(customLayout, mainContent, statusBar)
	myWindow.Resize(fyne.NewSize(700, 600))
	myWindow.SetContent(fullWindow)
	myWindow.ShowAndRun()
}
