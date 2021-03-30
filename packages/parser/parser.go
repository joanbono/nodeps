package parser

import (
	"io/ioutil"
	"strings"

	vsn "github.com/hashicorp/go-version"
	"github.com/joanbono/nodeps/packages/requests"
	"github.com/tidwall/gjson"
)

//Parser will parse the package.json file
func Parser(packageJson string) (string, string, string, string, [][]string, error) {
	var err error
	read, err := ioutil.ReadFile(packageJson)

	var table [][]string

	appName := gjson.Get(string(read), "name")
	description := gjson.Get(string(read), "description")
	license := gjson.Get(string(read), "license")
	author := gjson.Get(string(read), "author")

	//deps
	dependencies := gjson.Get(string(read), "dependencies")
	dependencies.ForEach(func(libName, libVersion gjson.Result) bool {
		lastVersion, repository, updated := CheckDependency(libName.Str, libVersion.Str)
		row := []string{updated, libName.Str, libVersion.Str, lastVersion, repository}
		table = append(table, row)
		return true
	})

	return appName.Str, description.Str, license.Str, author.Str, table, err
}

// CheckDependency will check the latest version
// on npms.io site
func CheckDependency(name, version string) (string, string, string) {
	var updated string
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
		println("Error getting versions")
	}

	if v1.LessThan(v2) {
		updated = "⚠️"
	} else if v1.Equal(v2) {
		updated = "✅"
	}

	return lastVersion.Str, repository.Str, updated
}
