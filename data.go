package data

type Data interface {
	Lines() [][]string
	LoadFrom() bool
	WriteTo() bool
}