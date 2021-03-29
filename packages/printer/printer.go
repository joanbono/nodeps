package printer

import "fmt"

func Header(appName, description, license, author string) {
	//print("header")
	fmt.Println(appName, description, license, author)
}
