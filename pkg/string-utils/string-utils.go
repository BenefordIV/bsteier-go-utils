package string_utils

import "strings"

// ContainsIgnoreCase takes in a target (t) and a value (v). Converts them to lower case and then performs the
// strings.Contains(t, v) to confirm if a value is contained within the target
func ContainsIgnoreCase(t, v string) bool {
	t = strings.ToLower(t)
	v = strings.ToLower(v)

	return strings.Contains(t, v)
}

func EqualsIgnoreCase(t, v string) bool {
	t = strings.ToLower(t)
	v = strings.ToLower(v)

	return strings.EqualFold(t, v)
}
