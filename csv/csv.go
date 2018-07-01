// Package csv provides a standardized way of working with CSV data.
package csv

import (
	"os"
	"io"
	"bufio"
	"encoding/csv"
	"blunders"
)

// Csv is the main type for the Csv package.
// Can be initialized with either the NewCsv() or Open() function, but should not
// be called directly.
//  - Location is the path to the file that is to be read.
//  - LinesRead is a count of the lines read from the file using the NextRecord() method.
//  - AllDataRead is used to determine if EOF has been reached while using the NextRecord() method.
//  - Data is a pointer to a encoding/csv instance of NewReader()
//  - Blunders is the implementation of a custom package that expands error recording and handling.
type Csv struct {
	Location string
	LinesRead int
	AllDataRead bool
	Data *csv.Reader
	file *os.File
	Blunders blunders.Blunders
}

//////////////////////////////////////////////////////////////////
// Init Functions
//////////////////////////////////////////////////////////////////

// NewCsv is the most basic way to initialize a Csv instance.
// Initializes the Blunders instance.
// This is where all Blunder Codes are defined.
func NewCsv() (new_file Csv) {
	new_file.Blunders = blunders.NewBlunders("CSV")
	new_file.Blunders.AddCode(1, "DataLocation")
	new_file.Blunders.AddCode(2, "LineProblem")
	return
}

// Open is a more direct way of creating a new Csv instance.
// Calls NewCsv() and also automatically opens the file  with OpenFile().
func Open(file_location string) (new_csv Csv) {
	new_csv = NewCsv()
	new_csv.OpenFile("./data.csv")
	return
}

//////////////////////////////////////////////////////////////////
// Open / Close Methods
//////////////////////////////////////////////////////////////////

// OpenFile attempts to open a csv file at the location provided.
// Will save a pointer to the file at .file and a pointer to the csv.Reader at .Data.
// Returns true if the file was loaded and false if it was not loaded.
func (c *Csv) OpenFile(file_location string) (load_success bool) {
	file, file_error := os.Open(file_location)
	if file_error != nil {
		c.Blunders.NewFatal(1, file_error.Error())
		load_success = false
		return
	}

	c.file = file
	c.Location = file_location

	c.Data = csv.NewReader(bufio.NewReader(c.file))

	load_success = true
	return
}

// Close closes the file connection that was opened during the OpenFile() method call.
func (c *Csv) Close() {
	// Be smart about where you use this.
	// Because the read is buffered, (through bufio) prematurely closing the file
	// will still allow some data (10-20 lines) to be read.
	c.file.Close()
}

//////////////////////////////////////////////////////////////////
// Read Methods
//////////////////////////////////////////////////////////////////

// NextRecord returns the next un-read line from the csv file.
// If it encounters an io.EOF as a read error it will set
// .AllDataRead to true.
// Increments .LinesRead by 1.
func (c *Csv) NextRecord() (line []string) {
	read_line, line_error := c.Data.Read()

	if line_error != nil {
		if line_error == io.EOF {
			c.AllDataRead = true
		} else {
			c.Blunders.NewFatal(2, line_error.Error())
		}
		return
	}

	c.LinesRead = c.LinesRead + 1

	line = read_line

	return
}

// HasMoreRecords returns false if there is either a fatal blunder or if an
// EOF was encountered during a NextRecord() call.
func (c *Csv) HasMoreRecords() bool {
	if c.AllDataRead {
		return false
	}
	if c.Blunders.HasFatal() {
		return false
	}
	return true
}