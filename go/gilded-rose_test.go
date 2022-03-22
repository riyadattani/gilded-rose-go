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

	UpdateQuality(items)

	approvals.VerifyString(t, items.String())
}
