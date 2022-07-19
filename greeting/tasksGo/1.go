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
	copy := *p

	return copy
}

func main() {

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
			{
				Name: "Vasya",
				Age:  22,
			},
		},
	}
	cp := CopyParent(p)
	cp.Children[0] = Child{}
	fmt.Println(&cp)
	fmt.Println(&p)
}
