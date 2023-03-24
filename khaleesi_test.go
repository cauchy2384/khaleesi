package khaleesi

import (
	"testing"

	"github.com/matryer/is"
)

func TestMock(t *testing.T) {
	t.Parallel()

	is := is.New(t)

	testCases := []struct {
		input  string
		mocked bool
	}{
		{"", false},
		{"Lorem ipsum dolor sit amet", false},
		{"Съешь еще этих мягких французских булок, да выпей чаю", true},
		{"Позвольте мне сражаться за вас, Кхалиси!", true},
		{"Дерись за меня дракон!", true},
	}

	for _, tt := range testCases {
		output, mocked, err := Mock(tt.input)
		is.NoErr(err)
		is.Equal(tt.mocked, mocked)
		if tt.mocked {
			is.True(tt.input != output)
		}
	}
}
