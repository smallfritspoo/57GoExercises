package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

/*
Prompt a user for an input string
Display the string and number of character
string contains.
*/

func PromptUser(question string, reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Printf("%s", question)
	scanner.Scan()
	return scanner.Text()
}

func StringCLRFToLF(input string) string {
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	return input
}

func CountString(count_string string) int {
	return len(count_string)
}

func main() {
	user_string := StringCLRFToLF(PromptUser("What is the string to count? ", os.Stdin))
	fmt.Printf("String: %s\nCount: %v", user_string, CountString(user_string))
}
