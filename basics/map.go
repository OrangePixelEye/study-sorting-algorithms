package basics

func UniqueStrings(elements []string) map[string]struct{} {
	myMap := make(map[string]struct{})
	for _, element := range elements {
		myMap[element] = struct{}{}
	}

	return myMap
}
