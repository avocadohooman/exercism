package accumulate

func Accumulate(collection []string, converter func(string) string) []string {
	var result = make([]string, len(collection))
	for i, str := range collection {
		result[i] = converter(str)
	}
	return result
}
