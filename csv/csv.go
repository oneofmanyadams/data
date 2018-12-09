package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
)

type Csv struct {
	DataLocation string
	DataPath string
	LocationIsFolder bool
	ReadableFiles []string

	file *os.File
	reader_writer *bufio.ReadWriter

	Errors []error

	InUsableState bool
}

func New(data_location string) (new_csv Csv) {
	new_csv.DataLocation = data_location
	new_csv.InUsableState = true
	new_csv.UpdateStats()
	return
}

func (c *Csv) UpdateStats() {
	// open file
	f, open_error := os.Open(c.DataLocation)
	defer f.Close()
	if open_error != nil {
		c.FatalError(open_error)
		return
	}

	// get basic stats
	c.file = f
	file_stats, file_stats_error := c.file.Stat()
	if file_stats_error != nil {
		c.FatalError(file_stats_error)
		return
	}
	c.LocationIsFolder = file_stats.IsDir()

	abs_path, abs_path_error := filepath.Abs(c.DataLocation)
	if abs_path_error != nil {
		c.FatalError(abs_path_error)
		return
	}
	c.DataPath = abs_path

	// Load sub-file names of all csv files
	// Only done if LocationIsFolder
	if c.LocationIsFolder {
		all_sub_file_names, dir_name_error := c.file.Readdirnames(0)
		if dir_name_error != nil {
			c.Error(dir_name_error)
		}
		var sub_files []string
		for _, name := range all_sub_file_names {
			if filepath.Ext(name) == ".csv" {
				sub_files = append(sub_files, name)
			}
		}
		c.ReadableFiles = sub_files
	}
}

//////////////////////////////////////////////////////
////		Error Methods
//////////////////////////////////////////////////////
func (c *Csv) Error(err error) {
	c.Errors = append(c.Errors, err)
}
func (c *Csv) FatalError(err error) {
	c.InUsableState = false
	c.Errors = append(c.Errors, err)
}

//////////////////////////////////////////////////////
////		Use Methods
//////////////////////////////////////////////////////

func (c *Csv) GenerateScanner() (scanner Scanner) {
	var base_reader io.Reader
	
	if c.LocationIsFolder {
		var reader_groups []io.Reader
		for _, sub_file := range c.ReadableFiles {
			fi, fi_error := os.Open(c.DataPath+"/"+sub_file)
			defer fi.Close()
			if fi_error != nil {
				c.FatalError(fi_error)
				return
			}
			reader_groups = append(reader_groups, bufio.NewReader(fi))
		}
		base_reader = io.MultiReader(reader_groups...)
	} else {
		fi, fi_error := os.Open(c.DataPath)
		defer fi.Close()
		if fi_error != nil {
			c.FatalError(fi_error)
			return
		}
		base_reader = bufio.NewReader(fi)
	}

	csv.NewReader(base_reader)

	return
}