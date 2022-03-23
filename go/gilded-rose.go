package main

import (
	"fmt"
)

type GildedRoseItem interface {
	GetQuality() int
	GetSellIn() int
	GetName() string

	ZeroQuality()
	IncrementQuality()
	DecrementQuality()
	DecrementSellIn()
}

type GildedRose struct {
	items GildedRoseItems
}

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
)

func NewGildedRose(items GildedRoseItems) *GildedRose {
	return &GildedRose{items: items}
}

func (g GildedRose) UpdateQuality() {
	for _, item := range g.items {
		g.doUpdateQuality(item)

	}
}

func (g GildedRose) doUpdateQuality(item GildedRoseItem) {
	switch item.GetName() {
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

type GildedRoseItems []GildedRoseItem

func (g GildedRoseItems) String() string {
	var s string
	for _, item := range g {
		s += fmt.Sprintf("%#v\n", item)
	}

	return s
}
