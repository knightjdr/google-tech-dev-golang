// stringsplosion is my solution to https://codingbat.com/prob/p117334
package stringsplosion

import "strings"

func stringSplosion(str string) string {
	var exploded strings.Builder
	splitString := strings.Split(str, "")
	for index := range splitString {
		bytes := []byte(strings.Join(splitString[:index+1], ""))
		exploded.Write(bytes)
	}
	return exploded.String()
}
