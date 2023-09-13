package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/dReam-dApps/dReams/dwidget"
	"github.com/dReam-dApps/dReams/menu"
	"github.com/dReam-dApps/dReams/rpc"
	"github.com/sirupsen/logrus"
)

// dReams dwidget NewVerticalEntries() example

// Name my app
const app_tag = "My_app"

func main() {
	// Initialize Gnomon fast sync
	menu.Gnomes.Fast = true

	// Initialize logger to Stdout
	menu.InitLogrusLog(logrus.InfoLevel)

	// Initialize fyne app
	a := app.New()

	// Initialize fyne window with size
	w := a.NewWindow(app_tag)
	w.Resize(fyne.NewSize(300, 100))
	w.SetMaster()

	// When window closes, stop Gnomon if running
	w.SetCloseIntercept(func() {
		if menu.Gnomes.Init {
			menu.Gnomes.Stop(app_tag)
		}
		w.Close()
	})

	// Initialize dwidget connection box
	connect_box := dwidget.NewVerticalEntries(app_tag, 1)

	// When connection button is pressed we will connect to wallet rpc,
	// and start Gnomon with NFA search filter if it is not running
	connect_box.Button.OnTapped = func() {
		rpc.GetAddress(app_tag)
		rpc.Ping()
		if rpc.Daemon.Connect && !menu.Gnomes.IsInitialized() && !menu.Gnomes.Start {
			go menu.StartGnomon(app_tag, "boltdb", []string{menu.NFA_SEARCH_FILTER}, 0, 0, nil)
		}
	}

	// Place connection box and start app
	w.SetContent(connect_box.Container)
	w.ShowAndRun()
}
