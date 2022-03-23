package item

import "fmt"

type Item struct {
	Name            string
	SellIn, Quality int
}

func NewItem(name string, sellIn int, quality int) *Item {
	return &Item{Name: name, SellIn: sellIn, Quality: quality}
}

func (i *Item) GetQuality() int {
	return i.Quality
}

func (i *Item) IncrementQuality() {
	i.Quality++
}

func (i *Item) DecrementQuality() {
	i.Quality--
}

func (i *Item) ZeroQuality() {
	i.Quality = 0
}

type Items []*Item

func (i Items) String() string {
	var s string
	for _, item := range i {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
