// Package data provides a comman interface for loading data from different formats.
package data


type Data interface {
	Lines() [][]string
	LoadFrom() bool
	WriteTo() bool
}