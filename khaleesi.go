package khaleesi

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// Mock replaces all Cyrillic characters in the input string with
// mock Cyrillic characters.
func Mock(input string) (string, bool, error) {
	kh, err := New()
	if err != nil {
		return "", false, err
	}

	output, mocked := kh.Mock(input)
	return output, mocked, nil
}

// Khaleesi is a mock Cyrillic generator.
type Khaleesi struct {
	re   *regexp.Regexp
	subs map[string][]string
}

// New creates a new Khaleesi instance.
func New() (*Khaleesi, error) {
	kh := Khaleesi{
		subs: replaces(),
	}

	expr := strings.Join(keys(kh.subs), "|") + "gi"

	var err error
	kh.re, err = regexp.Compile(expr)
	if err != nil {
		return nil, fmt.Errorf("compile regexp: %w", err)
	}

	return &kh, nil
}

// Mock replaces all Cyrillic characters in the input string with
// mock Cyrillic characters.
func (k *Khaleesi) Mock(input string) (string, bool) {

	if len(input) == 0 {
		return input, false
	}

	output := k.re.ReplaceAllStringFunc(input, func(s string) string {
		ss := k.subs[s]
		idx := rand.Intn(len(ss))
		return ss[idx]
	})

	return output, input != output
}
