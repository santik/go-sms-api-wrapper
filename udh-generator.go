package main

import (
	"strings"
	"fmt"
)

type UdhGenerator struct {
}

func (generator UdhGenerator) generate(total int, sequence int, reference int) (string)  {
	octet1 := "05"
	octet2 := "00"
	octet3 := "03"
	octet4 := generator.dechexStr(reference)
	octet5 := generator.dechexStr(total)
	octet6 := generator.dechexStr(sequence)

	s := []string{octet1, octet2, octet3, octet4, octet5, octet6}

	return strings.Join(s, "")
}

func (generator UdhGenerator) dechexStr(ref int)(string)  {

	start := ""
	if ref <= 15 {
		start = string('0')
	}

	s:= []string{start, fmt.Sprintf("%x", ref)}

	return strings.Join(s, "")
}