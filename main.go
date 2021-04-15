package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	filenamePtr := flag.String("filename", "", "Filename to read")
	indentPtr := flag.Int("indent", 0, "Amount of spaces to indent")

	flag.Parse()

	file, err := os.Open(*filenamePtr)

	if err != nil {
		fmt.Println("failed to open")
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	indention := strings.Repeat(" ", *indentPtr)
	for _, line := range text {
		fmt.Printf("%s%s\n", indention, line)
	}
}
