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
	if p == nil {
		return Parent{}
	}

	cp := *p

	cp.Children = make([]Child, len(p.Children))
	copy(cp.Children, p.Children)

	return cp
}

func main() {
	cp := CopyParent(nil)

	fmt.Printf("nil case: %+v \n", cp)

	p := &Parent{Name: "123"}
	cp2 := CopyParent(p)
	cp2.Name = "321"

	fmt.Printf("filled case p: %+v \n", p)
	fmt.Printf("filled case cp: %+v \n", cp2)
}
