package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	newFile, err := os.Create("./23_files/example1.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	newFile.WriteString("hi abhinav.")
	bytes := []byte(" how are you?")
	newFile.Write(bytes)

	f, err := os.Open("./23_files/example2.txt")
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

	data, err := os.ReadFile("./23_files/example1.txt") // not a viable solution for large file
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

	// read and write to another file (streaming fashion)
	sourceFile, err := os.Open("./23_files/example1.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	DestFile, err := os.Create("./23_files/example3.txt")
	if err != nil {
		panic(err)
	}
	defer DestFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(DestFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written To new file successfully")

	// Delete a file
	er := os.Remove("./23_files/test.txt")
	if er != nil {
		panic(er)
	}

	fmt.Println("file deleted successfully")
}
