// This is a sample application for showing how to use the xval package
package main

import (
	"flag"
	// Since there's also a global var named "xval", this package will be aliased as "xvdp" (xval package)
	xvp "github.com/landaire/xval"
	"fmt"
)

const (
	serialUsage = "Console's serial number"
	xvalUsage = "Console's X value displayed in the console's system info"
)

var (
	serial = flag.String("serial", "", serialUsage)
	xval   = flag.String("xval", "", xvalUsage)
)

func init() {
	// Set up short flags so that e.g. either -s or -serial can be used
	flag.StringVar(serial, "s", "", serialUsage)
	flag.StringVar(xval, "x", "", xvalUsage)
}

func main() {
	flag.Parse()

	if *serial == "" || *xval == "" {
		flag.Usage()
		return
	}

	// The xval package does validation on the serial/xval
	key, decrpytedXval, err := xvp.Decrypt(*serial, *xval)

	if err != nil {
		fmt.Errorf("The following error occurred: %s", err)
		return
	}

	flagDescriptions := xvp.TextResult(decrpytedXval)

	fmt.Printf("Decrypted xval: 0x%X\n", decrpytedXval)
	fmt.Printf("Decryption key: 0x%X\n", key)
	fmt.Println("Results:")

	for _, description := range flagDescriptions {
		fmt.Printf("  - %s", description)
	}
}
