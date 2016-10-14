package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fi, err := file.Stat() //파일정보
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())

	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	file2, _ := os.Create(os.Args[2])
	defer file2.Close()
	_, err = file2.Write([]byte(string(data)))

	fmt.Println(string(data))
}
