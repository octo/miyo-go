package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/octo/miyo-go/miyo"
)

var (
	address = flag.String("addr", os.Getenv("MIYO_ADDRESS"), "address of the Miyo cube")
	apiKey  = flag.String("apikey", os.Getenv("MIYO_APIKEY"), "API key of the Miyo cube")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	if *address == "" || *apiKey == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s -addr=<addr> -apikey=<apikey>\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := miyo.Connect(ctx, *address, miyo.WithAPIKey(*apiKey))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("# Circuits")
	fmt.Println()
	cc, err := conn.CircuitAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cc {
		fmt.Printf("*   %s (%s)\n", c.Name, c.Status())
	}
	fmt.Println()

	fmt.Println("# Devices")
	fmt.Println()
	devs, err := conn.DeviceAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devs {
		fmt.Printf("*   %s %s (%s)\n", dev.Type, dev.ID, dev.Status())
	}
	fmt.Println()
}
