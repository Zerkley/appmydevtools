package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Settings() *fyne.Container {
	mainView := container.NewHBox(canvas.NewText("text5", color.White), canvas.NewText("text6", color.White))

	return mainView
}
