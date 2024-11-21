package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() int
}

type IntSliceIterator struct {
	collection []int
	index      int
}

func (it *IntSliceIterator) HasNext() bool {
	return it.index < len(it.collection)
}

func (it *IntSliceIterator) Next() int {
	if !it.HasNext() {
		panic("No more elements in the iterator")
	}
	element := it.collection[it.index]
	it.index++
	return element
}

type IntCollection struct {
	items []int
}

func NewIntCollection(items []int) *IntCollection {
	return &IntCollection{items: items}
}

func (c *IntCollection) CreateIterator() Iterator {
	return &IntSliceIterator{
		collection: c.items,
		index:      0,
	}
}

func main() {
	collection := NewIntCollection([]int{1, 2, 3, 4, 5})

	iterator := collection.CreateIterator()

	fmt.Println("Iterating through the collection:")
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
