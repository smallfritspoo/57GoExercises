package main

import (
	"io"
	"strings"
	"testing"
)

var question string = "query string"
var user_input string = "user input string"
var test_quote string = "quote string"
var test_author string = "author string"
var test_collected_data string
var test_formatted_data string

func User_Return(input string) io.Reader {
	return strings.NewReader(input)
}

func TestCollectInput(t *testing.T) {
	test_collected_data = CollectInput(question, User_Return(user_input))
	if test_collected_data != user_input {
		t.Error("User input string interpreted incorrectly")
	}
}

func TestFormatQuote(t *testing.T) {
	test_formatted_data = FormatQuote(test_quote, test_author)
	if test_formatted_data != "author string says, \"quote string\"" {
		t.Error("quote string formatted incorrectly")
	}
}
