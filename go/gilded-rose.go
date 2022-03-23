package main

import (
	"github.com/riyadattani/item"
)

type GildedRose struct {
	items item.Items
}

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
)

func NewGildedRose(items item.Items) *GildedRose {
	return &GildedRose{items: items}
}

func (g GildedRose) UpdateQuality() {
	for _, item := range g.items {
		g.doUpdateQuality(item)

	}
}

func (g GildedRose) doUpdateQuality(item *item.Item) {
	switch item.Name {
	case AgedBrie:
		if item.GetQuality() < 50 {
			item.IncrementQuality()
		}
		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 {
			if item.GetQuality() < 50 {
				item.IncrementQuality()
			}
		}
	case BackstagePasses:
		if item.GetQuality() < 50 {
			item.IncrementQuality()
			if item.SellIn < 11 {
				if item.GetQuality() < 50 {
					item.IncrementQuality()
				}
			}
			if item.SellIn < 6 {
				if item.GetQuality() < 50 {
					item.IncrementQuality()
				}
			}
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 {
			item.ZeroQuality()
		}
	case Sulfuras:
		break
	default:
		if item.GetQuality() > 0 {
			item.DecrementQuality()
		}

		item.SellIn = item.SellIn - 1

		if item.SellIn < 0 {
			if item.GetQuality() > 0 {
				item.DecrementQuality()
			}
		}
	}

}
