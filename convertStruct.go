package main

import "strings"

type Convert struct {
	from  string
	to    string
	value float64
}

func (c Convert) getFunctionName() string {
	functionName := strings.Join([]string{c.from, strings.Title(c.to)}, "To")
	return functionName
}
