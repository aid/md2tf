package main

const (
	State_MD               = iota
	State_MD_Comment       = iota
	State_MD_FencedBlock   = iota
	State_MD_LanguageBlock = iota
)

var debug = false

// Struct that maps the temporary filenames that are initially
// created to the final filenames to which these files should
// finally be linked.
type MD_OutputPair struct {
	temp_filename   string
	output_filename string
}

// Struct that stores essential data for each of the Markdown
// files that are processed by this tool
type MD_File struct {
	input_filename string
	languages      []string
	output_pairs   []MD_OutputPair
}
