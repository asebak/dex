package main

import (
	"fmt"
	"os"
	"os/signal"

	"hameid.net/cdex/dex/internal/relayer"
)

func main() {
	app := relayer.NewRelayer(
		os.Getenv("DEX_VALIDATOR_CONTRACTS_FILE"),
		os.Getenv("DEX_VALIDATOR_NETWORKS_FILE"),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\nReceived signal %s", sig)
			app.Quit()
			os.Exit(0)
		}
	}()

	app.Initialize()

	done := make(chan bool)

	// app.RunOnBridgeNetwork()
	app.RunOnExchangeNetwork()

	<-done
}