package stringsplosion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsplosion(t *testing.T) {
	inputs := []string{"Code", "abc", "ab"}
	want := []string{"CCoCodCode", "aababc", "aab"}
	for index, input := range inputs {
		assert.Equal(t, want[index], stringSplosion(input), "String should explode correctly")
	}
}
