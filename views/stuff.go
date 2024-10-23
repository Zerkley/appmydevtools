package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Stuff() *fyne.Container {
	mainView := container.NewHBox(canvas.NewText("text3", color.White), canvas.NewText("text4", color.White))

	return mainView
}
