package stringsplosion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsplosion(t *testing.T) {
	// TEST: strings possibly containing non UTF-8 characters
	inputs := []string{"Code", "abc", "äb"}
	want := []string{"CCoCodCode", "aababc", "ääb"}
	for index, input := range inputs {
		assert.Equal(t, want[index], stringSplosion(input), "String should explode correctly")
	}

	// TEST: strings possibly containing non UTF-8 characters
	inputs = []string{"Code", "abc", "ab"}
	want = []string{"CCoCodCode", "aababc", "aab"}
	for index, input := range inputs {
		assert.Equal(t, want[index], stringSplosionUTF(input), "String should explode correctly")
	}
}
