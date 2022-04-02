// setup helps setting up a MIYO Cube by discovering the device's address and requesting an API key.
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

	if addr := *address; addr == "" {
		addr, err := miyo.FindCube(ctx)
		if err != nil {
			log.Fatalf("miyo.FindCube(): %v", err)
		}

		fmt.Printf("MIYO_ADDRESS=%q; export MIYO_ADDRESS;\n", addr)
		fmt.Printf("echo \"MIYO Cube address: %s\";\n", addr)
		*address = addr
	}

	if ak := *apiKey; ak == "" {
		ak, err := miyo.APIKey(ctx, *address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Requesting API key failed: %v\n", err)
			fmt.Fprintln(os.Stderr, "Press the physical button on the MIYO Cube and try again.")
			os.Exit(1)
		}

		fmt.Printf("MIYO_APIKEY=%q; export MIYO_APIKEY;\n", ak)
		fmt.Printf("echo \"MIYO API key: %s\";\n", ak)
	}
}
