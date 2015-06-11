package main

import (
	"bufio"
	"fmt"
	"github.com/hni/gsafeed/lib"
	"os"
)

func main() {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
    xml := gsafeed.GsaUrlFeed(urls, os.Args[1])
    fmt.Print(xml)
}
