package main

import (
	"fmt"
	"math"
)

//generic converter to stone that accepts conversion rate
func genericStoneConversion(value float64, rate float64, rate_name string) (int64, float64) {
	stone := value / rate
	fmt.Printf("%v%s = %v Stone\n", value, rate_name, stone)
	// get value for stone that doesnt include any lbs
	baseStone := math.Floor(stone) * 14
	// get the number of remaining lbs to 2 decimal places
	remainder := math.Floor(((stone*14)-baseStone)*100) / 100
	fmt.Println("or")
	fmt.Printf("%v Stone; %v lbs\n", int64(stone), remainder)
	return int64(stone), remainder
}

//lbsToStone converts lbs to stone.
func lbsToStone(lbs float64) {
	fmt.Println("- Convert lbs to stone -")
	// 14 pounds in a stone
	genericStoneConversion(lbs, 14, "lbs")
}

//kgToStone converts kg to stone
func kgToStone(kg float64) {
	fmt.Println("- Convert kg to stone -")
	// 6.35029318 kg in a stone
	genericStoneConversion(kg, 6.35029318, "kg")
}

//gramsToStone converts grams to stone
func gramsToStone(grams float64) {
	fmt.Println("- Convert grams to stone -")
	// 6350.29318 grams in a stone
	genericStoneConversion(grams, 6350.29318, "grams")
}

//ozToStone converts oz to stone
func ozToStone(oz float64) {
	fmt.Println("- Convert oz to stone -")
	// 224 oz in a stone
	genericStoneConversion(oz, 224, "oz")
}

//toLbs section

//toOz section

//toGrams section
