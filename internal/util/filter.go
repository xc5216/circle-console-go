package util

func Filter[T any](slice []T, f func(T) bool) []T {
	filtered := []T{}
	for _, s := range slice {
		if f(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
