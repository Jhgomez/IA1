package navigation

import (
	// "fmt"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var unimatchApp fyne.App = app.New()
var backStack = []fyne.Window{}

// func SetUpHomeScreen(windowTitle string, content fyne.CanvasObject) {
// 	if unimatchApp == nil {
// 		unimatchApp = app.New()

// 		currentWindow = unimatchApp.NewWindow(windowTitle)

// 		currentWindow.SetContent(content)

// 		currentWindow.ShowAndRun()
// 	}
// }

func NavigateWithNewWindow(
	windowTitle string,
	content fyne.CanvasObject,
	shouldHide bool,
	onClose func(),
) {
	if unimatchApp == nil {
		panic(errors.New("Fyne App object can not be null"))
	}

	backStack = append(backStack, unimatchApp.NewWindow(windowTitle))

	backstackLen := len(backStack)

	currentWindow := backStack[backstackLen - 1]

	currentWindow.SetContent(content)

	currentWindow.SetCloseIntercept(func() {
		if shouldHide {
			if onClose != nil {
				// if call back defined you're responsible for popping out nav stack
				onClose()
			} else {
				PopBackstack()
			}
		} else {
			if onClose != nil {
				// if call back defined you're responsible for popping out nav stack
				onClose()
			} else {
				if backstackLen == 1 {
					backStack[0].Close()
				}
			}
		}
	})

	if len(backStack) == 1 {
		// fmt.Println("First window")
		// fmt.Println(&backStack)
		if shouldHide {
			panic(errors.New("can not hide first screen on app start"))
		}
		
		currentWindow.ShowAndRun()

		return
	}

	// currentWindow := unimatchApp.NewWindow(windowTitle)
	// currentWindow.SetContent(content)

	prevWindow := backStack[len(backStack) - 2]

	if shouldHide {
		prevWindow.Hide()
	}

	// fmt.Println("New window")
	// fmt.Println(&backStack)

	currentWindow.Show()
}

func PopBackstack() {
	lastIndex := len(backStack) - 1
	
	currentWindow := backStack[lastIndex]

	backStack = backStack[:lastIndex]

	prevWindow := backStack[len(backStack) - 1]

	// fmt.Println("Pop window")
	// fmt.Println(&backStack)

	currentWindow.Close()
	prevWindow.Show()
}

func NavigateWithCurrentWindow(windowTitle string, content fyne.CanvasObject) {
	if unimatchApp == nil {
		panic(errors.New("Fyne App object can not be null"))
	}

	currentWindow := backStack[len(backStack) - 1]

	currentWindow.SetContent(content)
	
	currentWindow.SetTitle(windowTitle)
	
	// fmt.Println("Replaced window title")
	// fmt.Println(backStack[len(backStack) - 1].Title())
	// fmt.Println(backStack)
}
