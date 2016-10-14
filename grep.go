package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) == 2 {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if len(os.Args) == 3 {
		file, _ := os.Open(os.Args[2])
		defer file.Close()
		scanner = bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
	}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), os.Args[1]) {
			fmt.Println(scanner.Text())
		}
	}

}
