package ast_demo

import "fmt"

type Example1 struct {
	// Foo Comments
	Foo string `json:"foo"`
}

type Example2 struct {
	// Aoo Comments
	Aoo int `json:"aoo"`
}

// print Hello World
func PrintHello(){
	fmt.Println("Hello World")
}