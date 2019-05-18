package commonquery

import "fmt"

type CommonQuery struct {
	HostPath string
	Options map[string][]string

	QueryPathSeparator string

	FullPath string
	FullPathLength int
}

/////////////////////////////////////////////////////////
//// Create/Modify Methods
/////////////////////////////////////////////////////////

func NewCommonQuery() (ncq CommonQuery) {
	ncq.Options = make(map[string][]string)
	ncq.QueryPathSeparator = "?"
	return
}

func (c *CommonQuery) AddOption(option_name string, option_value string) {
	if existing_values, option_exists := c.Options[option_name]; option_exists {
		value_exists_in_options := false
		for _, existing_val := range existing_values {
			if existing_val == option_value {
				value_exists_in_options = true
				break
			}
		}
		if !value_exists_in_options {
			c.Options[option_name] = append(existing_values, option_value)	
		}
	} else {
		c.Options[option_name] = []string{option_value}
	}
}

/////////////////////////////////////////////////////////
//// Build Methods
/////////////////////////////////////////////////////////

func (c *CommonQuery) BuildFullPath() {
	c.FullPath = c.HostPath+c.QueryPathSeparator
	c.FullPath = c.FullPath+c.BuildQueryFromOptions()
	c.FullPathLength = len(c.FullPath)
}

func (c CommonQuery) BuildQueryFromOptions() (query string) {
	first_value := true
	for option_name, option_values := range c.Options {
		for _, val := range option_values {
			if first_value {
				query = query+option_name+"="+val
				first_value = false
			} else {
				query = query+"&"+option_name+"="+val
			}
		}
	}
	return
}

/////////////////////////////////////////////////////////
//// Utility Methods
/////////////////////////////////////////////////////////

func (c CommonQuery) Display() {
	fmt.Println("HostPath:	", c.HostPath)
	fmt.Println("")
	fmt.Println("Options:	", c.Options)
	fmt.Println("")
	fmt.Println("Separator:	", c.QueryPathSeparator)
	fmt.Println("")
	fmt.Println("FullPath:	", c.FullPath)
	fmt.Println("")
	fmt.Println("FullPathLen:	", c.FullPathLength)
	fmt.Println("")
}