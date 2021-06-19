package main

import (
	"log"
	"os"
	"syscall"
)

//
// Moves the temporary files created by process_md into their final
// locations.
func relink_md_files(md_files *[]MD_File) {
	var err error

	for _, md_file := range *md_files {
		for _, output_pair := range md_file.output_pairs {
			// Unlink any previous files at the destination file path; if such exists...
			if _, err := os.Stat(output_pair.output_filename); err == nil {
				err = syscall.Unlink(output_pair.output_filename)
				if err != nil {
					log.Fatal(err)
				}
			}

			// Link the temporary file to the final path
			err = syscall.Link(output_pair.temp_filename, output_pair.output_filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
