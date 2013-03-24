package testing

import (
	"fmt"
	"gompress/helper"
)

func StartTesting(args []string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func TestSanitizedName(filename []string) {
	newName := make([]string, len(filename))
	for key, name := range filename {
		newName[key] = helper.SanitizedName(name)
	}
	for _, name := range newName {
		fmt.Println(name)
	}
}
