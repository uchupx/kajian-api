package helper

import "strconv"

// Convert string to int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

// Convert int to string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Convert int64 to string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Convert string to pointer
func StringToPointer(s string) *string {
	return &s
}

// Convert int to pointer
func IntToPointer(i int) *int {
	return &i
}
