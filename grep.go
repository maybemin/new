package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 2 {
		ioscanner := bufio.NewScanner(os.Stdin)

		for ioscanner.Scan() {
			if strings.Contains(ioscanner.Text(), os.Args[1]) {
				fmt.Println(ioscanner.Text())
			}
		}
	}

	if len(os.Args) == 3 {
		file, _ := os.Open(os.Args[2])
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), os.Args[1]) {
				fmt.Println(scanner.Text())
			}
		}
	}

}
