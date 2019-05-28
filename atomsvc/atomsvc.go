package atomsvc

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

type AtomSvc struct {
	Errors []error
	Service Service `xml:"service"`
}

type Service struct {
	XMLName xml.Name `xml:"service"`
	Workspace Workspace `xml:"workspace"`
}

type Workspace struct {
	XMLName xml.Name `xml:"workspace"`
	Collections []Collection `xml:"collection"`
}

type Collection struct {
	XMLName xml.Name `xml:"collection"`
	Href string `xml:"href,attr"`
	HrefLength int
	HostPath string
	Query string
	Options map[string][]string
	Errors []error
}

func FromFile(file_location string) (a AtomSvc) {
	file_reference, file_open_error := os.Open(file_location)
	if (file_open_error != nil) {
		a.Errors = append(a.Errors, file_open_error)
		return
	}
	defer file_reference.Close()

	byte_val, io_error := ioutil.ReadAll(file_reference)
	if (io_error != nil) {
		a.Errors = append(a.Errors, io_error)
		return
	}

	xml.Unmarshal(byte_val, &a.Service)

	for _, cols := range a.Service.Workspace.Collections {
		cols.Options = make(map[string][]string)
	}

	return
}

func (a AtomSvc) AllCollections() (processed_collections []Collection) {
	for _, col := range a.Service.Workspace.Collections {
		col.ParseHref()
		processed_collections = append(processed_collections, col)
	}
	return
}

func (c *Collection) ParseHref() {

	// url_r, parse_error := url.Parse(c.Href)
	// if (parse_error != nil) {
	// 	c.Errors = append(c.Errors, parse_error)
	// }
	
	// This is weird because we are using just the query part instead of the whole url string.
	// This is because the particular reporting service we are using has a weird (and dumb) url structure.
	// Default version of this should probably just PathUnescape c.Href directly.
	result, unescape_error := url.PathUnescape(c.Href)
	if (unescape_error != nil) {
		c.Errors = append(c.Errors, unescape_error)
	}

	c.HrefLength = len(result)

	// More weird stuff to compensate for weird url structure.
	// Test url does not separate path from query with a "?".
	// So we have to do this manually.
	segmentation_count := 0
	var query_string string
	for _, s := range strings.Split(result, "&") {
		if segmentation_count == 0 {
			c.HostPath = s
		} else if segmentation_count == 1 {
			query_string = query_string+s
		} else {
			query_string = query_string+"&"+s
		}
		segmentation_count++
	}
	c.Query = query_string

	split_ops, query_parse_error := url.ParseQuery(c.Query)
	if (query_parse_error != nil) {
		c.Errors = append(c.Errors, query_parse_error)
	}	

	c.Options = split_ops
}

func (c Collection) Display() {
	fmt.Println("HREF Length")
	fmt.Println("	", c.HrefLength)
	fmt.Println("")
	fmt.Println("Host/Path")
	fmt.Println("	", c.HostPath)
	fmt.Println("")
	fmt.Println("Query")
	fmt.Println("	", c.Query)
	fmt.Println("")
	fmt.Println("Options")
	for key, vals := range c.Options {
		fmt.Println("	", key)
		for _, v := range vals {
			fmt.Println("		", v)
		}
	}
	fmt.Println("")
	fmt.Println("Errors")
	fmt.Println("	", c.Errors)
	fmt.Println("")
}