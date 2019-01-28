package stringsplosion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsplosion(t *testing.T) {
	inputs := []string{"Code", "abc", "äb", "ぅべほけ"}
	want := []string{"CCoCodCode", "aababc", "ääb", "ぅぅべぅべほぅべほけ"}
	for index, input := range inputs {
		assert.Equal(t, want[index], stringSplosion(input), "String should explode correctly")
	}
}
