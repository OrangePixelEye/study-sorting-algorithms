package typeconvertions

func ZeroOrOneIntoBool(val int) bool {
	return val == 1
}

func BooleanIntoInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}
