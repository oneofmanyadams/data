package commonquery

type CommonQuery struct {
	FullPath string
	FullPathLength string
	HostPath string
	Options map[string][]string
}

func NewCommonQuery() (ncq CommonQuery) {
	ncq.Options = make(map[string][]string)
	return
}

func (c *CommonQuery) LoadFromFile(file_location string) {

}

func (c CommonQuery) SaveToFile(file_location string) {

}

func (c *CommonQuery) BuildPath() {

}

func (c *CommonQuery) BreakdownPath() {
	
}