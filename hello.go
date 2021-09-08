package main

import (
    "fmt"
)

type  Person struct {
    name string
    address string
    age int
}

func modify(age int, name string, p Person) {
    age = age * 2
    name = "new name"
    p.age = age
    p.name = name
    p.address = "yeapie"

}

func modMap(m map[int] string) {
    m[2] = "hello"
    m[3] = "goodbye!"
    delete(m, 1)
}

func modSlice(s [] int) {
    for k, v := range s {
        fmt.Println(k)
        s[k] = v*2
    }

    s = append(s, 10)
}

func main() {
    age := 10
    name := "sup shady"
    p := Person{
        "Lukman",
        "muchener str",
        20,
    }

    modify(age, name, p)

    fmt.Println(age, name, p)
    fmt.Println("********************************")

    m := map[int]string {
        1: p.name,
        2: p.address,
    }

    fmt.Print(" Before modification: ")
    fmt.Println(m)

    modMap(m)
    fmt.Print(" After modification: ")
    fmt.Println(m)
    fmt.Println("********************************")


    s := [] int {1,2,3}
    s = append(s, 4)

    fmt.Print(" Before modification: ")
    fmt.Println(s)

    modSlice(s)
    fmt.Print(" After modification: ")
    fmt.Println(s)
}