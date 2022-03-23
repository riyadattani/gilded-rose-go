package main

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		doUpdateQuality(item)

	}
}

func doUpdateQuality(item *Item) {
	switch item.name {
	case AgedBrie:
		if item.quality < 50 {
			item.quality = item.quality + 1
		}
		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			if item.quality < 50 {
				item.quality = item.quality + 1
			}
		}
	case BackstagePasses:
		if item.quality < 50 {
			item.quality = item.quality + 1
			if item.sellIn < 11 {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
			if item.sellIn < 6 {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			item.quality = 0
		}
	case Sulfuras:
		break
	default:
		if item.quality > 0 {
			item.quality = item.quality - 1
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			if item.quality > 0 {
				item.quality = item.quality - 1
			}
		}
	}
}
