package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/v2fly/geoip/lib"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"

	// Default config file changed to my preferred location
	configFile = flag.String("config", "config.json", "Path to the configuration file")
	listInput  = flag.Bool("list-input", false, "List all available input formats")
	listOutput = flag.Bool("list-output", false, "List all available output formats")
	version    = flag.Bool("version", false, "Print version and exit")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  geoip [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nFor more info, see: https://github.com/Loyalsoldier/geoip\n")
	}

	flag.Parse()

	if *version {
		fmt.Printf("geoip version %s\n", Version)
		os.Exit(0)
	}

	if *listInput {
		lib.ListInputFormats()
		os.Exit(0)
	}

	if *listOutput {
		lib.ListOutputFormats()
		os.Exit(0)
	}

	// Load and validate the configuration file
	instance, err := lib.NewInstance()
	if err != nil {
		log.Fatalf("failed to create instance: %v", err)
	}

	if err := instance.InitConfig(*configFile); err != nil {
		log.Fatalf("failed to initialize config from %q: %v", *configFile, err)
	}

	// Run the conversion pipeline
	if err := instance.Run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}

	log.Println("done")
}
