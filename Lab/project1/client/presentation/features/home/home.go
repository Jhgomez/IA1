package home

import (
	// "image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"unimatch/presentation/navigation"
	"unimatch/presentation/features/student"
)

func HomeScreen() fyne.CanvasObject {
	// Title
	title := canvas.NewText("Unimatch", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 28 // default is 14

	// // Title
	// title := widget.NewLabelWithStyle("Career Path Assistant", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Description
	description := widget.NewLabelWithStyle(
	"Bienvenido, Unimatch es un sistema inteligente que te ayudara a encontrar carreras\nafines a tus apitudes, habilidades e intereses",
		fyne.TextAlignCenter,
		fyne.TextStyle{},
	)

	// Buttons
	studentBtn := widget.NewButtonWithIcon("Student", theme.AccountIcon(), func() {
		navigation.NavigateWithNewWindow(
			"Formulary",  //windowTitle
			student.StudenFirstFormulary(),  //content
			true, // shouldHide
			fyne.NewSize(700, 500), // windows size
			nil, // onClose
		)
	})
	adminBtn := widget.NewButtonWithIcon("Administrator", theme.SettingsIcon(), func() {
		// Handle admin button click
	})

	buttons := container.NewHBox(layout.NewSpacer(), studentBtn, layout.NewSpacer(), adminBtn, layout.NewSpacer())

	// Layout: VBox (Title > Description > Buttons)
	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		description,
		layout.NewSpacer(),
		buttons,
		layout.NewSpacer(),
	)

	return content
}