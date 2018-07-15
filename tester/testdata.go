package main

import (
	"blunders"
	"data/meta"
	"strconv"
)

type TestData struct {
	Lines []TdLine
	MetaLocation string
	Blunders *blunders.BlunderBus
	MetaData meta.Meta
}

type TdLine struct {
	Name string
	Age int
	Position string
}

func NewTestData(meta_location string) (td TestData) {
	// implement blunders
	td.Blunders = blunders.NewBlunderBus()

	//Load Meta
	td.MetaLocation = meta_location	
	td.MetaData = meta.NewMeta(td.MetaLocation)
	var sample_line TdLine
	td.MetaData.DetermineRequiredFields(sample_line)
	td.Blunders.IncludeBlundersFrom(td.MetaData.Blunders)
	td.MetaData.Blunders = td.Blunders

	return
}

func (td *TestData) CreateDefaultMetaFile() {
	var sample_line TdLine
	td.MetaData.GenerateMetaFile(td.MetaLocation, sample_line)
}

// Source required Methods
func (td *TestData) LoadLine(raw_line []string, line_number int) (result bool) {
	if line_number == 0 && td.MetaData.HasTitleRow {
		return
	}
	var new_line TdLine

	new_line.Name = raw_line[td.MetaData.P("Name")]
	
	var age_error error
	new_line.Age, age_error = strconv.Atoi(raw_line[td.MetaData.P("Age")])
	if age_error != nil {
		td.Blunders.New("DATACONVERSION-age", age_error.Error())
	}
	
	new_line.Position = raw_line[td.MetaData.P("Position")]

	td.Lines = append(td.Lines, new_line)

	result = true
	return
}

func (td *TestData) DataLocation() string {
	return td.MetaData.DataLocation
}

func (td *TestData) BlunderBus() *blunders.BlunderBus {
	return td.Blunders
}