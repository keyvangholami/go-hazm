package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var dataPath string = filepath.Join(filepath.Dir(os.Args[0]), "data")
var defaultWords string = filepath.Join(dataPath, "words.dat")
var defaultStopWords string = filepath.Join(dataPath, "stop_words.dat")
var defaultVerbs string = filepath.Join(dataPath, "verbs.dat")
var informalWords string = filepath.Join(dataPath, "informal_words.dat")
var informalVerbs string = filepath.Join(dataPath, "informal_verbs.dat")

var numbers string = "۰۱۲۳۴۵۶۷۸۹"

func MakeTrans(A, B string) map[rune]rune {
	trans := make(map[rune]rune)
	aRunes := []rune(A)
	bRunes := []rune(B)
	for i, aRune := range aRunes {
		trans[aRune] = bRunes[i]
	}
	return trans
}

// WordsList returns a list of words.
//
// examples:
//
//	words, err := WordsList("")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(words[1]) // Output: ('آب', 549005877, ('N', 'AJ')) #(id, word, (tag1, tag2, ...))
//
// Args:
//
//	words_file (string, optional): The path of the file containing the words.
//
// Returns:
//
//	(tuple[string,string,tuple[string,string]]): A list of words.
func WordsList(wordsFile string) []interface{} {
	if wordsFile == "" {
		wordsFile = defaultWords
	}
	file, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []interface{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		if len(fields) == 3 {
			word := fields[0]
			count, _ := strconv.Atoi(fields[1])
			tags := strings.Split(fields[2], ",")
			items = append(items, []interface{}{word, count, tags})
		}
	}
	return items
}

// StopWordsList Returns a list of stop words.
//
// Examples:
//
//	stopWords, err := StopWordsList("")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(stopWords[:4]) // Output: ['محسوب', 'اول', 'بسیار', 'طول']
//
// Args:
//
//	stopWords_file (string, optional): The path of the file containing the stop words.
//
// Returns:
//
//	(list[string]): A list of stop words.
func StopWordsList(stopWordsFile string) []string {
	if stopWordsFile == "" {
		stopWordsFile = defaultStopWords
	}
	file, err := os.Open(stopWordsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []string
	for scanner.Scan() {
		word := scanner.Text()
		items = append(items, word)
	}
	return items
}

func verbsList() []string {
	file, err := os.Open(defaultVerbs)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []string
	for scanner.Scan() {
		word := scanner.Text()
		items = append(items, word)
	}
	return items
}

func PastRoots() string {
	roots := ""
	for _, verb := range verbsList() {
		split := strings.Split(verb, "#")
		roots += split[0] + "|"
	}
	return roots[:len(roots)-1]
}

func PresentRoots() string {
	roots := ""
	for _, verb := range verbsList() {
		split := strings.Split(verb, "#")
		roots += split[1] + "|"
	}
	return roots[:len(roots)-1]
}

func RegexReplace(patterns [][2]string, text string) string {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern[0])
		text = re.ReplaceAllString(text, pattern[1])
	}
	return text
}
