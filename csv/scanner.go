package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Scanner struct {
	SourceFiles []*os.File
	BaseReader bufio.Reader
	Reader *csv.Reader	
	EOF bool
	activeLine []string
}

func (s *Scanner) Scan() (eof bool) {
	if s.Reader == nil {
		s.EOF = true
		eof = true
		s.Close()
		return
	}
	new_line, nl_error := s.Reader.Read()
	if nl_error == io.EOF {
		s.EOF = true
		s.activeLine = []string{}
		eof = true
		s.Close()
		return
	}
	if nl_error != nil {
		s.EOF = true
		eof = true
		fmt.Println(nl_error.Error())
		s.Close()
	}
	s.activeLine = new_line
	return
}

func (s *Scanner) Line() []string {
	return s.activeLine
}

func (s *Scanner) Close() {
	for _, file := range s.SourceFiles {
		file.Close()
	}
}