package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	vsn "github.com/hashicorp/go-version"
	"github.com/joanbono/nodeps/packages/requests"
	"github.com/tidwall/gjson"
)

// Parser will parse the package.json file and look for the latest version on npmjs
func Parser(packageJson string) (string, string, string, string, [][]string, error) {
	var err error
	var table [][]string
	read, err := ioutil.ReadFile(packageJson)

	appName := gjson.Get(string(read), "name")
	description := gjson.Get(string(read), "description")
	license := gjson.Get(string(read), "license")
	author := gjson.Get(string(read), "author")

	//fmt.Printf("%v", appName.Str)

	dependencies := gjson.Get(string(read), "dependencies")
	dependencies.ForEach(func(libName, libVersion gjson.Result) bool {
		results := requests.MakeRequest(fmt.Sprintf(`["%s"]`, libName.Str))
		latestVersion := gjson.Get(fmt.Sprintf("%s", results), fmt.Sprintf("%s.collected.metadata.version", libName.Str))
		repository := gjson.Get(string(string(results)), fmt.Sprintf("%s.collected.metadata.links.repository", libName.Str))
		updated := CheckDependency(libVersion.Str, latestVersion.Str)
		row := []string{updated, libName.Str, libVersion.Str, latestVersion.Str, repository.Str}
		table = append(table, row)

		return true
	})

	return appName.Str, description.Str, license.Str, author.Str, table, err
}

// CheckDependency will compare current version with latest version
func CheckDependency(installedVersion, latestVersion string) string {
	var updated string

	removeTilde := strings.Replace(installedVersion, "~", ``, -1)
	oldVersion := strings.Replace(removeTilde, "^", ``, -1)

	v1, err := vsn.NewVersion(oldVersion)
	v2, err := vsn.NewVersion(latestVersion)

	if err != nil {
		//println("Error getting versions")
		updated = "❓"
	} else {
		if v1.LessThan(v2) {
			updated = "⚠️"
		} else if v1.Equal(v2) {
			updated = "✅"
		}
	}
	return updated
}
