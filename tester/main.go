package main

import (
	"blunders"
	"data/source"
	"os"
	"fmt"
)

func main() {
	blndrs := blunders.NewBlunderBus()

	test_data := NewTestData("testdata.xml")
	source.LoadCsvDataInto(&test_data) //This pointer thing is weird

	for _, line := range test_data.Lines {
		fmt.Println(line)
	}

	// The include needs to happen at the bottom? (why?)
	blndrs.IncludeBlundersFrom(test_data.Blunders)
	blndrs.LogTo(os.Stdout)
}