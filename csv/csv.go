package csv

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"simple/sclid"
)

type Csv struct {
	DataLocation csvDataLocation
}

type csvDataLocation struct {
	OriginalPath string
	AbsPath string
	LastModified time.Time
	IsDir bool
	IsCsv bool

	SubDataLocations []csvDataLocation
	UsableFiles []string

	//UseModCutOff bool
	//ModCutOff time.Time

	errs []error
}

func New(data_location string) (new_csv Csv) {
	new_csv.DataLocation = newDataLocation(data_location)
	return
}

func newDataLocation(data_location string) (dl csvDataLocation) {
	dl.OriginalPath = data_location

	// Get full path for location.
	full_path, abs_error := filepath.Abs(data_location)
	if abs_error != nil {
		dl.errs = append(dl.errs, abs_error)
		return
	}
	dl.AbsPath = full_path

	// Verify file/folder path is valid and grab some basic stats.
	stat_data, stat_error := os.Stat(dl.AbsPath)
	if 	stat_error != nil {
		dl.errs = append(dl.errs, stat_error)
		return
	}
	dl.IsDir = stat_data.IsDir()
	dl.LastModified = stat_data.ModTime()

	// If location is directory, get all sub files
	if dl.IsDir {
		sub_stat, dir_error := ioutil.ReadDir(dl.AbsPath)
		if dir_error != nil {
			dl.errs = append(dl.errs, dir_error)
			return	
		}
		for _, ss := range sub_stat {
			nsl := newDataLocation(dl.AbsPath + "/" + ss.Name())
			if nsl.IsCsv || nsl.IsDir {
				dl.SubDataLocations = append(dl.SubDataLocations, nsl)
				dl.UsableFiles = append(dl.UsableFiles, nsl.UsableFiles...)
			}
		}
	}

	// If file, determine if CSV
	if !dl.IsDir {
		if filepath.Ext(dl.AbsPath) == ".csv" {
			dl.IsCsv = true
			dl.UsableFiles = append(dl.UsableFiles, dl.AbsPath)
		} else {
			dl.IsCsv = false
		}
	}

	return
}


func (c *Csv) UpdateStats() {
}

//////////////////////////////////////////////////////
////		Error Methods
//////////////////////////////////////////////////////
func (c *Csv) Error(err error) {
}
func (c *Csv) FatalError(err error) {
}

//////////////////////////////////////////////////////
////		Use Methods
//////////////////////////////////////////////////////

func (c *Csv) Scanner() (scanner Scanner) {
	return
}

//////////////////////////////////////////////////////
////		Info Methods
//////////////////////////////////////////////////////

func (c *Csv) Info() {
	var display_string [][]string
	display_string = append(display_string, []string{"Data Location Info", "", ""})
	display_string = append(display_string, []string{"", "Original Path:", c.DataLocation.OriginalPath})
	display_string = append(display_string, []string{"", "Abs Path:", c.DataLocation.AbsPath})
	display_string = append(display_string, []string{"", "Last Modified:", c.DataLocation.LastModified.Format("2006-01-02 15:04:05")})
	loc_is_csv := "False"
	if c.DataLocation.IsCsv {
		loc_is_csv = "True"
	}
	display_string = append(display_string, []string{"", "Is A CSV file:", loc_is_csv})
	
	loc_is_dir := "False"
	if c.DataLocation.IsDir {
		loc_is_dir = "True"
	}
	display_string = append(display_string, []string{"", "Is Directory:", loc_is_dir})
	display_string = append(display_string, []string{"", "Usable Sub Files:", ""})
	for _, uf := range c.DataLocation.UsableFiles {
		display_string = append(display_string, []string{"", "", uf})
	}

	sclid.Display(display_string)
}