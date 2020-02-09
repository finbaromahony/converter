package main

import (
	"fmt"
	"math"
)

var ratesMap = map[string]float64{
	"stoneToLbs":   0.071428571,
	"stoneToKg":    0.157473044,
	"stoneToGrams": 0.000157473,
	"stoneToOz":    0.004545455,
	"lbsToStone":   14,
	"lbsToKg":      2.20462262,
	"lbsToGrams":   0.002204623,
	"lbsToOz":      16,
	"kgToStone":    6.35029318,
	"kgToLbs":      0.45359237,
	"kgToGrams":    0.001,
	"kgToOz":       0.028349523,
	"gramsToStone": 6350.29318,
	"gramsToLbs":   453.59237,
	"gramsToKg":    1000,
	"gramsToOz":    28.3495231,
	"ozToStone":    220,
	"ozToLbs":      0.0625,
	"ozToKg":       35.2739619,
	"ozToGrams":    0.035273962,
}

//return float that has 2 decimals max
func toTwoDecimals(value float64) float64 {
	return math.Floor((value)*100) / 100
}

//generic converter to stone that accepts conversion rate
func genericStoneConversion(value float64, rate float64, rate_name string) (int64, float64) {
	stone := value / rate
	fmt.Printf("%v %s = %v Stone\n", value, rate_name, stone)
	// get value for stone that doesnt include any lbs
	baseStone := math.Floor(stone) * 14
	// get the number of remaining lbs to 2 decimal places
	// remainder := math.Floor(((stone*14)-baseStone)*100) / 100
	remainder := toTwoDecimals((stone * 14) - baseStone)
	fmt.Println("or")
	fmt.Printf("%v Stone; %v lbs\n", int64(stone), remainder)
	return int64(stone), remainder
}

//generic convert to lbs that accepts conversion rate
func genericUnitConversion(value float64, rate float64, rate_name string) {
	fmt.Printf("%f\n", value)
	fmt.Printf("%f\n", rate)
	new_value := value / rate
	fmt.Printf("%v = %v %s\n", value, toTwoDecimals(new_value), rate_name)
	fmt.Printf("%v = %v %s\n", value, new_value, rate_name)

}

//lbsToStone converts lbs to stone.
func lbsToStone(lbs float64) {
	fmt.Println("- Convert lbs to stone -")
	// 14 pounds in a stone
	fmt.Printf("Rates conversion from lbsToStone is %v\n", ratesMap["lbsToStone"])
	genericStoneConversion(lbs, ratesMap["lbsToStone"], "lbs")
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
func stoneToLbs(stones float64) {
	fmt.Println("- Convert Stones to lbs -")
	genericUnitConversion(stones, 0.071428571, "stones")
}

func ozToLbs(oz float64) {
	fmt.Println("- Convert oz to lbs - ")
	genericUnitConversion(oz, 0.4536, "ozs")
}

//toOz section

//toGrams section
