package main

import (
	"html/template"
	"os"
)

func main() {

	file, err2 := os.ReadFile("medium-pom.xml")
	if err2 != nil {
		return
	}
	tmpl, err := template.New("test").Parse(string(file))
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"Dependencies": []interface{}{
			map[string]string{
				"groupId":    "org.postgresql",
				"artifactId": "postgresql",
				"version":    "42.5.0",
			},
		},
	}
	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
