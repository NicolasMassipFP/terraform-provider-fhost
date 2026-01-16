// Package main implements the Terraform provider for SMC (Security Management Center).
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/terraform-providers/terraform-provider-smc/internal/provider"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version = "0.0.1"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// Set up cleanup handler for graceful shutdown
	// This ensures all SMC API sessions are properly closed when the provider exits
	setupCleanupHandler()

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		// Also update the tfplugindocs generate command to either remove the
		// -provider-name flag or set its value to the updated provider name.
		Address: "registry.terraform.io/forcepoint/smc",
		Debug:   debug,
	}

	// errors not caught. best effort cleanup
	provider.DeletePendingResources(context.Background())

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Cleanup on normal exit
	cleanup()
}

// setupCleanupHandler registers signal handlers and cleanup routines
func setupCleanupHandler() {
	// Create a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Start a goroutine to handle cleanup on interrupt
	go func() {
		<-sigChan
		log.Println("Received interrupt signal, cleaning up SMC connections...")
		cleanup()
		os.Exit(0)
	}()
}

// cleanup logs out all active SMC client sessions
func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	smc.SaveAndCloseVpns(ctx)

	// try to delete resources that could not be deleted due to
	// pending dependencies
	provider.DeletePendingResources(ctx)

	// logout from all the providers
	smc.GetClientManager().CleanupAll(ctx)
}
