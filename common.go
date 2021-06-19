package main

// Struct used to record the state of the tool as it
// moves through a Markdown file.
const (
	State_MD               = iota // General Markdown Prose
	State_MD_Comment       = iota // Within a Markdown (HTML) Comment
	State_MD_FencedBlock   = iota // Within a Markdown Fenced Block that we do not care about
	State_MD_LanguageBlock = iota // Within a Markdown Fenced Block that we do care about
)

// Control for whether we show debug output, or not
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
