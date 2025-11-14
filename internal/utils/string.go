package utils

// UniqueStrings returns a new slice with only unique strings from the input slice.
// The order of elements is preserved.
func UniqueStrings(slice []string) []string {
	keys := make(map[string]struct{})
	list := make([]string, 0, len(slice))
	for _, item := range slice {
		if _, exists := keys[item]; !exists {
			keys[item] = struct{}{}
			list = append(list, item)
		}
	}
	return list
}
