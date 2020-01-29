package lib

import (
	"encoding/json"
	"strings"
)

//IsJSON : check if valid json
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

//RemoveNewLine : removes new line character
func RemoveNewLine(char string) string {

	char = strings.TrimPrefix(char, "\n") //remove new line from beginning
	char = strings.TrimSuffix(char, "\n") //remove new line from end of line

	return char
}
