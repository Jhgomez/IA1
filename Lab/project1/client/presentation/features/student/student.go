package student

import (
	"regexp"
	"fmt"
	"strconv"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"

	"unimatch/presentation/navigation"
)

type Question struct {
	Text         string
	Options      []string
	Required     int
	CheckBoxes   []*widget.Check
	Label        *canvas.Text
	ContainerBox *fyne.Container // used later for styling
}

var questions []Question

var interestEntries map[string]*widget.Entry

func validateForm() {
	// Validate interest entries sum
	total := 0
	for label, entry := range interestEntries {
		text := entry.Text
		if text == "" {
			text = "0"
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			msg := fmt.Sprintf("Valor inválido en el campo de interés: %s", label)
			// dialog.ShowError(msg, w)
			navigation.ShowErrorDialog(msg)
			return
		}
		total += value
	}

	if total != 100 {
		// dialog.ShowError(errors.New("Los porcentajes deben sumar exactamente 100%"), w)
		navigation.ShowErrorDialog("Los porcentajes deben sumar exactamente 100%")
		return
	}

	// Validate questions
	for i, q := range questions {
		// Count checked boxes
		selected := 0
		for _, c := range q.CheckBoxes {
			if c.Checked {
				selected++
			}
		}

		// Reset label to normal
		q.Label.Color = theme.ForegroundColor()
		q.Label.TextStyle = fyne.TextStyle{}
		q.Label.Refresh()


		if selected != q.Required {
			// Bold label to highlight the error
			q.Label.Color = color.NRGBA{R: 255, G: 0, B: 0, A: 255} // red
			q.Label.TextStyle = fyne.TextStyle{Bold: true}
			q.Label.Refresh()


			msg := fmt.Sprintf("La pregunta '%d' requiere seleccionar %d opción(es).", i + 1, q.Required)
			navigation.ShowErrorDialog(msg)
			return
		}
	}

	// ✅ Everything is valid
	// navigation.ShowDialog("Exito", "Todos los campos han sido validados correctamente")
	// navigation.NavigateWithNewWindow(
	// 	"Formulary",  //windowTitle
	// 	studentSecondaryFormulary(),  //content
	// 	true, // shouldHide
	// 	fyne.NewSize(700, 500), // windows size
	// 	nil, // onClose
	// )
}

func StudentPrimaryFormulary() fyne.CanvasObject {
	// Title
	title := canvas.NewText("Bienvenido", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 28 // default is 14

	// // Title
	// title := widget.NewLabelWithStyle("Career Path Assistant", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Description
	description := widget.NewLabelWithStyle(
	"Selecciona la informacion que mas te \nrepesenta y te describe",
		fyne.TextAlignCenter,
		fyne.TextStyle{Italic: true},
	)

	
}


func studenSecondaryFormulary() fyne.CanvasObject {
	// Title
	title := canvas.NewText("Bienvenido", theme.PrimaryColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 28 // default is 14

	// // Title
	// title := widget.NewLabelWithStyle("Career Path Assistant", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Description
	description := widget.NewLabelWithStyle(
	"Selecciona la informacion que mas te \nrepesenta y te describe",
		fyne.TextAlignCenter,
		fyne.TextStyle{Italic: true},
	)

	// ---------- Left Interests Section ----------
	interests := []string{"Science", "Art", "Technology", "Business", "Education"}
	interestEntries = make(map[string]*widget.Entry)
	var interestRows []fyne.CanvasObject

	spacer1 := canvas.NewRectangle(nil)
	spacer1.SetMinSize(fyne.NewSize(80, 10))

	for _, interest := range interests {
		label := widget.NewLabel(interest)

		entry := widget.NewEntry()
		entry.SetPlaceHolder("0–100")
		entry.Resize(fyne.NewSize(90, 36))

		entryContainer := container.NewWithoutLayout(entry)

		entry.OnChanged = func(s string) {
			clean := regexp.MustCompile(`\D`).ReplaceAllString(s, "")
			if len(clean) > 3 {
				clean = clean[1:4]
			}
			
			entry.SetText(clean)
		}

		interestEntries[interest] = entry

		row := container.NewHBox(label, layout.NewSpacer() , entryContainer, spacer1)
		interestRows = append(interestRows, row)
	}


	leftScroll := container.NewVScroll(container.NewVBox(interestRows...))
	leftScroll.SetMinSize(fyne.NewSize(300, 300)) // Optional: set a fixed size

	leftDescription := widget.NewLabelWithStyle(
	"Ingresa el porcentaje de interes \n sobre los siguientes campos",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	spacer := canvas.NewRectangle(nil)
	spacer.SetMinSize(fyne.NewSize(10, 10))

	leftColumn := container.NewVBox(
		leftDescription,
		spacer,
		leftScroll,
	)

	// ---------- Right Questions Section ----------

	questions = []Question{
		{
			Text:     "¿Qué ambientes prefieres? (Selecciona 2)",
			Options:  []string{"Interior", "Exterior", "Remoto", "Híbrido"},
			Required: 2,
		},
		{
			Text:     "¿Qué tipo de tareas disfrutas? (Selecciona 1)",
			Options:  []string{"Creativas", "Analíticas", "Prácticas", "Sociales"},
			Required: 1,
		},
	}

	var questionBoxes []fyne.CanvasObject
	checkboxContainer := container.NewVBox()

	for i := range questions {
		q := &questions[i]

		// Label
		q.Label = canvas.NewText(q.Text, theme.ForegroundColor())
		q.Label.TextStyle = fyne.TextStyle{Bold: false}


		checkboxGroup := container.NewVBox()

		// Checkboxes
		q.CheckBoxes = []*widget.Check{}
		for i, opt := range q.Options {
			check := widget.NewCheck(opt, nil)
			q.CheckBoxes = append(q.CheckBoxes, check)
			checkboxGroup.Add(q.CheckBoxes[i])

			check.OnChanged = func(checked bool) {
				if !checked {
					return // always allow unchecking
				}

				// Count how many are selected
				count := 0
				for _, c := range q.CheckBoxes {
					if c.Checked {
						count++
					}
				}

				// If too many are selected, uncheck this one
				if count > q.Required {
					check.SetChecked(false)
				}
			}
		}

		// Wrap question components in a VBox
		q.ContainerBox = container.NewVBox(
			q.Label,
			checkboxGroup,
		)

		questionBoxes = append(questionBoxes, q.ContainerBox)
		checkboxContainer.Add(questionBoxes[i])
	}

	
	rightScroll := container.NewVScroll(checkboxContainer)

	rightScroll.SetMinSize(fyne.NewSize(300, 300)) // fixed width for symmetry

	rightTitle := canvas.NewText("Preferencias", theme.ForegroundColor())
	rightTitle.TextStyle = fyne.TextStyle{Bold: true}
	rightTitle.TextSize = 20
	rightTitle.Alignment = fyne.TextAlignCenter

	rightColumn := container.NewVBox(rightTitle, spacer1, rightScroll)

	validateBtn := widget.NewButton("Validar", func() {
		validateForm()
	})

	bottom := container.NewHBox(
		layout.NewSpacer(), // pushes button to the right
		validateBtn,
		spacer,
	)

	// ---------- Layout ----------
	columns := container.NewHBox(
		leftColumn,
		spacer,
		rightColumn,
	)

	layoutContent := container.NewVBox(
		title,
		description,
		layout.NewSpacer(),
		columns,
		spacer,
		bottom,
		spacer,
	)

	return layoutContent
}
