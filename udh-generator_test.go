package main

import "testing"

type decHexTestDataItem struct {
	number int
	dechex string
}

type generateTestDataItem struct {
	total int
	sequence int
	reference int
	udh string
}

var dechexTests = []decHexTestDataItem{
	{ 100, "64" },
	{ 10, "0a" },
}

var generateTests = []generateTestDataItem{
	{1, 1, 123, "0500037b0101"},
	{1, 0, 123, "0500037b0100"},
	{12, 1, 123, "0500037b0c01"},
}


func TestDechexStr(t *testing.T) {

	g := UdhGenerator{}

	for _, pair := range dechexTests {
		result := g.dechexStr(pair.number)
		if result != pair.dechex {
			t.Error(
				"For", pair.number,
				"expected", pair.dechex,
				"got", result,
			)
		}
	}
}

func TestGenerate(t *testing.T)  {

	g := UdhGenerator{}

	for _, item := range generateTests {
		result := g.generate(item.total, item.sequence, item.reference)
		if result != item.udh {
			t.Error(
				"expected", item.udh,
				"got", result,
			)
		}
	}
}
