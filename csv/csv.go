package csv

import (
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"time"
)

type Csv struct {
	Path string
	HasMoreRecords bool
	ActiveRecord []string
	Errors []error
}

func New(data_location string) (new_csv Csv) {
	new_csv.Path = data_location
	return
}

func (c *Csv) LoadData() () {
	
} 

func (c Csv) LoadNextRecord() (next_record []string) {
	return
}

