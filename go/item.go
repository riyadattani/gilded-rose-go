package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

func NewItem(name string, sellIn int, quality int) *Item {
	return &Item{name: name, sellIn: sellIn, quality: quality}
}

type Items []*Item

func (i Items) String() string {
	var s string
	for _, item := range i {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
