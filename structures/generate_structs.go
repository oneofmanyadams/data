package structures

import "io/ioutil"

func Permutations(start_string string, options []string) (perms []string) {

	for _, opp := range options {
		perms = append(perms, start_string+opp)
	}

	return
}

func Gen() {
	// If you add to types you need to make sure to add the same amount to fields
	types := make(map[string]string)
	types["S"] = "string"
	types["I"] = "int"
	types["F"] = "float"
	types["D"] = "time.Time"
	types["B"] = "bool"

	var ids []string
	for id := range types {
		ids = append(ids, id)
	}

	// this must have the same amount of records as there are types
	fields := []string{"A", "B", "C", "D", "E"}

	var main_string string

	for _, i := range ids {
		perms := Permutations(i, ids)
		for _, i := range perms {
			perms := Permutations(i, ids)
			for _, i := range perms {
				perms := Permutations(i, ids)
				for _, i := range perms {
					perms := Permutations(i, ids)
					for _, i := range perms {
						the_string := "type "+i+" struct {\n"
						field_count := 0
						for _, f := range fields {
							the_string = the_string + "	"+ f
							the_string = the_string + "	"+ types[string(i[field_count])]
							the_string = the_string + "\n"
							field_count = field_count + 1
						}
						the_string = the_string +"}\n"
						main_string = main_string + the_string
					}	
				}
			}
		}
	}
	
	ioutil.WriteFile("dump.txt", []byte(main_string), 0644)
}