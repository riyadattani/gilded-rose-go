package item

import "fmt"

type Item struct {
	Name            string
	SellIn, Quality int
}

func NewItem(name string, sellIn int, quality int) *Item {
	return &Item{Name: name, SellIn: sellIn, Quality: quality}
}

type Items []*Item

func (i Items) String() string {
	var s string
	for _, item := range i {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
