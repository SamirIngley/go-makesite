package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type post struct{ Content string }

// reads data from a file -> returns data as a string
func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

// renderTemplate reads file data, creates template -> executes template to STDOUT with file data
func renderTemplate(filename string) {

	cont := post{Content: readFile(filename)}
	tmp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	var err error
	err = tmp.Execute(os.Stdout, cont)
	if err != nil {
		panic(err)
	}
}

// saveToFile creates a file, saves readThis to FILENAME
func saveToFile(filename string, readThis string) {

	file, err := os.Create(filename)
	cont := post{Content: readFile(readThis)}

	tmp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	err = tmp.Execute(file, cont)
	if err != nil {
		panic(err)
	}
}

// walks the current file path and checks for ext parameter
func checkIfTxt(ext string) []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	extList := []string{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(path) == ext {
			extList = append(extList, fileInfo.Name())
		}
		return nil
	})

	fmt.Println(extList)
	return extList
}

func main() {

	// fmt.Printf(readFile("first-post.txt"))
	// renderTemplate("first-post.txt")
	// saveToFile("first-post.html", "first-post.txt")

	// examplePtr := flag.String("file", "first-post", " Help text.")
	// dirPtr := flag.String("file", "first-post", " Help text.")

	flag.Parse()

	checkIfTxt(".txt")

	// if examplePtr
	// creates html file from ptr name, and data from ptr text
	// saveToFile(*examplePtr+".html", *examplePtr+".txt")
	// else
	// readFromFile(*dirPtr)
	// ioutil.readDir https://golang.org/pkg/io/ioutil/
}
