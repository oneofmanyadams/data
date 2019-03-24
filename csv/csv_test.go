package csv

import (
	"fmt"
)

/////////////////////////////////////////////////////////
//// Examples
/////////////////////////////////////////////////////////

func ExampleNew_loadData() {
	csv_data := New("path/to/csvfile.csv")
	for csv_data.LoadData(); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {
		fmt.Println(csv_data.ActiveRecord) // <-reads one line at a time
	}
	fmt.Println(csv_data.ActiveRecord) // <- This will return nil
}

func ExampleNew_writeData() {
	csv_file := New("path/to/csvfile.csv")
	csv_file.WriteData()
	csv_file.WriteRecord([]string{"1","one"})
	csv_file.WriteRecord([]string{"2","two"})
	csv_file.WriteRecordsToFile()
}

func ExampleLoadNewData() {
	for csv_data := LoadNewData("path/to/csvfile.csv"); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {
		fmt.Println(csv_data.ActiveRecord) // <-reads one line at a time
	}
}

func ExampleWriteNewData() {
	csv_file := WriteNewData("path/to/csvfile.csv")
	csv_file.WriteRecord([]string{"1","one"})
	csv_file.WriteRecord([]string{"2","two"})
	csv_file.WriteRecordsToFile()
}