// data/csv package is a simple wrapper for the builtin csv package.
// simplifies a lot of the typical and basic logic of reading and writing csv data.
package csv

import (
	"os"
	"bufio"
	"encoding/csv"
)

type Csv struct {
	Path string
	HasMoreRecords bool
	ActiveRecord []string
	Errors []error

	SkipFirstLine bool

	fileReference *os.File
	reader *csv.Reader
	writer *csv.Writer
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
// !!! This skips the first line by default. 
func LoadNewData(data_location string) (new_csv Csv) {
	new_csv = New(data_location)
	new_csv.SkipFirstLine = true
	new_csv.LoadData()
	return
}

// LoadData opens the file and reads the first record into Csv.ActiveRecord
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
	if c.SkipFirstLine && c.HasMoreRecords {
		c.LoadNextRecord()
	}

	return
}

// LoadNextRecord populates the ActiveRecord value.
// Sets HasMoreRecords value
func (c *Csv) LoadNextRecord() {
	read_record, read_error := c.reader.Read()
	if read_error != nil {
		if read_error.Error() == "EOF" {
			c.ActiveRecord = nil
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

// ForceClose allows a manual closing of the file if needed.
func (c *Csv) ForceClose() {
	c.ActiveRecord = nil
	c.HasMoreRecords = false
	c.fileReference.Close()
}

/////////////////////////////////////////////////////////
//// Write Functions
/////////////////////////////////////////////////////////

// WriteNewData is a quick way of creating a new Csv instance and Writer in the same call.
func WriteNewData(data_location string) (new_csv Csv) {
	new_csv = New(data_location)
	new_csv.WriteData()
	return
}

// WriteData opens the file and creates the csv reference.
// If the file does not exist it is created.
func (c *Csv) WriteData() () {
	// Open csv file
	file_reference, file_open_error := os.Create(c.Path)
	if (file_open_error != nil) {
		c.Errors = append(c.Errors, file_open_error)
		return
	}
	c.fileReference = file_reference

	// Read file reference into csv reader
	c.writer = csv.NewWriter(c.fileReference)
	
	return
}

// Write Record writes the rovided record to the csv file reference.
// The data will not actually be written to the file until WriteRecordsToFile is called.
func (c *Csv) WriteRecord(record []string) {
	write_error := c.writer.Write(record)
	if write_error != nil {
		c.Errors = append(c.Errors, write_error)
	}
}

// WriteRecordsToFile is a simple wrapper for Flush.
// I'm sure a better way to do this is possible.
func (c *Csv) WriteRecordsToFile() {
	c.writer.Flush()
}