package item

import "fmt"

type Item struct {
	Name            string
	SellIn, quality int
}

func NewItem(name string, sellIn int, quality int) *Item {
	return &Item{Name: name, SellIn: sellIn, quality: quality}
}

func (i *Item) GetQuality() int {
	return i.quality
}

func (i *Item) IncrementQuality() {
	i.quality++
}

func (i *Item) DecrementQuality() {
	i.quality--
}

func (i *Item) ZeroQuality() {
	i.quality = 0
}

type Items []*Item

func (i Items) String() string {
	var s string
	for _, item := range i {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
