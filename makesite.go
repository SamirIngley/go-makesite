package main

import (
	"flag"
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

func activeFlag(name string) bool {
	active := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			active = true
		}
	})

	return active
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
	// saveToFile("first-post.html", "first-post.txt")

	examplePtr := flag.String("file", "first-post", " Help text.")
	flag.Parse()
	dirPtr := flag.String("dir", ".", "pull files from")
	flag.Parse()

	saveToFile(*examplePtr+".html", *examplePtr+".txt")

	if activeFlag("dir") {
		files := readFile(*dirPtr)

		for _, file := range files {

		}

	} else if activeFlag("md") {

		saveToFile(template, fileName)

	}

	content := readFile(*filePtr)
	template := renderTemplate(content)
	saveToFile(template, fileName)

}
