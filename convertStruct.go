package main

import (
	"fmt"
	"strings"
)

type Convert struct {
	from  string
	to    string
	value float64
}

func (c Convert) getFunctionName() string {
	functionName := strings.Join([]string{c.from, strings.Title(c.to)}, "To")
	return functionName
}

//Decide which conversion to call
func (c Convert) executeFunction(function string, value float64) {
	switch function {
	case "lbsToStone":
		lbsToStone(value)
	case "kgToStone":
		kgToStone(value)
	case "gramsToStone":
		gramsToStone(value)
	case "ozToStone":
		ozToStone(value)
	case "stoneToLbs":
		stoneToLbs(value)
	default:
		fmt.Printf("Rates conversion from %s is %f\n", function, ratesMap[function])
		genericUnitConversion(value, ratesMap[function], c.to)
	}
}
