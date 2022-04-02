package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/octo/miyo-go/miyo"
)

var (
	address = flag.String("addr", os.Getenv("MIYO_ADDRESS"), "address of the Miyo cube")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	if *address == "" {
		flag.Usage()
		os.Exit(1)
	}

	apiKey, err := miyo.APIKey(ctx, *address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Requesting API key failed: %v\n", err)
		fmt.Fprintln(os.Stderr, "Press the physical button on the MIYO Cube and try again.")
		os.Exit(1)
	}

	fmt.Println(apiKey)
}
