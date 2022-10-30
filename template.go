package main

import (
	"html/template"
	"os"
)

type values struct {
}

func main() {

	file, err2 := os.ReadFile("x.txt")
	if err2 != nil {
		return
	}
	tmpl, err := template.New("test").Parse(string(file))
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"httpLibrary": "gin",
	}
	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
