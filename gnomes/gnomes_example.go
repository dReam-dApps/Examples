package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/civilware/Gnomon/structures"
	"github.com/dReam-dApps/dReams/gnomes"
	"github.com/dReam-dApps/dReams/rpc"
	"github.com/sirupsen/logrus"
)

// dReams gnomes StartGnomon() example

// Name my app
const app_tag = "My_app"

// Log output
var logger = structures.Logger.WithFields(logrus.Fields{})

// Gnomon instance from gnomes package
var gnomon = gnomes.NewGnomes()

func main() {
	// Initialize Gnomon fast sync true to sync db immediately
	gnomon.SetFastsync(true, false, 100)

	// Initialize rpc address to rpc.Daemon var
	rpc.Daemon.Rpc = "127.0.0.1:10102"

	// Initialize logger to Stdout
	gnomes.InitLogrusLog(logrus.InfoLevel)

	rpc.Ping()
	// Check for daemon connection, if daemon is not connected we won't start Gnomon
	if rpc.Daemon.IsConnected() {
		// Initialize NFA search filter and start Gnomon
		filter := []string{gnomes.NFA_SEARCH_FILTER}
		gnomes.StartGnomon(app_tag, "boltdb", filter, 0, 0, nil)

		// Exit with ctrl-C
		var exit bool
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			exit = true
		}()

		// Gnomon will continue to run if daemon is connected
		for !exit && rpc.Daemon.IsConnected() {
			contracts := gnomon.GetAllOwnersAndSCIDs()
			logger.Printf("[%s] Index contains %d contracts\n", app_tag, len(contracts))
			time.Sleep(3 * time.Second)
			rpc.Ping()
		}

		// Stop Gnomon
		gnomon.Stop(app_tag)
	}

	logger.Printf("[%s] Done\n", app_tag)
}
