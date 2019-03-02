package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringCRLFToLF(name string) string {
	return strings.Replace(name, "\r\n", "", -1)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello what is your name?")

	text, _ := reader.ReadString('\n')

	fmt.Printf("Hello, %s, nice to meet you!\n", StringCRLFToLF(text))

}
