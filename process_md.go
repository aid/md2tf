package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

//
// 1. Process the provided Markdown file.
// 2. Comment out all lines within the markdown file,
// 3. EXCEPT for those lines within a FencedBlock that is tagged with the supplied language
//
func process_file(input_file *os.File, output_file *os.File, language string) error {
	md_regex := get_md_regex()

	// We start with the file state as Markdown
	var state = State_MD
	var line_n = 1

	// Run through each line of the file
	scanner := bufio.NewScanner(input_file)
	scanner.Split(ScanLinesReturningTerminator)
	for scanner.Scan() {
		line := scanner.Text()

		switch state {

		// Normal Markdown prose...
		case State_MD:
			// Check for the start of a Language Block
			matches := md_regex.md_lang_start_regex.FindStringSubmatch(line)
			if len(matches) > 0 {
				var match_language = string(bytes.Trim([]byte(matches[1]), " \t"))
				if debug {
					fmt.Printf("Found language '%s' on line %d.\n", match_language, line_n)
				}
				if match_language == language {
					if debug {
						fmt.Printf("Entering State_MD_LanguageBlock on line %d\n", line_n)
					}
					state = State_MD_LanguageBlock
				} else {
					if debug {
						fmt.Printf("Entering State_MD_FencedBlock on line %d\n", line_n)
					}
					state = State_MD_FencedBlock
				}
			}
			// Check for the start of a Markdown (HTML) Comment
			if md_regex.md_comment_start_regex.MatchString(line) {
				state = State_MD_Comment
			}
			fmt.Fprintf(output_file, "# %s", line)

		// Within a Markdown (HTML) comment...
		case State_MD_Comment:
			// Check for the end of a Markdown (HTML) Comment
			// Note that in state State_MD_Comment we're specifically NOT checking for TF content
			if md_regex.md_comment_end_regex.MatchString(line) {
				state = State_MD
			}
			fmt.Fprintf(output_file, "# %s", line)

		// Within a Markdown Fenced Block that is not for this language
		case State_MD_FencedBlock:
			if md_regex.md_lang_end_regex.MatchString(line) {
				state = State_MD
			}
			fmt.Fprintf(output_file, "# %s", line)

		// Within a Markdown Fenced Block containing this language
		case State_MD_LanguageBlock:
			if md_regex.md_lang_end_regex.MatchString(line) {
				fmt.Fprintf(output_file, "# %s", line)
				state = State_MD
			} else {
				// This is a line within our language block; so we output it
				// WITHOUT the comment symbol (hash):
				fmt.Fprintf(output_file, "%s", line)
			}

		// State error - we should never get here...
		default:
			log.Fatal("Unexpected state in md2tf_process_file")
		}

		line_n += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// return a nil error
	return nil
}

//
// Use the provided input_filename and the provided language
// to create a temporary file; returning the file and its
// filename to the caller.
//
// The caller is responsible for closing the returned file.
func create_temp_file(input_filename string, language string) (temp_filename string, temp_file *os.File, err error) {
	if debug || true {
		fmt.Printf("md2tf_create_temp_file: input_filename='%s' language='%s'\n", input_filename, language)
	}
	input_dir := path.Dir(input_filename)
	input_base := path.Base(input_filename)
	input_extension := path.Ext(input_filename)
	temp_pattern := input_base[0:len(input_base)-len(input_extension)] + ".*." + language

	if debug || true {
		fmt.Printf("md2tf_create_temp_file: temp_dir='%s' temp_pattern='%s'\n", input_dir, temp_pattern)
	}
	temp_file, err = ioutil.TempFile(input_dir, temp_pattern)
	if err != nil {
		return "", nil, err
	}
	temp_filename = temp_file.Name()

	if debug || true {
		fmt.Printf("md2tf_create_temp_file: temp_filename='%s' \n", temp_filename)
	}

	// All good; return the filename of the tempfile, a handle to the temp file and a nil error
	return temp_filename, temp_file, nil
}

//
// Open the provided Markdown file
// Comment out all lines within the markdown file.
// EXCEPT for those lines within a FencedBlock that is tagged with the supplied language
func process_file_with_language(input_filename string, language string) (temp_filename string, output_filename string, err error) {
	if debug {
		fmt.Printf("Entered md2tf_process_file: input_filename='%s' language='%s'\n", input_filename, language)
	}

	// Open the input file for reading
	input_file, err := os.Open(input_filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input_file: '%s'\n", input_filename)
		log.Fatal(err)
	}
	defer input_file.Close()

	// define an output filename, based upon input filename
	input_extension := path.Ext(input_filename)
	output_filename = input_filename[0:len(input_filename)-len(input_extension)] + "." + language

	// Create a temporary file to which we write our output; and set it to close (deferred)
	temp_filename, temp_file, err := create_temp_file(input_filename, language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temprorary file for input_file '%s' and language '%s'\n", input_filename, language)
		log.Fatal(err)
	}
	defer temp_file.Close()

	if debug || true {
		fmt.Printf("md2tf_process_file: temp_filename='%s'\n", temp_filename)
	}
	// fmt.Printf("Using output filename '%s'.\n", output_filename)

	err = process_file(input_file, temp_file, language)
	if err != nil {
		return "", "", err
	}

	// All good; return the temporary and output filenames and a nil error..
	return temp_filename, output_filename, nil
}

//
// Takes a slice of MD_File objects and uses the data within
// to process each input file into one temporary file for
// each required language within that file.
//
// This function updates the MD_File slice to include a mapping
// of temporary to final (output) filenames; to allow
// the caller to place the files in their final location
// if no errors have been detected.
func process_md_files(md_files *[]MD_File) {
	for _, md_file := range *md_files {
		for _, language := range md_file.languages {
			// process this markdown file for the given language; storing the output
			// in a temporary file
			temp_filename, output_filename, err := process_file_with_language(md_file.input_filename, language)
			if err != nil {
				log.Fatal(err)
			}
			// record the temporary filename to final (output) filename for this file/language combination
			md_file.output_pairs = append(md_file.output_pairs, MD_OutputPair{
				temp_filename:   temp_filename,
				output_filename: output_filename,
			})
		}
	}
}
