package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
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

func getHosts(file *os.File) chan Host {
	c := make(chan Host)
	scanner := bufio.NewScanner(file)

	go func() {
		defer close(c)
		for scanner.Scan() {
			c <- Host{Url: scanner.Text()}
		}
	}()

	return c
}

func digester(hosts <-chan Host, results chan<- *Footprint) {
	for host := range hosts {
		if footprint, err := host.Get(); err != nil {
			log.Println(err)
		} else {
			results <- footprint
		}
	}
}

func main() {
	file := openFile(os.Args)
	defer file.Close()

	const digesterCount = 1024
	var wg sync.WaitGroup
	wg.Add(digesterCount)

	hosts := getHosts(file)
	footprints := make(chan *Footprint)

	for i := 0; i < digesterCount; i++ {
		go func() {
			digester(hosts, footprints)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(footprints)
	}()

	for f := range footprints {
		fmt.Println(f)
	}
}
