package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var author string
var quote string

/*
DRY user collection.
function paramaterized for io reader for testing purposes.
*/
func CollectInput(query string, reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Printf("%s", query)
	scanner.Scan()
	return scanner.Text()
}

/*
Constraints:
  - Use a single output statement
  - Use string concatenation instead of interpolation
*/
func FormatQuote(quote, author string) string {
	return author + " says, \"" + quote + "\""
}

func main() {
	quote = CollectInput("What is the Quote? ", os.Stdin)
	author = CollectInput("Who Said it? ", os.Stdin)
	fmt.Println(FormatQuote(quote, author))
}
