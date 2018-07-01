// Package data provides a comman interface for loading data from different formats.
package data


type Data interface {
	NextRecord() []string
	HasMoreRecords() bool 
}