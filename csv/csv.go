package csv

import (
	"os"
	"bufio"
	"encoding/csv"
//	"path/filepath"
//	"time"
)

type Csv struct {
	Path string
	HasMoreRecords bool
	ActiveRecord []string
	Errors []error

	fileReference *os.File
	reader *csv.Reader
}

func New(data_location string) (new_csv Csv) {
	new_csv.Path = data_location
	return
}

/////////////////////////////////////////////////////////
//// Read Functions
/////////////////////////////////////////////////////////

// LoadNewData is a quick way of defining and loading CSV data in one function call.
// Useful for keeping entire csv read logic inside of one for loop.
func LoadNewData(data_location string) (new_csv Csv) {
	new_csv = New(data_location)
	new_csv.LoadData()
	return
}

// LoadData opens the file and reads the first record.
// The first record is read into Cav.ActiveRecord
func (c *Csv) LoadData() () {
	// Open csv file
	file_reference, file_open_error := os.Open(c.Path)
	if (file_open_error != nil) {
		c.Errors = append(c.Errors, file_open_error)
		return
	}
	c.fileReference = file_reference

	// Read file reference into csv reader
	c.reader = csv.NewReader(bufio.NewReader(c.fileReference))

	// Load first record
	c.LoadNextRecord()

	return
} 

// LoadNextRecord populates the ActiveRecord value.
// Sets HasMoreRecords value
func (c *Csv) LoadNextRecord() {
	read_record, read_error := c.reader.Read()
	if read_error != nil {
		if read_error.Error() == "EOF" {
			c.HasMoreRecords = false
			c.fileReference.Close()
		} else {
			c.Errors = append(c.Errors, read_error)
		}
	} else {
		c.HasMoreRecords = true
	}
	c.ActiveRecord = read_record	
}

/////////////////////////////////////////////////////////
//// Write Functions
/////////////////////////////////////////////////////////
