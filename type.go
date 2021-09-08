package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
	lastUpdate time.Time
}

func (c *Counter) increment() {
	c.count++
	c.lastUpdate = time.Now()
}

func (c Counter) toString() string {
	return fmt.Sprintf("Count: %d Updated: %v", c.count, c.lastUpdate)
}

func main() {
	var c Counter
	fmt.Println("Before increment")
	fmt.Println(c.toString())
	c.increment()
	fmt.Println("After increment")
	fmt.Println(c.toString())
}
