package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

type Convert struct {
	from  string
	to    string
	value float64
}

func (c Convert) getFunctionName() string {
	functionToExec := strings.Join([]string{c.from, strings.Title(c.to)}, "To")
	return functionToExec
}

func main() {
	fmt.Println("## Conversion program ##")
	from, to, value := getParser()
	var conversion = Convert{
		from:  from,
		to:    to,
		value: value,
	}
	functionToExec := conversion.getFunctionName()
	executeFunction(functionToExec, conversion.value)
}

//Decide which conversion to call
func executeFunction(function string, value float64) {
	switch function {
	case "lbsToStone":
		lbsToStone(value)
	case "kgToStone":
		kgToStone(value)
	case "gramsToStone":
		gramsToStone(value)
	case "ozToStone":
		ozToStone(value)
	}
}

//getParser constructs the argparse.
//It formats the input parameters
//It constructs the command line help
func getParser() (string, string, float64) {
	parser := argparse.NewParser("converter", "Convert Units of measurement from one type to another")
	var convertFrom *string = parser.Selector("f", "from", []string{"kg", "stone", "lbs", "grams", "oz"}, &argparse.Options{Required: true, Help: "Measurement Unit to convert from"})
	var convertTo *string = parser.Selector("t", "to", []string{"kg", "stone", "lbs", "grams", "oz"}, &argparse.Options{Required: true, Help: "Measurement Unit to convert to"})
	var value *float64 = parser.Float("v", "value", &argparse.Options{Required: true, Help: "Value to Convert"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
	}
	return *convertFrom, *convertTo, *value
}

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
