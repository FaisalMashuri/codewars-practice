package _Kyu

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func MakeUpperCase(s string) string {
	return strings.ToUpper(s)
}

func TestMakeUpperCase(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "test 1",
			input:  "hello",
			output: "HELLO",
		},
		{
			name:   "test 2",
			input:  "hello world",
			output: "HELLO WORLD",
		},
		{
			name:   "test 3",
			input:  "hello world !",
			output: "HELLO WORLD !",
		},
		{
			name:   "test 4",
			input:  "heLlO wORLd !",
			output: "HELLO WORLD !",
		},
		{
			name:   "test 5",
			input:  "1,2,3 hello world!",
			output: "1,2,3 HELLO WORLD!",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			res := MakeUpperCase(test.input)
			assert.Equal(t, test.output, res)
		})
	}

}
