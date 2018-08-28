package accumulate

func Accumulate(input []string, converter func(string) string) []string {
	output := make([]string, 0)

	for _, word := range input {
		output = append(output, converter(word))
	}

	return output
}
