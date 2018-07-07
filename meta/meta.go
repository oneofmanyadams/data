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
	DataAge time.Time `xml:"-"`
	DataPoints []DataPoint `xml:"DataPoint"`
	PointPositions map[string]int `xml:"-"`
	Blunders blunders.Blunders `xml:"-"`
}

type DataPoint struct {
	Name string
	Position int
}

func NewMeta(meta_location string) (meta Meta) {
	meta.Blunders = blunders.NewBlunders("META")
	meta.Blunders.AddCode(1, "File")
	meta.Blunders.AddCode(2, "Marshalling")
	meta.Blunders.AddCode(3, "DataPoint")
	
	
	meta_type := mime.TypeByExtension(filepath.Ext(meta_location))
	if meta_type != "application/xml" {
		meta.Blunders.NewFatal(1, "Meta file not in xml format. Trying to use: "+meta_location)
	}

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

	meta.PointPositions = make(map[string]int)
	for _, dp := range meta.DataPoints {
		meta.PointPositions[dp.Name] = dp.Position
	}

	return
}

func (m *Meta) GenerateMetaFile(data_points []string, output_location string) {

	old_data_points := m.DataPoints
	if len(data_points) > 0 {
		m.DataPoints = nil

		for po, dp := range data_points {
			var new_dp DataPoint
			new_dp.Name = dp
			new_dp.Position = po
			m.DataPoints = append(m.DataPoints, new_dp)
		}
	}

	file, file_error := os.Create(output_location)
	if file_error != nil {
		m.Blunders.New(1, "Unable to open file for sample meta data: \""+file_error.Error()+"\"")
	}

	marshaled_data, marshal_error := xml.MarshalIndent(m, "", "	")
	if marshal_error != nil {
		m.Blunders.New(2, "Unable to marshal meta data: \""+marshal_error.Error()+"\"")
	}

	file.Write(marshaled_data)

	m.DataPoints = old_data_points

	file.Close()

}

func (m Meta) HasDataPoints(required_points []string) (has_all bool) {
	has_all = true
	for _, rp := range required_points {
		if _, rp_exists := m.PointPositions[rp]; !rp_exists {
			m.Blunders.NewFatal(3, "Meta file missing data point \""+rp+"\"")
			has_all = false
		}
	}
	return
}

func (m *Meta) LoadDataLocationInfo() {
	if m.LocationIsFolder {
		// Don't need this yet...
	} else {
		file, file_error := os.Open(m.DataLocation)
		if file_error != nil {
			m.Blunders.NewFatal(1, "Unable to open Data File: "+file_error.Error())
		}
	
		file_stats, file_stat_error := file.Stat()
		if file_stat_error != nil {
			m.Blunders.NewFatal(1, "Unable to stat Data File: "+file_stat_error.Error())
		}
		m.DataAge = file_stats.ModTime()
	}
}

func (m Meta) DisplayMeta() {
	fmt.Println("Location:", m.DataLocation)
	fmt.Println("IsFolder:", m.LocationIsFolder)
	fmt.Println("MimeType:", m.DataType)
	fmt.Println("HasTitleRow:", m.HasTitleRow)
	fmt.Println("DataAge:", m.DataAge)

	for _, dp := range m.DataPoints {
		fmt.Println("	", dp)
	}

	for point, position := range m.PointPositions {
		fmt.Println("	", point, " => ", position)
	}

}