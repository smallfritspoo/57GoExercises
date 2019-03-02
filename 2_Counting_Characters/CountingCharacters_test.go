package main

import (
	"runtime"
	"strings"
	"testing"
)

const osenv string = runtime.GOOS

var question string = "string"
var win_input string = "a\r\n"
var lin_input string = "a\n"

func TestPromptUser(t *testing.T) {
	if osenv == "windows" {
		win_input = strings.TrimRight(win_input, "\r\n")
		reader := strings.NewReader(win_input)
		returned_data := PromptUser(question, reader)
		if returned_data != win_input {
			t.Error("Returned value doesn't match input")
		}
	} else {
		lin_input = strings.TrimRight(lin_input, "\n")
		reader := strings.NewReader(lin_input)
		returned_data := PromptUser(question, reader)
		if returned_data != lin_input {
			t.Error("Returned value doesn't match input")
		}
	}
}

func TestStringCLRFToLF(t *testing.T) {
	if osenv == "windows" {
		input := "Test_windows\r\n"
		output := "Test_windows"
		if StringCLRFToLF(input) != output {
			t.Error("Correct", osenv, " characters not stripped")
		}

	} else {
		input := "Test_linux\n"
		output := "Test_linux"
		if StringCLRFToLF(input) != output {
			t.Error("Correct", osenv, " characters not stripped")
		}
	}
}

func TestCountString(t *testing.T) {
	input := "five"
	output := 4
	if CountString(input) != output {
		t.Error("Characters counted incorrectly")
	}
}
