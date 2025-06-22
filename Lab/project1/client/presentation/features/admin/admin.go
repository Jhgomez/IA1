package admin

import (
	"fmt"
	"strings"

	"unimatch/data/repository/admin"
	"unimatch/presentation/navigation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	facultyMap := make(map[string][]*adminrepo.EditableCareer)

	for _, c := range careers {
		facultyMap[c.Faculty] = append(facultyMap[c.Faculty], c)
	}

	var facultyButtons []fyne.CanvasObject

	for faculty, facultyCareers := range facultyMap {
		// Button to represent each faculty
		btn := widget.NewButton(faculty, func(fac string, careers []*adminrepo.EditableCareer) func() {
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


func adminCareerSelectionScreen(careers []*adminrepo.EditableCareer) fyne.CanvasObject {
	var careerButtons []fyne.CanvasObject

	for _, career := range careers {
		c := career // capture for closure
		btn := widget.NewButton(c.Career, func() {
			navigation.NavigateWithNewWindow(
				c.Career,  //windowTitle
				careerEditScreen(c), //content
				true, // shouldHide
				fyne.NewSize(700, 500), // windows size
				func () {
					navigation.PopBackstack()
				}, // onClose
			)
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

func careerEditScreen(career *adminrepo.EditableCareer) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Editar carrera: "+career.Career, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	repo := adminrepo.GetAdminRepository()

	aptitudeBox, aptitudeEntries := buildEditableList(career.Aptitude, func (index int) {
		repo.DeleteFact(career.CareerId, career.Aptitude[index], "", "")

		career.Aptitude[index] = ""
	})
	skillBox, skillEntries := buildEditableList(career.Skill, func (index int) {
		repo.DeleteFact(career.CareerId, "", career.Skill[index], "")
		career.Skill[index] = ""	
	})
	interestBox, interestEntries := buildEditableList(career.Interest, func (index int) {
		repo.DeleteFact(career.CareerId, "", "", career.Interest[index])
		career.Interest[index] = ""	
	})

	aptitudesSection := container.NewVBox(
		widget.NewLabelWithStyle("Aptitudes", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		aptitudeBox,
	)
	skillsSection := container.NewVBox(
		widget.NewLabelWithStyle("Habilidades", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		skillBox,
	)
	interestsSection := container.NewVBox(
		widget.NewLabelWithStyle("Intereses", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		interestBox,
	)

	deleteBtn := widget.NewButton("Borrar carrera", func() {
		// handle delete
	})


	// Save and delete buttons
	saveBtn := widget.NewButton("Guardar", func() {
		var newAptitudes []string
		for _, entry := range aptitudeEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newAptitudes = append(newAptitudes, text)
			}
		}

		var newSkills []string
		for _, entry := range skillEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newSkills = append(newSkills, text)
			}
		}

		var newInterests []string
		for _, entry := range interestEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newInterests = append(newInterests, text)
			}
		}

		// fmt.Println(career.Aptitude)
		// fmt.Println(newAptitudes)
		repo.UpdateCareer(career.CareerId, newAptitudes, newSkills, newInterests, career.Aptitude, career.Skill, career.Interest)

		career.Aptitude = newAptitudes
		career.Skill = newSkills
		career.Interest = newInterests
		// fmt.Println(career.Aptitude)
	})

	topRow := container.NewHBox(
		layout.NewSpacer(),
		saveBtn,
	)

	bottomRow := container.NewHBox(
		layout.NewSpacer(),
		deleteBtn,
	)

	content := container.NewVBox(
		title,
		topRow,
		container.NewHBox(
			aptitudesSection,
			layout.NewSpacer(),
			skillsSection,
			layout.NewSpacer(),
			interestsSection,
		),
		layout.NewSpacer(),
		bottomRow,
	)

	return content
}

// option 1 for aptitudes, 2 for skills, 3 for interests
func buildEditableList(items []string, onClick func(index int)) (*fyne.Container, []*widget.Entry) {
	var entries []*widget.Entry
	list := container.NewVBox()

	var refresh func()

	refresh = func() {
		list.Objects = nil
		for i, entry := range entries {
			// delete button
			trash := widget.NewButtonWithIcon("", theme.DeleteIcon(), func(index int) func() {
				return func() {
					entries = append(entries[:index], entries[index+1:]...)
					refresh()
					onClick(index)
				}
			}(i))

			// content := container.NewHBox(entry, trash)

			// Make the entry expand as much as possible
			entryContainer := container.NewBorder(nil, nil, nil, trash, entry)
			list.Add(entryContainer)
		}
		list.Refresh()
	}

	for _, item := range items {
		if item == "" {
			continue
		}

		entry := widget.NewEntry()
		entry.SetText(item)
		entry.Wrapping = fyne.TextWrapOff // Optional: prevent line breaks
		entries = append(entries, entry)
	}

	refresh()

	scroll := container.NewVScroll(list)
	scroll.SetMinSize(fyne.NewSize(240, 300)) // Set width to accommodate widest content

	return container.NewVBox(scroll), entries
}
