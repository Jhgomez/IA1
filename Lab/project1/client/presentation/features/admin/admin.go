package admin

import (
	"fmt"

	"unimatch/data/repository/admin"
	"unimatch/presentation/navigation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	
)

func addSpacing(width float32, height float32) fyne.CanvasObject {
	spacing := canvas.NewRectangle(nil)
	spacing.SetMinSize(fyne.NewSize(width, height))

	return spacing
}

func AdminFacultySelectionScreen() fyne.CanvasObject {
	careers, error := adminrepo.GetAdminRepository().GetCareers()

	if error != nil {
		fmt.Println(error)
		return nil
	}
	
	// Map: Faculty -> List of Careers
	facultyMap := make(map[string][]adminrepo.EditableCareer)

	for _, c := range careers {
		facultyMap[c.Faculty] = append(facultyMap[c.Faculty], c)
	}

	var facultyButtons []fyne.CanvasObject

	for faculty, facultyCareers := range facultyMap {
		// Button to represent each faculty
		btn := widget.NewButton(faculty, func(fac string, careers []adminrepo.EditableCareer) func() {
			return func() {

				navigation.NavigateWithNewWindow(
					"Admin",  //windowTitle
					adminCareerSelectionScreen(careers), //content
					true, // shouldHide
					fyne.NewSize(400, 500), // windows size
					nil, // onClose
				)
			}
		}(faculty, facultyCareers))

		facultyButtons = append(facultyButtons, btn)
	}

	title := canvas.NewText("Facultades", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	content := container.NewVBox(
		title,
		addSpacing(0, 16),
		container.NewVBox(facultyButtons...),
	)

	return container.NewScroll(content)
}


func adminCareerSelectionScreen(careers []adminrepo.EditableCareer) fyne.CanvasObject {
	var careerButtons []fyne.CanvasObject

	for _, career := range careers {
		c := career // capture for closure
		btn := widget.NewButton(c.Career, func() {
			// Handle career selection here
			println("Selected Career:", c.Career)
			// TODO: navigate to AdminCareerEditorScreen(c)
		})
		careerButtons = append(careerButtons, btn)
	}

	// Title
	title := canvas.NewText("Carreras en la facultad", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	// "Agregar carrera" button
	addCareerBtn := widget.NewButton("Agregar carrera", func() {
		println("Agregar nueva carrera")
		// TODO: Navigate to career creation form
	})

	// Scrollable content (title + list of careers)
	scrollArea := container.NewVScroll(
		container.NewVBox(
			title,
			addSpacing(0, 16),
			container.NewVBox(careerButtons...),
		),
	)
	scrollArea.SetMinSize(fyne.NewSize(400, 400))

	// Layout with the button at the bottom
	content := container.NewBorder(nil, addCareerBtn, nil, nil, scrollArea)

	return content
}
