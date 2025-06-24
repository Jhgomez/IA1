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
	"unimatch/data/repository/knowledge"
)

var updateRequired = false

func addSpacing(width float32, height float32) fyne.CanvasObject {
	spacing := canvas.NewRectangle(nil)
	spacing.SetMinSize(fyne.NewSize(width, height))

	return spacing
}


func CheckIfUpdateNeeded() {
	if updateRequired {
		_, err := knowledgerepo.GetKnowledgeRepo().LoadKnowledgeBase()

		if err != nil {
			fmt.Printf("error in knowledge repo %v", err)
		}

		fmt.Println("Knowledge updated")
	}
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
					adminCareerSelectionScreen(&careers), //content
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


func adminCareerSelectionScreen(careers *[]*adminrepo.EditableCareer) fyne.CanvasObject {
	var careerButtons []fyne.CanvasObject
	list := container.NewVBox()

	// "Agregar carrera" button
	addCareerBtn := widget.NewButton("Agregar carrera", func() {
		println("Agregar nueva carrera")
		// TODO: Navigate to career creation form
	})

	for index, career := range *careers {
		c := career // capture for closure
		btn := widget.NewButton(c.Career, func() {
			navigation.NavigateWithNewWindow(
				c.Career,  //windowTitle
				careerEditScreen(c, func () {
					adminrepo.GetAdminRepository().DeleteCareer(c.CareerId)
					updateRequired = true
					
					list.Objects = nil
					
					// fmt.Println(len(careers))
					// fmt.Println(careers)
					*careers = append((*careers)[:index], (*careers)[index+1:]...)

					// fmt.Println(len(careers))
					// fmt.Println(careers)
					navigation.PopBackstack()

					careerButtons = append(careerButtons[:index], careerButtons[index+1:]...)

					for _, button := range careerButtons {
						list.Add(button)
					}

					list.Refresh()

				}), //content
				true, // shouldHide
				fyne.NewSize(900, 550), // windows size
				func () {
					navigation.PopBackstack()
				}, // onClose
			)
		})
		careerButtons = append(careerButtons, btn)
		list.Add(btn)
	}

	// Title
	title := canvas.NewText("Carreras en la facultad", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	// Scrollable content (title + list of careers)
	scrollArea := container.NewVScroll(
		container.NewVBox(
			title,
			addSpacing(0, 16),
			list,
		),
	)
	scrollArea.SetMinSize(fyne.NewSize(400, 400))

	// Layout with the button at the bottom
	content := container.NewBorder(nil, addCareerBtn, nil, nil, scrollArea)

	return content
}

func careerEditScreen(career *adminrepo.EditableCareer, onDeleteCareer func()) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Editar carrera: "+career.Career, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	repo := adminrepo.GetAdminRepository()

	aptitudeEntries := []*widget.Entry{}

	aptitudeBox := buildEditableList(career.Aptitude, &aptitudeEntries, func (index int) {
		if index > len(career.Aptitude)-1 {
			return
		}

		// fmt.Println(aptitudeEntries)
		repo.DeleteFact(career.CareerId, career.Aptitude[index], "", "")
		updateRequired = true

		career.Aptitude = append(career.Aptitude[:index], career.Aptitude[index+1:]...)
		// fmt.Println(career.Aptitude)
	})

	skillEntries := []*widget.Entry{}

	skillBox := buildEditableList(career.Skill, &skillEntries, func (index int) {
		if index > len(career.Skill)-1 {
			return
		}

		// fmt.Println(skillEntries)
		repo.DeleteFact(career.CareerId, "", career.Skill[index], "")
		updateRequired = true

		career.Aptitude = append(career.Skill[:index], career.Skill[index+1:]...)
		// career.Skill[index] = ""	
	})

	interestEntries := []*widget.Entry{}

	interestBox := buildEditableList(career.Interest, &interestEntries, func (index int) {
		if index > len(career.Interest)-1 {
			return
		}

		// fmt.Println(interestEntries)
		repo.DeleteFact(career.CareerId, "", "", career.Interest[index])
		updateRequired = true

		career.Aptitude = append(career.Interest[:index], career.Interest[index+1:]...)
		// career.Interest[index] = ""	
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
		onDeleteCareer()
	})


	// Save and delete buttons
	saveBtn := widget.NewButton("Guardar", func() {
		var newAptitudes []string
		for _, entry := range aptitudeEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newAptitudes = append(newAptitudes, text)
			} else {
				navigation.ShowErrorDialog("No pueden haber apitudes vacias")
				return
			}
		}

		var newSkills []string
		for _, entry := range skillEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newSkills = append(newSkills, text)
			} else {
				navigation.ShowErrorDialog("No pueden haber habilidades vacias")
				return
			}
		}

		var newInterests []string
		for _, entry := range interestEntries {
			text := strings.TrimSpace(entry.Text)
			if text != "" {
				newInterests = append(newInterests, text)
			} else {
				navigation.ShowErrorDialog("No pueden haber intereses vacios")
				return
			}
		}

		// fmt.Println(career.Aptitude)
		// fmt.Println(newAptitudes)
		repo.UpdateCareer(career.CareerId, newAptitudes, newSkills, newInterests, career.Aptitude, career.Skill, career.Interest)

		career.Aptitude = newAptitudes
		career.Skill = newSkills
		career.Interest = newInterests
		// fmt.Println(career.Aptitude)
		updateRequired = true
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

func buildEditableList(items []string, entries *[]*widget.Entry, onClick func(index int)) (*fyne.Container) {
	list := container.NewVBox()
	footer := container.NewVBox()

	var refresh func()

	refresh = func() {
		list.Objects = nil
		for i, entry := range *entries {
			trash := widget.NewButtonWithIcon("", theme.DeleteIcon(), func(index int) func() {
				return func() {
					*entries = append((*entries)[:index], (*entries)[index+1:]...)
					refresh()
					onClick(index)
				}
			}(i))

			entryContainer := container.NewBorder(nil, nil, nil, trash, entry)
			list.Add(entryContainer)
		}
		list.Refresh()
	}

	// Initial entries
	for _, item := range items {
		if item == "" {
			continue
		}
		entry := widget.NewEntry()
		entry.SetText(item)
		*entries = append(*entries, entry)
	}
	refresh()

	// Agregar button
	addButton := widget.NewButton("Agregar", func() {
		newEntry := widget.NewEntry()
		*entries = append(*entries, newEntry)
		refresh()
	})
	footer.Add(addButton)

	scroll := container.NewVScroll(list)
	scroll.SetMinSize(fyne.NewSize(240, 300))

	full := container.NewVBox(
		scroll,
		footer,
	)

	return full
}

