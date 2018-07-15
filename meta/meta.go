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
	HasTitleRow bool
	DataAge time.Time `xml:"-"`
	DataPoints []DataPoint `xml:"DataPoint"`
	PointPositions map[string]int `xml:"-"`
	Blunders *blunders.BlunderBus `xml:"-"`
}

type DataPoint struct {
	Name string
	Position int
}

func NewMeta(meta_location string) (meta Meta) {
	meta.Blunders = blunders.NewBlunderBus()
	
	file, file_error := os.Open(meta_location)
	if file_error != nil {
		meta.Blunders.NewFatal("FILE", "Unable to open Meta File: "+file_error.Error())
		return
	}

	byte_val, read_error := ioutil.ReadAll(file)
	if read_error != nil {
		meta.Blunders.NewFatal("FILE", "Unable to read Meta File: "+read_error.Error())
		return
	}

	unmarshal_error := xml.Unmarshal(byte_val, &meta)
	if unmarshal_error != nil {
		meta.Blunders.NewFatal("MARSHALLING", "Unable to Unmarshal meta data: "+unmarshal_error.Error())
		return
	}

	meta.PointPositions = make(map[string]int)
	for _, dp := range meta.DataPoints {
		meta.PointPositions[dp.Name] = dp.Position
	}

	file.Close()

	return
}

func (m *Meta) P(point_name string) (point_position int) {
	point_position = m.PointPositions[point_name]
	return
}

func (m *Meta) GenerateMetaFile(output_location string, data_points []string) {

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
		m.Blunders.New("FILE", "Unable to open file for sample meta data: \""+file_error.Error()+"\"")
	}

	marshaled_data, marshal_error := xml.MarshalIndent(m, "", "	")
	if marshal_error != nil {
		m.Blunders.New("MARSHALLING", "Unable to marshal meta data: \""+marshal_error.Error()+"\"")
	}

	file.Write(marshaled_data)

	m.DataPoints = old_data_points

	file.Close()

}

func (m *Meta) Require(point string) (has_point bool) {
	if _, point_exists := m.PointPositions[point]; !point_exists {
		m.Blunders.NewFatal("DATAPOINT", "Meta file missing data point \""+point+"\"")
		has_point = false
	} else {
		has_point = true
	}
	return
}

func (m *Meta) LoadDataLocationInfo() {
	if m.LocationIsFolder {
		// Don't need this yet...
	} else {
		file, file_error := os.Open(m.DataLocation)
		if file_error != nil {
			m.Blunders.NewFatal("FILE", "Unable to open Data File: "+file_error.Error())
			return
		}
	
		file_stats, file_stat_error := file.Stat()
		if file_stat_error != nil {
			m.Blunders.NewFatal("FILE", "Unable to stat Data File: "+file_stat_error.Error())
			return
		}
		m.DataAge = file_stats.ModTime()
	}
}