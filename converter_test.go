package main

import (
	"fmt"
	"testing"
)

func TestGenericStoneConversion(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		s string
		v int64
		r float64
	}{
		{1, 1, "a", 1, 0},
		{10, 2, "b", 5, 0},
		{345, 6.35029318, "c", 54, 4.59},
		{345, 14, "d", 24, 9},
		{345, 6350.29318, "e", 0, 0.76},
		{345, 224, "f", 1, 7.56},
	}

	for _, table := range tables {
		value, remainder := genericStoneConversion(table.x, table.y, table.s)
		if value != table.v && remainder != table.r {
			t.Errorf(
				"Value of (%v/%v) did not match %v stones and remainder %v actually got %v %v", table.x, table.y, table.v, table.r, value, remainder,
			)
		}
	}

	value, remainder := genericStoneConversion(5, 5, "string")
	fmt.Println(value)
	fmt.Println(remainder)
}
