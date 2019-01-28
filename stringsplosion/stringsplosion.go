// stringsplosion is my solution to https://codingbat.com/prob/p117334
package stringsplosion

func stringSplosion(str string) string {
	var exploded string
	runeSlice := []rune(str)
	for index := range runeSlice {
		exploded += string(runeSlice[:index+1])
	}
	return exploded
}
