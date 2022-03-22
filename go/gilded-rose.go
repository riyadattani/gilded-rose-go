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
		if item.name != AgedBrie && item.name != BackstagePasses {
			if item.highQuality() {
				if item.name != Sulfulars {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.lowQuality() {
				item.quality = item.quality + 1
				if item.name == BackstagePasses {
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
			}
		}

		if item.name != Sulfulars {
			item.sellIn = item.sellIn - 1
		}

		if item.sellIn < 0 {
			if item.name != AgedBrie {
				if item.name != BackstagePasses {
					if item.highQuality() {
						if item.name != Sulfulars {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.lowQuality() {
					item.quality = item.quality + 1
				}
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
