package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	vsn "github.com/hashicorp/go-version"
	"github.com/joanbono/nodeps/packages/printer"
	"github.com/joanbono/nodeps/packages/requests"
	"github.com/tidwall/gjson"
)

type PackageJsonStruct struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Version      string      `json:"version"`
	License      string      `json:"license"`
	Author       string      `json:"author"`
	Main         string      `json:"main"`
	Dependencies interface{} `json:"dependencies"`
}

//parser will parse the package.json file
func Parser(packageJson string) {
	read, err := ioutil.ReadFile(packageJson)
	if err != nil {
		fmt.Printf("cannot open %s\v", packageJson)
		os.Exit(1)
	}

	appName := gjson.Get(string(read), "name")
	description := gjson.Get(string(read), "description")
	license := gjson.Get(string(read), "license")
	author := gjson.Get(string(read), "author")

	printer.Header(appName.Str, description.Str, license.Str, author.Str)

	//deps
	dependencies := gjson.Get(string(read), "dependencies")
	dependencies.ForEach(func(libName, libVersion gjson.Result) bool {
		println(libName.Str, libVersion.Str)
		lastVersion, repository, updated := CheckDependency(libName.Str, libVersion.Str)
		fmt.Printf("\n%v - %v - %v - %v\n\n", lastVersion, libVersion, updated, repository)
		os.Exit(2)
		return true
	})

}

func CheckDependency(name, version string) (string, string, bool) {
	var updated bool
	var url = `https://api.npms.io/v2/package/` + name
	response := requests.MakeRequest(url)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	lastVersion := gjson.Get(string(string(bodyBytes)), "collected.metadata.version")
	repository := gjson.Get(string(string(bodyBytes)), "collected.metadata.links.repository")

	removeTilde := strings.Replace(version, "~", ``, -1)
	oldVersion := strings.Replace(removeTilde, "^", ``, -1)

	v1, err := vsn.NewVersion(oldVersion)
	v2, err := vsn.NewVersion(lastVersion.Str)

	if err != nil {
		println(err)
	}

	if v1.LessThan(v2) {
		updated = false
	} else if v1.Equal(v2) {
		updated = true
	}

	return lastVersion.Str, repository.Str, updated
}
