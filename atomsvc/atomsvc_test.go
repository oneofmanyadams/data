package atomsvc

/////////////////////////////////////////////////////////
//// Examples
/////////////////////////////////////////////////////////

func ExampleReadAtomSvc() {
	a_svc := ReadAtomSvc("path/to/example/file.atomsvc")
	for _, c := range a_svc.Service.Workspace.Collections {
		c.ParseHref()
		c.Display()
	}
}

func ExampleToJsonFile() {
	a_svc := ReadAtomSvc("path/to/example/file.atomsvc")
	for _, c := range a_svc.Service.Workspace.Collections {
		c.ParseHref()
		c.ToJsonFile("path/to/example/file.json")
	}
}

func ExampleFromJsonFile() {
	var asvc AtomSvc
	var new_collec Collection
	new_collec.FromJsonFile("path/to/example/file.json")
	asvc.Service.Workspace.Collections = append(asvc.Service.Workspace.Collections, new_collec)
	
	for _, c := range asvc.Service.Workspace.Collections {
		c.Display()
	}
	
}