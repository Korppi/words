package words

import (
	"bufio"
	"embed"
	"log"
	"strings"
)

//go:embed words_alpha.txt
var content embed.FS

// Reverse given string s
func Reverse(s string) (reversed string) {
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return
}

// Check if word is case-insensitive palindrome.
// Example: ABba returns true
func IsPalindrome(s string) bool {
	lowercaseString := strings.ToLower(s)
	return IsPalindromeCaseSensitive(lowercaseString)
}

// Check if word is case-sensitive palindrome.
// Example: ABba returns false
func IsPalindromeCaseSensitive(s string) bool {
	reversed := Reverse(s)
	return s == reversed
}

// Check if s1 and s2 are anagrams (both have same letters in different order)
// Example of anagrams, words "rat" and "tar" are anagrams
func IsAnagram(s1 string, s2 string) bool {
	// Remove empty spaces and make all lowercase
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)

	for _, s1Rune := range s1 {
		s1Map[s1Rune]++
	}

	for _, s2Rune := range s2 {
		s2Map[s2Rune]++
	}

	for letter, s1Count := range s1Map {
		s2Count, found := s2Map[letter]
		if !found || s1Count != s2Count {
			return false
		}
	}
	return true
}

// Return word length
func WordLength(s string) (count int) {
	for range s {
		count++
	}
	return
}

// Return how many times substring occurs in s
func CountOccurences(s string, substring string) int {
	return 0
}

// Returns true if world contains every letter in alphabets atleast once
func IsPangram(s string) bool {
	var alphabetCounter [26]int
	s = strings.ToLower(s)
	for i := 0; i < WordLength(s); i++ {
		alphabetCounter[s[i]-'a']++
	}

	for _, v := range alphabetCounter {
		if v == 0 {
			return false
		}
	}
	return true
}

// Return true if r is vowel
func IsVowel(r rune) bool {
	vowels := "AEIOUaeiou"
	return strings.ContainsRune(vowels, r)
}

// Return true if r is consonant
func IsConsonant(r rune) bool {
	consonants := "BCDFGHJKLMNPQRSTVXYZbcdfghjklmnpqrstvxyz"
	return strings.ContainsRune(consonants, r)
}

// Returns count of consonants in given string s
func CountConsonants(s string) (count int) {
	for _, r := range s {
		if IsConsonant(r) {
			count++
		}
	}
	return
}

// Returns count of vowels in given string s
func CountVowels(s string) (count int) {
	for _, r := range s {
		if IsVowel(r) {
			count++
		}
	}
	return
}

func SearchPossibleWords(letters string) (result []string) {

	file, err := content.Open("words_alpha.txt") //os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	letterLen := len(letters)
	letters = strings.ToLower(letters)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		originalText := text
		if len(text) <= letterLen {
			for _, letter := range letters {
				text = strings.Replace(text, string(letter), "", 1)
			}
			if text == "" {
				result = append(result, originalText)
			}
		}
	}
	return
}
