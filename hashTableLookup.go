package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
		}
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces: ", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)

	fmt.Println()
	if lookup(hash, 120) {
		fmt.Printf("%d is not found\n", 120)
	} else {
		fmt.Printf("%d is found\n", 120)
	}

	if lookup(hash, 121) {
		fmt.Printf("%d is not found\n", 121)
	} else {
		fmt.Printf("%d is found\n", 121)
	}

	if lookup(hash, 99) {
		fmt.Printf("%d is not found\n", 99)
	} else {
		fmt.Printf("%d is found\n", 99)
	}

	if lookup(hash, 11) {
		fmt.Printf("%d is not found\n", 11)
	} else {
		fmt.Printf("%d is found\n", 11)
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}
