package main

import (
	"fmt"
	"os"
)

func main() {

	newFile, err := os.Create("./23_files/example.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	newFile.WriteString("hi abhinav.")
	bytes := []byte(" how are you?")
	newFile.Write(bytes)

	f, err := os.Open("./23_files/example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileinfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("file name: ", fileinfo.Name())
	fmt.Println("file or folder: ", fileinfo.IsDir())
	fmt.Println("file size: ", fileinfo.Size())
	fmt.Println("file permission: ", fileinfo.Mode())
	fmt.Println("file modified at: ", fileinfo.ModTime())
	fmt.Println("file modified at: ", fileinfo.ModTime())

	defer f.Close()

	buf := make([]byte, 12)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	println("data", buf, d)

	for i := 0; i < len(buf); i++ {
		fmt.Print(string(buf[i]))
	}

	fmt.Println()

	data, err := os.ReadFile("./23_files/example.txt") // not a viable solution for large file
	if err != nil {
		panic(err)
	}
	fmt.Println("file string data: ", string(data))

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	dirFileinfo, err := dir.ReadDir(2) // -1 for all the folder and files
	if err != nil {
		panic(err)
	}

	for _, fi := range dirFileinfo {
		fmt.Println("file name: ", fi.Name(), "directory name: ", fi.IsDir())
	}

}
