package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dReam-dApps/dReams/bundle"
)

// Name my app
const app_tag = "My_app"

func main() {
	// Initialize app color to bundle var
	bundle.AppColor = color.Black

	// Initialize fyne app with Dero theme
	a := app.New()
	a.Settings().SetTheme(bundle.DeroTheme(bundle.AppColor))

	// Initialize fyne window with size and icon from bundle package
	w := a.NewWindow(app_tag)
	w.SetIcon(bundle.ResourceBlueBadge3Png)
	w.Resize(fyne.NewSize(300, 100))
	w.SetMaster()

	// Initialize fyne container and add some various widgets for viewing purposes
	cont := container.NewVBox()
	cont.Add(container.NewAdaptiveGrid(3, widget.NewLabel("Label"), widget.NewEntry(), widget.NewButton("Button", nil)))
	cont.Add(container.NewAdaptiveGrid(3, widget.NewLabel("Label"), widget.NewCheck("Check", nil), widget.NewButton("Button", nil)))
	cont.Add(widget.NewPasswordEntry())
	cont.Add(widget.NewSlider(0, 100))

	// Widget to change theme
	change_theme := widget.NewRadioGroup([]string{"Dark", "Light"}, func(s string) {
		switch s {
		case "Dark":
			bundle.AppColor = color.Black
		case "Light":
			bundle.AppColor = color.White
		default:

		}

		a.Settings().SetTheme(bundle.DeroTheme(bundle.AppColor))
	})
	change_theme.Horizontal = true
	cont.Add(container.NewCenter(change_theme))

	// Add a image from bundle package
	gnomon_img := canvas.NewImageFromResource(bundle.ResourceGnomonIconPng)
	gnomon_img.SetMinSize(fyne.NewSize(45, 45))
	cont.Add(container.NewCenter(gnomon_img))

	// Adding last widget
	select_entry := widget.NewSelect([]string{"Choice 1", "Choice 2", "Choice 3"}, nil)
	cont.Add(select_entry)

	// Place widget container and start app
	w.SetContent(cont)
	w.ShowAndRun()
}
