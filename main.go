package main

import (
	"flag"
	"fmt"

	"github.com/joanbono/nodeps/packages/parser"
)

var (
	packageJson string
	versionFlag bool
	version     string
)

func init() {
	flag.StringVar(&packageJson, "pkg", "package.json", "package.json file")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Printf("\nNODEPS %v\n\n", version)
		return
	}
	// if packageJson == "" {
	// 	flag.PrintDefaults()
	// 	return
	// }

	parser.Parser(packageJson)
}
