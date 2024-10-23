package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/zerkley/appmydevtools/databases"
	"github.com/zerkley/appmydevtools/helpers"
	customToolsLayouts "github.com/zerkley/appmydevtools/layouts"
	"github.com/zerkley/appmydevtools/views"
)

func main() {
	// Inits
	myApp := app.New()
	myWindow := myApp.NewWindow("AppmyDevtools")
	settings, err := helpers.SettingsLoader()

	if err != nil {
		fmt.Printf("Error loading settings: %v", err)
	}

	go databases.MongoAccess(settings.Database.Dev)

	// Components
	customLayout := &customToolsLayouts.MainWindowLayout{}
	updatingComponent := views.Home()
	sideBar := views.Sidebar(updatingComponent)
	statusBar := views.StatusBar()

	// Main Rendering
	mainContent := container.New(layout.NewHBoxLayout(), sideBar, updatingComponent)
	fullWindow := container.New(customLayout, mainContent, statusBar)

	// Window settings
	myWindow.Resize(fyne.NewSize(700, 600))
	myWindow.SetContent(fullWindow)
	myWindow.ShowAndRun()
}
