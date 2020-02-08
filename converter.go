package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

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
