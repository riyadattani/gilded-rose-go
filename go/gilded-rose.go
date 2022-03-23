package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

func (i Item) lowQuality() bool {
	return i.quality < 50
}

func (i Item) highQuality() bool {
	return i.quality > 0
}

type Items []*Item

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfulars       = "Sulfuras, Hand of Ragnaros"
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		doUpdateQuality(item)

	}
}

func doUpdateQuality(item *Item) {
	switch item.name {
	case AgedBrie:
		if item.lowQuality() {
			item.quality = item.quality + 1
		}
		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			if item.lowQuality() {
				item.quality = item.quality + 1
			}
		}
	case BackstagePasses:
		if item.lowQuality() {
			item.quality = item.quality + 1
			if item.sellIn < 11 {
				if item.lowQuality() {
					item.quality = item.quality + 1
				}
			}
			if item.sellIn < 6 {
				if item.lowQuality() {
					item.quality = item.quality + 1
				}
			}
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			item.quality = 0
		}
	case Sulfulars:
		break
	default:
		if item.highQuality() {
			item.quality = item.quality - 1
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			if item.highQuality() {
				item.quality = item.quality - 1
			}
		}
	}
}

//
func (i Items) String() string {
	var s string
	for _, item := range i {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
