package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

type Items []*Item

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfulars       = "Sulfuras, Hand of Ragnaros"
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].name != AgedBrie && items[i].name != BackstagePasses {
			if items[i].quality > 0 {
				if items[i].name != Sulfulars {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == BackstagePasses {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != Sulfulars {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != AgedBrie {
				if items[i].name != BackstagePasses {
					if items[i].quality > 0 {
						if items[i].name != Sulfulars {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
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
