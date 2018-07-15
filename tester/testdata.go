package main

import (
	"blunders"
	"data/meta"
	"reflect"
	"strconv"
)

type TestData struct {
	Lines []TdLine
	MetaLocation string
	RequiredFields []string
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

	// Load required fields
	td.DetermineRequiredFields()

	//Load Meta
	td.MetaLocation = meta_location	
	td.MetaData = meta.NewMeta(td.MetaLocation)
	td.MetaData.HasFields(td.RequiredFields)
	td.Blunders.IncludeBlundersFrom(td.MetaData.Blunders)
	td.MetaData.Blunders = td.Blunders

	return
}

// Utility functions
// Maybe move this to the Meta package? Pass the field as an argument
func (td *TestData) DetermineRequiredFields() (req_fields []string) {
	var sample_line TdLine
	tp := reflect.TypeOf(sample_line)
	for i := 0; i < tp.NumField(); i++ {
		req_fields = append(req_fields, tp.Field(i).Name)
	}
	td.RequiredFields = req_fields
	return
}

func (td *TestData) CreateDefaultMetaFile() {
	td.MetaData.GenerateMetaFile(td.MetaLocation, td.RequiredFields)
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