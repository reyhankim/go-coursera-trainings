package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	findIan(line)
}

func findIan(input string) {
	re := regexp.MustCompile(`(?s).*i.*a.*n.*`)
	isIanFound := re.Match([]byte(input))
	if (isIanFound) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}