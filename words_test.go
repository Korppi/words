package words

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult bool
	}{
		{"Test", false},
		{"", true},
		{"A", true},
		{"AA", true},
		{"AAA", true},
		{"The eht", true},
		{"E ", false},
		{"rotator", true},
		{"ROTator", true},
	}
	for _, test := range testCases {
		result := IsPalindrome(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%s'", result, test.expectedResult, test.input)
		}
	}
}

func TestIsPalindromeCaseSensitive(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult bool
	}{
		{"Test", false},
		{"", true},
		{"A", true},
		{"AA", true},
		{"AAA", true},
		{"The eht", false},
		{"E ", false},
		{"rotator", true},
		{"ROTator", false},
	}
	for _, test := range testCases {
		result := IsPalindromeCaseSensitive(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%s'", result, test.expectedResult, test.input)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult string
	}{
		{"Test", "tseT"},
		{"", ""},
		{"123456789", "987654321"},
		{"0", "0"},
		{"0 ", " 0"},
		{" 0 ", " 0 "},
	}
	for _, test := range testCases {
		reversed := Reverse(test.input)
		if reversed != test.expectedResult {
			t.Errorf("Got '%s' but expected '%s' for '%s'", reversed, test.expectedResult, test.input)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		input1         string
		input2         string
		expectedResult bool
	}{
		{"rat", "tAr", true},
		{"Hello, World!", "World, hello!", true},
		{"rat", "tAre", true},
		{"rate", "tAr", false},
		{"rat", "tar", true},
		{"rat", "tar", true},
		{"Tom Marvolo Riddle", "I am Lord Voldemort", true},
	}

	for _, test := range testCases {
		result := IsAnagram(test.input1, test.input2)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%s' and '%s'", result, test.expectedResult, test.input1, test.input2)
		}
	}
}

func TestWordLength(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult int
	}{
		{"", 0},
		{" ", 1},
		{"1", 1},
		{"asd", 3},
		{"Hello, 世🖖", 9},
	}

	for _, test := range testCases {
		result := WordLength(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%s'", result, test.expectedResult, test.input)
		}
	}
}

func TestIsPangram(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"abcdefghijklmnopqrstuvwxy", false},
		{"abcdefghijklmnopqrstuvwxyzABCEDFLKF", true},
		{"ABCdefghijklmnopqrstuvwxyz", true},
	}

	for _, test := range testCases {
		result := IsPangram(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%s'", result, test.expectedResult, test.input)
		}
	}
}

func TestIsVowel(t *testing.T) {
	testCases := []struct {
		input          rune
		expectedResult bool
	}{
		{'a', true},
		{'e', true},
		{'i', true},
		{'o', true},
		{'u', true},
		{'A', true},
		{'E', true},
		{'I', true},
		{'O', true},
		{'U', true},
		{'X', false},
	}

	for _, test := range testCases {
		result := IsVowel(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%v'", result, test.expectedResult, test.input)
		}
	}
}

func TestIsConsonant(t *testing.T) {
	testCases := []struct {
		input          rune
		expectedResult bool
	}{
		{'a', false},
		{'e', false},
		{'i', false},
		{'o', false},
		{'u', false},
		{'A', false},
		{'E', false},
		{'I', false},
		{'O', false},
		{'U', false},
		{'X', true},
	}

	for _, test := range testCases {
		result := IsConsonant(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%v'", result, test.expectedResult, test.input)
		}
	}
}

func TestCountConsonants(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult int
	}{
		{"", 0},
		{"iii", 0},
		{"asd", 2},
		{"I do like this", 6},
		{"RRttrT", 6},
	}

	for _, test := range testCases {
		result := CountConsonants(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%v'", result, test.expectedResult, test.input)
		}
	}
}

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult int
	}{
		{"", 0},
		{"iii", 3},
		{"asd", 1},
		{"I do like this", 5},
		{"RRttrT", 0},
	}

	for _, test := range testCases {
		result := CountVowels(test.input)
		if result != test.expectedResult {
			t.Errorf("Got '%v' but expected '%v' for '%v'", result, test.expectedResult, test.input)
		}
	}
}

func TestSearchPossibleWords(t *testing.T) {
	testCases := []struct {
		letters       string
		expectedCount int
	}{
		{"aa", 2},
		{"a", 1},
	}
	for _, test := range testCases {
		result := SearchPossibleWords(test.letters)
		if len(result) != test.expectedCount {
			t.Errorf("Got '%v' but expected '%v' for '%s'", result, test.expectedCount, test.letters)
		}
	}
}
