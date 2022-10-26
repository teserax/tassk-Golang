package main

import "fmt"

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	*p = Parent{}

	return *p
}
func main() {

	// cp := CopyParent(nil) // Parent{}

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}
	cp := CopyParent(p)

	fmt.Printf("%v", cp)
}
