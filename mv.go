package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
}
