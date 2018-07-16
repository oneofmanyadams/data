package data

import (
	"data/csv"
	"blunders"
//	"fmt"
)

type Data interface {
	LoadLine([]string, int) bool
	DataLocation() string
	BlunderBus() *blunders.BlunderBus
}

func Loader(data_source Data) {  // Gotta pass as "&Source" when calling this.
	csv_data := csv.Open(data_source.DataLocation())
	data_source.BlunderBus().IncludeBlundersFrom(csv_data.Blunders)
	csv_data.Blunders = data_source.BlunderBus()

	for line_count := 0; csv_data.HasMoreRecords(); line_count++ {
		raw_line := csv_data.NextRecord()
		if len(raw_line) > 0 {
			data_source.LoadLine(raw_line, line_count)
		}
	}
	
}