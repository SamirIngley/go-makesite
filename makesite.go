package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type post struct{ Content string }

// readFile reads the data from a file
func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func renderTemplate(filename string) {

	cont := post{Content: readFile(filename)}
	tmp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	var err error
	err = tmp.Execute(os.Stdout, cont)
	if err != nil {
		panic(err)
	}
}

func saveToFile(filename string, readThis string) {

	file, err := os.Create(filename)
	cont := post{Content: readFile(readThis)}

	tmp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	err = tmp.Execute(file, cont)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf(readFile("first-post.txt"))
	renderTemplate("first-post.txt")
	saveToFile("first-post.html", "first-post.txt")
}
