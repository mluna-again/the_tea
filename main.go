package main

import (
	"flag"
	"fmt"
	"os"
	"the_tea/menu"
)

var item string

var helpMsg = `
Usage:

thetea -item menu
`

func usage() {
	fmt.Fprintln(os.Stderr, helpMsg)
	os.Exit(1)
}

func main() {
	flag.StringVar(&item, "item", "", "component to show case")
	flag.Parse()

	if item == "" {
		usage()
		return
	}

	switch item {
	case "menu":
		menu.Demo()
		return

	default:
		usage()
		return
	}
}
