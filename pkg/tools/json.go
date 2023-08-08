package tools

import (
	"encoding/json"
	"strings"

	"github.com/iancoleman/orderedmap"
)

// IndentJSON takes a JSON string and the number of indentations and returns
// the indented JSON string.
func IndentJSON(jsonStr string, indent int) (string, error) {
	// Create an ordered map
	orderedMap := orderedmap.New()
	if err := orderedMap.UnmarshalJSON([]byte(jsonStr)); err != nil {
		return "", err
	}

	// Determine whether to use tabs or spaces for indentation
	indentString := " "
	if indent%4 == 0 {
		indentString = "\t"
		indent /= 4
	}

	// Indent the ordered map to a string
	prettyBytes, err := json.MarshalIndent(orderedMap, "", strings.Repeat(indentString, indent))
	if err != nil {
		return "", err
	}

	prettyString := string(prettyBytes[:])

	// Print the indented JSON
	return prettyString, nil
}
