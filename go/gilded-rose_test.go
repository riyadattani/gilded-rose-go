package main

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func Test_Foo(t *testing.T) {
	var items = Items{
		&Item{"foo", 0, 0},
		&Item{"bar", 2, 1},
	}

	approvals.VerifyAll(t, "update quality", items, func(item interface{}) string {
		i := item.(*Item)
		items := Items{i}
		UpdateQuality(items)
		return items.String()
	})

	//approvals.VerifyString(t, items.String())
}
