package csv

import (
	"os"
	"io"
	"bufio"
	"encoding/csv"
	"blunders"
)

type Csv struct {
	Blunders blunders.Blunders
	Location string
	LinesRead int
	LineRecords int
	AllDataRead bool
	Data *csv.Reader
	file *os.File
}

func NewCsv() (new_file Csv) {
	new_file.Blunders = blunders.NewBlunders("CSV")
	new_file.Blunders.AddCode(1, "DataLocation")
	new_file.Blunders.AddCode(2, "LineProblem")
	return
}

func (c *Csv) Open(file_location string) (load_success bool) {
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

func OpenNew(file_location string) (new_csv Csv) {
	new_csv = NewCsv()
	new_csv.Open("./data.csv")
	return
}

func (c *Csv) Close() {
	//this is closing the file too early!
	//defer c.file.Close()
}

func (c *Csv) AllDone() bool {
	if c.AllDataRead {
		return true
	}
	if c.Blunders.HasFatal() {
		return true
	}
	return false
}

func (c *Csv) NextLine() (line []string) {
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
	c.LineRecords = len(read_line)

	line = read_line

	return
}