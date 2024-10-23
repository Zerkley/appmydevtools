package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RenderElement(update *fyne.Container, component *fyne.Container) {
	update.Objects = []fyne.CanvasObject{component}
	update.Refresh()
}

func Sidebar(updatingComponent *fyne.Container) *fyne.Container {
	button1 := widget.NewButton("Home", func() {
		RenderElement(updatingComponent, Home())
	})
	button2 := widget.NewButton("Stuff", func() {
		RenderElement(updatingComponent, Stuff())
	})
	button3 := widget.NewButton("Settings", func() {
		RenderElement(updatingComponent, Settings())
	})
	mainView := container.NewBorder(nil, nil, nil, canvas.NewLine(color.White), container.NewVBox(button1, button2, button3))

	return mainView
}
