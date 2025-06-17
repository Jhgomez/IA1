package navigation

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var unimatchApp fyne.App
var currentWindow fyne.Window

func SetUpHomeScreen(windowTitle string, content fyne.CanvasObject) {
	if unimatchApp == nil {
		unimatchApp = app.New()

		currentWindow = unimatchApp.NewWindow(windowTitle)

		currentWindow.SetContent(content)

		currentWindow.ShowAndRun()
	}
}

func NavigateWithNewWindow(windowTitle string, content fyne.CanvasObject) {
	if unimatchApp == nil {
		panic(errors.New("Fyne App object can not be null"))
	}

	if currentWindow == nil {
		panic(errors.New("Master Window can not be null"))
	}

	prevWindow := currentWindow

	currentWindow := unimatchApp.NewWindow(windowTitle)
	currentWindow.SetContent(content)

	prevWindow.Close()
	currentWindow.Show()

}

func NavigateWithCurrentWindow(content fyne.CanvasObject) {
	currentWindow.SetContent(content)
}
