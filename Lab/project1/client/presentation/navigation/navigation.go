package navigation

import (
	// "fmt"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
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
	shouldHidePrevWindow bool,
	size fyne.Size,
	onClose func(),
) {
	if unimatchApp == nil {
		panic(errors.New("Fyne App object can not be null"))
	}

	backStack = append(backStack, unimatchApp.NewWindow(windowTitle))

	backstackLen := len(backStack)

	currentWindow := backStack[backstackLen - 1]

	currentWindow.SetContent(content)

	currentWindow.Resize(size)

	currentWindow.SetCloseIntercept(func() {
		if shouldHidePrevWindow {
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
		if shouldHidePrevWindow {
			panic(errors.New("can not hide first screen on app start"))
		}

		currentWindow.CenterOnScreen()
		
		currentWindow.ShowAndRun()

		return
	}

	// currentWindow := unimatchApp.NewWindow(windowTitle)
	// currentWindow.SetContent(content)

	prevWindow := backStack[len(backStack) - 2]

	if shouldHidePrevWindow {
		prevWindow.Hide()
	}

	// fmt.Println("New window")
	// fmt.Println(&backStack)
	currentWindow.CenterOnScreen()

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

func NavigateWithCurrentWindow(windowTitle string, size fyne.Size, content fyne.CanvasObject) {
	if unimatchApp == nil {
		panic(errors.New("Fyne App object can not be null"))
	}

	currentWindow := backStack[len(backStack) - 1]

	currentWindow.SetContent(content)

	currentWindow.Resize(size)
	
	currentWindow.SetTitle(windowTitle)

	currentWindow.CenterOnScreen()
	
	// fmt.Println("Replaced window title")
	// fmt.Println(backStack[len(backStack) - 1].Title())
	// fmt.Println(backStack)
}

func ShowDialog(title, msg string) {
	dialog.ShowInformation(title, msg, backStack[len(backStack) - 1])
}

func ShowErrorDialog(msg string) {
	dialog.ShowError(errors.New(msg), backStack[len(backStack) - 1])
}