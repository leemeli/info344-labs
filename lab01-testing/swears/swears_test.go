package main

import "testing"

func TestCountSwears(t *testing.T) {

	cases := []struct {
		name           string
		input          []string
		expectedOutput []SwearCounts
	}{
		{
			name:           "empty KnownSwears",
			input:          []string{},
			expectedOutput: []SwearCounts{},
		},
		{
			name:           "non-empty KnownSwears",
			input:          []string{"Ha", "fuck", "bitch", "lol"},
			expectedOutput: []SwearCounts{"fuck": 1},
		},
	}
	knownSwears := loadKnownSwears("known_swears.txt")
	for _, c := range cases {
		if output := countSwears(knownSwears, c.input); output != c.expectedOutput {
			t.Errorf("%s: got %d but expected %d", c.name, output, c.expectedOutput)
		}
	}
}

func TestLoadKnownSwears(t *testing.T) {

}

func TestLoadWords(t *testing.T) {

}

func TestOpenFile(t *testing.T) {
	// hint os.Args = []string{"cmd", "path/to/file.txt"}
}
