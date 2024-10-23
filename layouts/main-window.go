package customToolsLayouts

import (
	"fyne.io/fyne/v2"
)

type MainWindowLayout struct {
}

func (m *MainWindowLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	// Assuming the width should fit 100% of the container's width
	if len(objects) > 0 {
		firstMinSize := objects[0].MinSize()
		return fyne.NewSize(firstMinSize.Width, firstMinSize.Height)
	}
	return fyne.NewSize(0, 0)
}

func (m *MainWindowLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) == 0 {
		return
	}

	// First item: 95% of the height, 100% of the width
	firstItemHeight := containerSize.Height * 0.95
	secondItemHeight := containerSize.Height * 0.05

	if len(objects) > 0 {
		firstItem := objects[0]
		firstItem.Resize(fyne.NewSize(containerSize.Width, firstItemHeight))
		firstItem.Move(fyne.NewPos(0, 0))
	}

	// Second item: 10% of the height, 100% of the width
	if len(objects) > 1 {
		secondItem := objects[1]
		secondItem.Resize(fyne.NewSize(containerSize.Width, secondItemHeight))
		secondItem.Move(fyne.NewPos(0, firstItemHeight)) // Positioned right after the first item
	}
}
