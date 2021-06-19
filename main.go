package main

import (
	"flag"
	"log"
	"os"
)

//
// Initialise a new MD_File structure based upon
// the provided input filename of a Markdown file.
func new_md_file(input_filename string) MD_File {
	languages := scan_md_languages(input_filename)

	return MD_File{
		input_filename: input_filename,
		languages:      languages,
		output_pairs:   make([]MD_OutputPair, 0, 2),
	}
}

//
// Process the arguments provided on the command line
// when this tool is called.
//
// Primarily; gather the directories and files included
// on the command line and use these to create a slice
// of MD_File objects that describe the selected files.
func process_args(md_files *[]MD_File) {
	flag.Parse()

	path_list := flag.Args()

	for _, path := range path_list {
		// Lets check it exists; and if it is a directory or file...
		stat, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("Path '%s' does not exist.\n", path)
			} else {
				log.Fatal(err)
			}
		} else {
			if stat.IsDir() {
				// md2tf_directories.PushBack(path)
				// mydir := e.Value
				// matches, err := filepath.Glob(mydir)
				// if err != nil {
				// 	fmt.Println(err)
				// }

				// for _, x := range matches {
				// 	fmt.Printf("\t\t%s\n", x)
				// }
			} else {
				md_file := new_md_file(path)
				*md_files = append(*md_files, md_file)
			}
		}
	}
}

// Entry point for this tool.
func main() {
	var md_files []MD_File = make([]MD_File, 0, 10)

	// Run through the provided arguments; adding MD_File objects
	// to the md_files list
	process_args(&md_files)

	// Run through the md_files list; processing each file for each
	// associated language
	process_md_files(&md_files)

	// If there have been no errors; use the 'outputs' data in
	// ms_files to link the move the temporry files to the
	// final (output) firenames
	relink_md_files(&md_files)
}
