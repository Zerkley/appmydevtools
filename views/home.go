package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Home() *fyne.Container {
	mainView := container.NewHBox(canvas.NewText("text1", color.White), canvas.NewText("text2", color.White))

	return mainView
}
