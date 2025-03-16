package iteration

func Repeat(character string) string {
	var repeated string // instantiated as an empty string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
