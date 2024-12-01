// package count the word with custom function
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the word number of s
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
