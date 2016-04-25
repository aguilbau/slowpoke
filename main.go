package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFile(args []string) *os.File {
	if len(args) == 1 {
		return os.Stdin
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func readLines(file *os.File) chan string {
	c := make(chan string)
	scanner := bufio.NewScanner(file)

	go func() {
		defer close(c)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()

	return c
}

func main() {
	file := openFile(os.Args)
	defer file.Close()
	c := readLines(file)

	for l := range c {
		fmt.Println(l)
	}
}
