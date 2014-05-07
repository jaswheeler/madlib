package main

import "testing"
import "io"

func TestParseString(t *testing.T) {
	madlib := "this is ((some)) test ((a)) another"
	if extractedWords := InputMadLib(madlib); len(extractedWords) != 2 {
		t.Errorf("InputMadLib failed to extract words. Length was: %v", len(extractedWords))
	}
}

func TestRemoveParentheses(t *testing.T) {
	rawWord := "((test))"
	if word := StripParentheses(rawWord); word != "test" {
		t.Errorf("StripParenthese Failed.  Output was: %v", word)
	}
}

func TestReplacePlaceholders(t *testing.T) {
	input := "this is ((replace me)) with ((me too))"
	expected := "this is replace1 with replace2"
	words := map[string]string{"((replace me))": "replace1", "((me too))": "replace2"}
	if replaced := ReplaceAll(input, words); replaced != expected {
		t.Errorf("Replaced failed.  Output was: %v", replaced)
	}
}

type InputStub struct {
	value string
}

func (is InputStub) Read(p []byte) (n int, err error) {
	copy(p, []byte(is.value))
	return len(p), nil
}

func TestReadUserInput(t *testing.T) {
	input := "this is a test\r\n"
	is := InputStub{value: input }
	result := UserInput(io.Reader(is))
	if result != input[0:len(input) - 2] {
		t.Errorf("UserInput failed. Output: %v", result)
	}
}
