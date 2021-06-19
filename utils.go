package main

import (
	"bytes"
)

// Returns true if the provided string, s, is present within
// the array slice, s, else returns false.
func string_present_in_list(s string, l *[]string) bool {
	for _, v := range *l {
		if v == s {
			return true
		}
	}
	return false
}

// Checks to see if two array slices of contain exactly
// the same strings; although without concern for the
// order.
//
// Returns true if both lists contains the same strings;
// else returns false.
func unordered_lists_are_equal(a *[]string, b *[]string) bool {
	// First check lengths match
	if len(*a) != len(*b) {
		return false
	}

	// Check all items of a are in b
	for _, a_v := range *a {
		if !string_present_in_list(a_v, b) {
			return false
		}
	}

	// Check all items of b are in a
	for _, b_v := range *b {
		if !string_present_in_list(b_v, a) {
			return false
		}
	}
	return true
}

// This is a customised version of the standard Go ScanLines
// function; however this splits into lines whilst including
// and line terminators - unlike the original Go function that
// strips the line terminators.
func ScanLinesReturningTerminator(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0 : i+1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
