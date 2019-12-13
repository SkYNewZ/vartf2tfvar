package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	// setup flags
	var fileArgs string
	flag.Usage = usage
	flag.StringVar(&fileArgs, "file", "variables.tf", "Input Terraform variables files")
	flag.Parse()

	if filepath.Ext(strings.TrimSpace(fileArgs)) != ".tf" {
		fmt.Println("Invalid .tf file")
		os.Exit(1)
	}

	// Open file
	bytes, err := ioutil.ReadFile(fileArgs)

	// if error while openning
	if err != nil {
		usage()
		fmt.Println(err)
		os.Exit(1)
	}

	// get content as string
	content := string(bytes)

	re, _ := regexp.Compile(`(?m)variable "([a-zA-Z_-]+)"`)
	for _, match := range re.FindAllStringSubmatch(content, -1) {
		fmt.Println(match[1] + " = \"\"\n")
	}
}
