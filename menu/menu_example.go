package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	dreams "github.com/dReam-dApps/dReams"
	"github.com/dReam-dApps/dReams/bundle"
	"github.com/dReam-dApps/dReams/menu"
)

// dReams menu PlaceMarket and PlaceAsset example

// Name my app
const app_tag = "My_app"

func main() {
	// Intialize Fyne window app and window into dReams app object
	a := app.New()
	w := a.NewWindow(app_tag)
	w.Resize(fyne.NewSize(900, 700))
	d := dreams.AppObject{
		App:    a,
		Window: w,
	}

	// Simple asset profile with wallet name entry and theme select
	line := canvas.NewLine(bundle.TextColor)
	profile := []*widget.FormItem{}
	profile = append(profile, widget.NewFormItem("Name", menu.NameEntry()))
	profile = append(profile, widget.NewFormItem("", layout.NewSpacer()))
	profile = append(profile, widget.NewFormItem("", container.NewVBox(line)))
	profile = append(profile, widget.NewFormItem("Theme", menu.ThemeSelect(&d)))
	profile = append(profile, widget.NewFormItem("", container.NewVBox(line)))

	// Rescan button function in asset tab
	rescan := func() {
		// What you want to scan wallet for
	}

	// Place asset and market layouts into tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Assets", menu.PlaceAssets(app_tag, widget.NewForm(profile...), rescan, bundle.ResourceDReamsIconAltPng, &d)),
		container.NewTabItem("Market", menu.PlaceMarket(&d)))

	// Place tabs as window content and run app
	d.Window.SetContent(tabs)
	d.Window.ShowAndRun()
}
