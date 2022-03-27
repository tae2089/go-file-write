package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b := []byte("Data to write5\n")

	f, err := os.Create("test1.txt")
	check(err)
	w := bufio.NewWriter(f)
	n, err := w.WriteString(string(b))
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	err = ioutil.WriteFile("test2.txt", b, 0644)
	check(err)
	//f, err := os.OpenFile("./hello2.txt", os.O_RDWR, 0755)
	//f.Seek(0, 2)
	//defer f.Close()
	//check(err)
	//_, err = f.Write(b)
	//check(err)
}
