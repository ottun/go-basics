package main

import (
	"fmt"
)

type person struct {
	FirstName string
	MiddleName *string
	LastName string
}

func failedUpdate(v *int) {
	x :=20
	v = &x
}

func update(v *int) {
	*v = 30
}

func main() {
	x :=10
	xref := &x

	fmt.Println(x)
	fmt.Println(xref)
	fmt.Println(*xref)
//-----------------------------
mn := "O."
p := person{"lukman", &mn, " Ottun"}

fmt.Println(*p.MiddleName)

fmt.Println("********************************")

failedUpdate(xref)
fmt.Println(x)

update(xref)
fmt.Println(x)

}

