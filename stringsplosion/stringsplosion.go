// stringsplosion is my solution to https://codingbat.com/prob/p117334
package stringsplosion

import (
	"strings"
)

// stringSplosion can process strings with non UTF-8 characters
func stringSplosion(str string) string {
	var exploded strings.Builder
	splitString := strings.Split(str, "")
	for index := range splitString {
		bytes := []byte(strings.Join(splitString[:index+1], ""))
		exploded.Write(bytes)
	}
	return exploded.String()
}

// stringSplosionUTF can only process strings with UTF-8 characters
func stringSplosionUTF(str string) string {
	byteArray := make([]byte, 0)
	for index := range str {
		byteArray = append(byteArray, []byte(str[0:index+1])...)
	}
	return string(byteArray)
}
