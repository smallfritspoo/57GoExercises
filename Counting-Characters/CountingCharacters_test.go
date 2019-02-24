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
	osenv := runtime.GOOS

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
