package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count      int
	lastUpdate time.Time
}

func (c *Counter) increment() {
	c.count++
	c.lastUpdate = time.Now()
}

func (c Counter) toString() string {
	return fmt.Sprintf("Count: %d Updated: %v", c.count, c.lastUpdate)
}

type IntTree struct {
	nodeVal           int
	leftVal, rightVal *IntTree
}

func (iTree *IntTree) insert(v int) *IntTree {
	if iTree == nil {
		return &IntTree{nodeVal: v}
	}

	if v < iTree.nodeVal {
		iTree.leftVal = iTree.leftVal.insert(v)
	}

	if v > iTree.nodeVal {
		iTree.rightVal = iTree.rightVal.insert(v)
	}

	return iTree
}

func (iTree *IntTree) Contains(v int) bool {
	switch {
	case iTree == nil:
		return false
	case v < iTree.nodeVal:
		return iTree.leftVal.Contains(v)
	case v > iTree.nodeVal:
		return iTree.rightVal.Contains(v)
	default:
		return true
	}
}

func main() {
	var c Counter
	fmt.Println("Before increment")
	fmt.Println(c.toString())
	c.increment()
	fmt.Println("After increment")
	fmt.Println(c.toString())

	var iTree *IntTree
	iTree = iTree.insert(5)
	iTree = iTree.insert(3)
	iTree = iTree.insert(7)
	iTree = iTree.insert(4)

	fmt.Println("*** Tree ***")
	fmt.Println("2 in tree", iTree.Contains(2))
	fmt.Println("4 in tree", iTree.Contains(4))

}
