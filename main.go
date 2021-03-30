package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/joanbono/nodeps/packages/parser"
	"github.com/joanbono/nodeps/packages/printer"
)

var red = color.New(color.FgRed)
var bold = color.New(color.FgHiWhite, color.Bold)

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
		fmt.Printf("\n\t🟢✝️  NODEPS %v\n\n", bold.Sprintf(version))
		return
	}

	appName, description, license, author, table, err := parser.Parser(packageJson)
	if err != nil {
		fmt.Printf("\n%v can't open %v\n\n", red.Sprintf("[-] ERROR: "), bold.Sprintf(packageJson))
	} else {
		printer.Table(appName, description, license, author, table)
	}
}
