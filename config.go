package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var dirs = []string{".", "partials"}
var extension = ".hbs"

var root string

var translation = map[string]string{}

var inputReader *bufio.Reader

func init() {
	inputReader = bufio.NewReader(os.Stdin)
	flag.StringVar(&root, "path", ".", "Path to the Ghost theme to translate")
}

func readString() string {
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	return strings.TrimSpace(input)
}
