package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dReam-dApps/dReams/rpc"
)

// dReams rpc connection example

// Name my app
const app_tag = "My_app"

func main() {
	// Initialize rpc addresses to rpc.Daemon and rpc.Wallet vars
	rpc.Daemon.Rpc = "127.0.0.1:10102"
	rpc.Wallet.Rpc = "127.0.0.1:10103"
	// Initialize rpc.Wallet.UserPass for rpc user:pass

	// Check for daemon connection
	rpc.Ping()

	// Check for wallet connection and get address
	rpc.GetAddress(app_tag)

	// Exit with ctrl-C
	var exit bool
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("[%s] Closing\n", app_tag)
		exit = true
	}()

	// Loop will check for daemon and wallet connection and
	// print wallet height and balance. It will keep
	// running while daemon and wallet are connected or until exit
	for !exit && rpc.IsReady() {
		rpc.Wallet.GetBalance()
		rpc.GetWalletHeight(app_tag)
		log.Printf("[%s] Height: %d   Dero Balance: %s\n", app_tag, rpc.Wallet.Height, rpc.FromAtomic(rpc.Wallet.Balance, 5))
		time.Sleep(3 * time.Second)
		rpc.Ping()
		rpc.EchoWallet(app_tag)
	}

	log.Printf("[%s] Not connected\n", app_tag)
}
