package printer

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var cyan = color.New(color.FgCyan)
var bold = color.New(color.FgHiWhite, color.Bold)

func PackageInfo(appName, description, license, author string) {
	//print("header")
	fmt.Println(appName, description, license, author)
}

func Table(appName, description, license, author string, table [][]string) {
	println()
	println()
	tablePrint := tablewriter.NewWriter(os.Stdout)
	tablePrint.SetHeader([]string{"", " Name ", " Used Version ", " Last Version ", " Repository "})
	tablePrint.SetHeaderColor(tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.BgBlueColor, tablewriter.Bold, tablewriter.FgHiWhiteColor})

	for _, v := range table {
		tablePrint.Append(v)
	}

	tablePrint.SetBorders(tablewriter.Border{Left: false, Top: false, Right: true, Bottom: false})
	tablePrint.SetAlignment(tablewriter.ALIGN_LEFT)
	tablePrint.Render()
	println()
	println()
}
