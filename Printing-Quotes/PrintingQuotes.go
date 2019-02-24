package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var author string
var quote string

//DRY user collection
func CollectInput(query string, reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Printf("%s", query)
	scanner.Scan()
	return scanner.Text()
}

func FormatQuote(quote, author string) string {
	return author + " says, \"" + quote + "\""
}

func main() {
	quote = CollectInput("What is the Quote? ", os.Stdin)
	author = CollectInput("Who Said it? ", os.Stdin)
	fmt.Println(FormatQuote(quote, author))
}
