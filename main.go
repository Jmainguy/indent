package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
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

	// Indent
	indention := strings.Repeat(" ", int(math.Abs(float64(*indentPtr))))
	if *indentPtr < 0 {
		for _, line := range text {
			indent := regexp.MustCompile(fmt.Sprintf(`^%s`, indention))
			newLineArray := indent.Split(line, 2)
			if len(newLineArray) > 1 {
				fmt.Printf("%s\n", newLineArray[1])
			} else {
				fmt.Printf("%s\n", line)
			}
		}

	} else {
		for _, line := range text {
			fmt.Printf("%s%s\n", indention, line)
		}
	}
}
