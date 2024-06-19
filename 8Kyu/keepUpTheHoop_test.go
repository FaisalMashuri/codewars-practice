package _Kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func KeepUpTheHoop(n int) string {
	if n >= 10 {
		return "Great, now move on to tricks"
	}
	return "Keep at it until you get it"
}

func TestKeepUpTheHoop(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		output string
	}{
		{
			name:   "test 1",
			input:  3,
			output: "Keep at it until you get it",
		},
		{
			name:   "test 2",
			input:  12,
			output: "Great, now move on to tricks",
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			res := KeepUpTheHoop(test.input)
			assert.Equal(t, test.output, res)
		})
	}
}
