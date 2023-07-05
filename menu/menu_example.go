package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/civilware/Gnomon/structures"
	"github.com/dReam-dApps/dReams/menu"
	"github.com/dReam-dApps/dReams/rpc"
	"github.com/sirupsen/logrus"
)

// dReams menu StartGnomon() example

// Name my app
const app_tag = "My_app"

// Log output
var logger = structures.Logger.WithFields(logrus.Fields{})

func main() {
	// Initialize Gnomon fast sync
	menu.Gnomes.Fast = true

	// Initialize rpc address to rpc.Daemon var
	rpc.Daemon.Rpc = "127.0.0.1:10102"

	// Initialize logger to Stdout
	menu.InitLogrusLog(runtime.GOOS == "windows")

	rpc.Ping()
	// Check for daemon connection, if daemon is not connected we won't start Gnomon
	if rpc.Daemon.Connect {
		// Initialize NFA search filter and start Gnomon
		filter := []string{menu.NFA_SEARCH_FILTER}
		menu.StartGnomon(app_tag, "boltdb", filter, 0, 0, nil)

		// Exit with ctrl-C
		var exit bool
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			exit = true
		}()

		// Gnomon will continue to run if daemon is connected
		for !exit && rpc.Daemon.Connect {
			contracts := menu.Gnomes.GetAllOwnersAndSCIDs()
			logger.Printf("[%s] Index contains %d contracts\n", app_tag, len(contracts))
			time.Sleep(3 * time.Second)
			rpc.Ping()
		}

		// Stop Gnomon
		menu.Gnomes.Stop(app_tag)
	}

	logger.Printf("[%s] Done\n", app_tag)
}
