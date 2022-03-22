package main

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func Test_Foo(t *testing.T) {
	names := []string{"Aged Brie", "Backstage passes to a TAFKAL80ETC concert"}
	sellins := []int{0, 1}
	quality := []int{0, 1}

	approvals.VerifyAllCombinationsFor3(t, "update quality", func(name, sellin, quality interface{}) string {
		item := &Item{name.(string), sellin.(int), quality.(int)}
		items := Items{item}
		UpdateQuality(items)
		return items.String()
	}, names, sellins, quality)
}
