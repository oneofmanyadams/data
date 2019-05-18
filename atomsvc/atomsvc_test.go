package atomsvc

/////////////////////////////////////////////////////////
//// Examples
/////////////////////////////////////////////////////////

func ExampleFromFile() {
	a_svc := FromFile("path/to/example/file.atomsvc")
	for _, c := range a_svc.Service.Workspace.Collections {
		c.ParseHref()
		c.Display()
	}
}
	
