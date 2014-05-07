package main

import "fmt"
import "strings"
import "regexp"
import "bufio"
import "os"
import "io"

func InputMadLib(madlib string) []string {
	re := regexp.MustCompile("\\(\\(.*?\\)\\)")
	words := re.FindAllString(madlib, -1)
	return words
}

func StripParentheses(word string) string {
	tail := len(word) - 2
	return word[2:tail]
}

func ReplaceAll(madlib string, words map[string]string) string {
	for k, v := range words {
		madlib = strings.Replace(madlib, k, v, -1)
	}
	return madlib
}

func UserInput(source io.Reader) string {
	var input string
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		input = scanner.Text()
		break
	}
	return input
}

func main() {
	fmt.Print("Input a MadLib: ")
	madlib := UserInput(os.Stdin)
	wordList := InputMadLib(madlib)
	if len(wordList) == 0 {
		fmt.Println("The MadLib should contain placeholders wrapped in double parentheses. Example:\n 'I like to ((verb))'!")
		return
	}

	wordMap := make(map[string]string)
	for i := range wordList {
		fmt.Printf("Enter a %v: ", StripParentheses(wordList[i]))
		wordMap[wordList[i]] = UserInput(os.Stdin)
	}

	fmt.Println(ReplaceAll(madlib, wordMap))
	return
}
