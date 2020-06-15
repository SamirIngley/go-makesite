package main

import (
	"io/ioutil"
	"fmt"
)


func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(filename string, data string) {
	cont := post{Content: data}
	tmp := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , cont)
	if err != nil {
		panic(err)
	}
}


func main() {
	fmt.Println("Hello, world!")
}
