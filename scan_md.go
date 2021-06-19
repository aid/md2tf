package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

//
// Scan the given Markdown file; returning a list of the languages
// used to tag Fenced Code Blocks within the file.
func scan_md_languages(input_filename string) []string {
	md_regex := get_md_regex()

	// We just the following map like a set to remove duplicates
	var language_map = make(map[string]bool)

	// Open the input file for reading
	input_file, err := os.Open(input_filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input_file: '%s'\n", input_filename)
		log.Fatal(err)
	}
	defer input_file.Close()

	// We start with the file as Markdown
	var state = State_MD
	var line_n = 1

	// Run through each line of the file
	scanner := bufio.NewScanner(input_file)
	scanner.Split(ScanLinesReturningTerminator)
	for scanner.Scan() {
		line := scanner.Text()

		switch state {
		case State_MD:
			// Check for the start of a Language Block
			matches := md_regex.md_lang_start_regex.FindStringSubmatch(line)
			if len(matches) > 0 {
				var language = string(bytes.Trim([]byte(matches[1]), " \t"))
				if language != "" {
					// fmt.Printf("Found language '%s' on line %d.\n", language, line_n)
					language_map[language] = true
					state = State_MD_FencedBlock
				}
			}
			// Check for the start of a Markdown (HTML) Comment
			if md_regex.md_comment_start_regex.MatchString(line) {
				state = State_MD_Comment
			}
		case State_MD_Comment:
			// Check for the end of a Markdown (HTML) Comment
			// Note that in state State_MD_Comment we're specifically NOT checking for TF content
			if md_regex.md_comment_end_regex.MatchString(line) {
				state = State_MD
			}
		case State_MD_FencedBlock:
			// Check for end of TF
			if md_regex.md_lang_end_regex.MatchString(line) {
				state = State_MD
			}
		default:
			log.Fatal("Unexpected state in md2tf_scan_for_languages")
		}

		line_n += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Convert the map to a list and return the result
	language_list := []string{}
	for key := range language_map {
		language_list = append(language_list, key)
	}
	return language_list
}
