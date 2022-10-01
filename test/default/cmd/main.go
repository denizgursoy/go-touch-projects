package main

import "fmt"

func main() {
	fmt.Println("{{.Port}}")

	{{ range .Array }}
		{{ .Field1 }}
		{{ .Field2 }}
	{{ end }}
}
