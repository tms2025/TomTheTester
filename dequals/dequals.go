package dequals

// Deep Equals (dequals) provides functions for deep equal comparsion of objects

// IgnoreTimestamps performs deep equal comparsion of 2 objects, ignoring any timestamp (time.Time) fields.
func IgnoreTimestamps[T any](a, b T) bool {
	return false
}
