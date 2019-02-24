package main

import (
	"runtime"
	"strings"
	"testing"
)

func TestPromptUser(t *testing.T) {
	input := "Test\r\n"
	question := "string"
	reader := strings.NewReader(input)
	returned_data := PromptUser(question, reader)
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	if returned_data != input {
		t.Error("Returned value doesn't match input")
	}
}

func TestStringCLRFToLF(t *testing.T) {
	input_windows := "Test_windows\r\n"
	output_windows := "Test_windows"
	input_linux := "Test_linux\n"
	output_linux := "Test_linux"
	if StringCLRFToLF(input_windows) != output_windows {
		t.Error("Correct windows characters not stripped")
	}
	if StringCLRFToLF(input_linux) != output_linux {
		t.Error("Correct Linux Characters not stripped")
	}
}

func TestCountString(t *testing.T) {
	input := "five"
	output := 4
	if CountString(input) != output {
		t.Error("Characters counted incorrectly")
	}
}
