package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, _ := os.Open("hello.txt")
	defer file.Close()

	buf := make([]byte, 1024)
	count, _ := file.Read(buf)
	fmt.Printf("file buf size is %d\n", count)

	stat, _ := file.Stat()
	curr_data_size := stat.Size()

	for {
		stat, _ := file.Stat()
		after_data_size := stat.Size()

		if curr_data_size < after_data_size {

			f, _ := os.Open("hello.txt")
			f.Seek(curr_data_size, os.SEEK_SET)
			bufsize := make([]byte, after_data_size-curr_data_size)
			io.ReadAtLeast(f, bufsize, 2)
			fmt.Printf("%s\n", string(bufsize))

			curr_data_size = after_data_size

		}
	}

}
