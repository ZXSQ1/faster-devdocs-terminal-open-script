package utils

import (
	"os"
	"strings"
)

/*
description: gets the value of an environment variable
arguments:

	environVarName: the name of the variable to get the value of

return:
  - the value of the variable
  - an empty string if there is no such variable
*/
func GetEnvironVar(environVarName string) string {
	environVars := os.Environ()

	for _, environVar := range environVars {
		name := strings.Split(environVar, "=")[0]
		value := strings.Split(environVar, "=")[1]

		if name == environVarName {
			return value
		}
	}

	return ""
}
