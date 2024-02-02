package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" > ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("echoing: ", text)
}
