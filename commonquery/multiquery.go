package commonquery

import (
	"os"

	"encoding/json"
	"io/ioutil"
)

type MultiQuery struct {
	Queries map[string]CommonQuery
	Errors []error
}

/////////////////////////////////////////////////////////
//// Create/Modify Methods
/////////////////////////////////////////////////////////

func NewMultiQuery() (nmc MultiQuery) {
	nmc.Queries = make(map[string]CommonQuery)
	return
}

func (m *MultiQuery) AddCommonQuery(identifier string, query CommonQuery) {
	m.Queries[identifier] = query
}

/////////////////////////////////////////////////////////
//// JSON Loading/Saving Methods
/////////////////////////////////////////////////////////

func (m_q *MultiQuery) LoadFrom(file_location string) {
	// Open JSON file
	file_reference, file_open_error := os.Open(file_location)
	defer file_reference.Close()
	if file_open_error != nil {
		m_q.Errors = append(m_q.Errors, file_open_error)
		return
	}
	// Read JSON file data
	json_bytes, data_read_error := ioutil.ReadAll(file_reference)
	if data_read_error != nil {
		m_q.Errors = append(m_q.Errors, data_read_error)
		return
	}
	json.Unmarshal(json_bytes, &m_q)

	// rebuild FullPath strings on each CommonQuery
	for key, common_query := range m_q.Queries {
		common_query.BuildFullPath()
		m_q.Queries[key] = common_query
	}
}

func (m_q *MultiQuery) SaveTo(file_location string) {
	//Marshal MoCollection
	json_data, marshal_error := json.MarshalIndent(m_q, "", "	")
	if marshal_error != nil {
		m_q.Errors = append(m_q.Errors, marshal_error)
		return
	}
	// Create/Open json file
	file_reference, file_open_error := os.Create(file_location)
	defer file_reference.Close()
	if file_open_error != nil {
		m_q.Errors = append(m_q.Errors, file_open_error)
		return
	}
	
	file_reference.Write(json_data)
}