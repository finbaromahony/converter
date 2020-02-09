package main

import (
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
}

func TestGenericConversion(t *testing.T) {
	tables := []struct {
		value          float64
		rate           float64
		rate_name      string
		float64_result float64
	}{
		{500, ratesMap["stoneToLbs"], "stoneToLbs", 7000},
		{500, ratesMap["stoneToKg"], "stoneToKg", 3175.15},
		{500, ratesMap["stoneToGrams"], "stoneToGrams", 3.175e+6},
		{500, ratesMap["stoneToOz"], "stoneToOz", 112000},
		{500, ratesMap["lbsToKg"], "lbsToKg", 226.796},
		{500, ratesMap["lbsToGrams"], "lbsToGrams", 226796},
		{500, ratesMap["lbsToOz"], "lbsToOz", 8000},
		{500, ratesMap["kgToLbs"], "kgToLbs", 1102.31},
		{500, ratesMap["kgToGrams"], "kgToGrams", 500000},
		{500, ratesMap["kgToOz"], "kgToOz", 17637},
		{500, ratesMap["gramsToLbs"], "gramsToLbs", 1.10231},
		{500, ratesMap["gramsToKg"], "gramsToKg", 0.5},
		{500, ratesMap["gramsToOz"], "gramsToOz", 17.637},
		{500, ratesMap["ozToLbs"], "ozToLbs", 31.25},
		{500, ratesMap["ozToKg"], "ozToKg", 14.1748},
		{500, ratesMap["ozToGrams"], "ozToGrams", 14174.8},
	}
	for _, table := range tables {
		result := genericUnitConversion(table.value, table.rate, table.rate_name)
		if result != table.value {
			t.Errorf(
				"%s Value of (%v/%v) did not match %v actually got %v", table.rate_name, table.value, table.rate, table.float64_result, result,
			)
		}
	}
}
