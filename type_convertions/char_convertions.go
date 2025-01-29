package typeconvertions

func ConvertCharToByte(str string) byte {
	return []byte(str)[0]
}

func ConvertByteToString(input byte) string {
	return string(input)
}
