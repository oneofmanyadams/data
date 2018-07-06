package meta

import (
	"os"
	"time"
	"io/ioutil"
	"encoding/xml"
	"blunders"
)

type Meta struct {
	DataLocation string
	LocationIsFolder bool
	DataType string
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
	
	
	file, file_error := os.Open(meta_location)
	if file_error != nil {
		meta.Blunders.NewFatal(1, "Unable to open Meta File: "+file_error.Error())
	}

	byte_val, read_error := ioutil.ReadAll(file)
	if read_error != nil {
		meta.Blunders.NewFatal(1, "Unable to read Meta File: "+read_error.Error())
	}

	unmarshal_error := xml.Unmarshal(byte_val, &meta)
	if unmarshal_error != nil {
		meta.Blunders.NewFatal(2, "Unable to Unmarshal meta data: "+unmarshal_error.Error())
	}
	return
}