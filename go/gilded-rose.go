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
		item.DecrementSellIn()

		if item.GetSellIn() < 0 {
			if item.GetQuality() < 50 {
				item.IncrementQuality()
			}
		}
	case BackstagePasses:
		if item.GetQuality() < 50 {
			item.IncrementQuality()
			if item.GetSellIn() < 11 {
				if item.GetQuality() < 50 {
					item.IncrementQuality()
				}
			}
			if item.GetSellIn() < 6 {
				if item.GetQuality() < 50 {
					item.IncrementQuality()
				}
			}
		}

		item.DecrementSellIn()

		if item.GetSellIn() < 0 {
			item.ZeroQuality()
		}
	case Sulfuras:
		break
	default:
		if item.GetQuality() > 0 {
			item.DecrementQuality()
		}

		item.DecrementSellIn()

		if item.GetSellIn() < 0 {
			if item.GetQuality() > 0 {
				item.DecrementQuality()
			}
		}
	}

}
