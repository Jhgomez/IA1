package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"

	"unimatch/presentation/navigation"
	"unimatch/presentation/features/home"
	"unimatch/data/repository/knowledge"
)


func main() {
	_, err := knowledgerepo.GetKnowledgeRepo().LoadKnowledgeBase()

	if err != nil {
		fmt.Printf("error in knowledge repo %v", err)
	}
	
	navigation.NavigateWithNewWindow(
		"Unimatch",  //windowTitle
		home.HomeScreen(),  //content
		false, // shouldHide
		fyne.NewSize(400, 500), // windows size
		nil, // onClose
	)

	// // messages send button
	// sendButton := widget.NewButton("Enviar", func() {
	// 	navigation.NavigateWithNewWindow(
	// 		"new", //windowTitle
	// 		widget.NewLabel("New Window"), //content
	// 		true,	// shouldHide
	// 		func() {	// onClose
	// 			// navigation.PopBackstack()
	// 		},
	// 	)
	// })

	// // sendButton := widget.NewButton("Enviar", func() {
	// // 	navigation.NavigateWithCurrentWindow(
	// // 		"new", //windowTitle
	// //		fyne.NewSize(400, 300), // size
	// // 		widget.NewLabel("New Window"), //content
	// // 		// true,	// shouldHide
	// // 		// func() {	// onClose
	// // 		// 	navigation.PopBackstack()
	// // 		// },
	// // 	)
	// // })

	// // Main content layout
	// home := container.NewVBox(sendButton,layout.NewSpacer())
		
	// navigation.NavigateWithNewWindow(
	// 	"home",  //windowTitle
	// 	home,  //content
	// 	false, // shouldHide
	//	fyne.NewSize(400, 300), // size
	// 	nil, // onClose
	// )
}