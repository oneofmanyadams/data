package meta

import (
	"os"
	"time"
	"mime"
	"io/ioutil"
	"encoding/xml"
	"path/filepath"
	"blunders"
	"fmt"
)

type Meta struct {
	DataLocation string
	LocationIsFolder bool
	DataType string
	HasTitleRow bool
	DataAge time.Time
	DataPoints []DataPoint `xml:"DataPoint"`
	PointPositions map[string]int
	Blunders blunders.Blunders
}

type DataPoint struct {
	Name string
	Position int
}

func NewMeta(meta_location string) (meta Meta) {
	meta.Blunders = blunders.NewBlunders("META")
	meta.Blunders.AddCode(1, "Load")
	meta.Blunders.AddCode(2, "Unmarshal")
	
	
	meta_type := mime.TypeByExtension(filepath.Ext(meta_location))
	if meta_type != "application/xml" {
		meta.Blunders.NewFatal(1, "Meta file not in xml format. Trying to use: "+meta_location)
	}

	file, file_error := os.Open(meta_location)
	if file_error != nil {
		meta.Blunders.NewFatal(1, "Unable to open Meta File: "+file_error.Error())
	}

	file_stats, file_stat_error := file.Stat()
	if file_stat_error != nil {
		meta.Blunders.NewFatal(1, "Unable to open Meta File: "+file_stat_error.Error())
	}
	meta.DataAge = file_stats.ModTime()

	byte_val, read_error := ioutil.ReadAll(file)
	if read_error != nil {
		meta.Blunders.NewFatal(1, "Unable to read Meta File: "+read_error.Error())
	}

	unmarshal_error := xml.Unmarshal(byte_val, &meta)
	if unmarshal_error != nil {
		meta.Blunders.NewFatal(2, "Unable to Unmarshal meta data: "+unmarshal_error.Error())
	}

	meta.PointPositions = make(map[string]int)
	for _, dp := range meta.DataPoints {
		meta.PointPositions[dp.Name] = dp.Position
	}

	return
}

func (meta Meta) DisplayMeta() {
	fmt.Println("Location:", meta.DataLocation)
	fmt.Println("IsFolder:", meta.LocationIsFolder)
	fmt.Println("MimeType:", meta.DataType)
	fmt.Println("HasTitleRow:", meta.HasTitleRow)
	fmt.Println("DataAge:", meta.DataAge)

	for _, dp := range meta.DataPoints {
		fmt.Println("	", dp)
	}

	for point, position := range meta.PointPositions {
		fmt.Println("	", point, " => ", position)
	}

}