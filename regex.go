package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type md_regex_struct struct {
	md_lang_start_regex    *regexp.Regexp
	md_lang_end_regex      *regexp.Regexp
	md_comment_start_regex *regexp.Regexp
	md_comment_end_regex   *regexp.Regexp
}

func get_md_regex() *md_regex_struct {
	var err error

	md_regex := md_regex_struct{
		md_lang_start_regex:    nil,
		md_lang_end_regex:      nil,
		md_comment_start_regex: nil,
		md_comment_end_regex:   nil,
	}

	md_lang_start_regex_string := "^```\\s*(\\w+)\\s*$"
	md_regex.md_lang_start_regex, err = regexp.Compile(md_lang_start_regex_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling md_lang_scan_regex: '%s'", md_lang_start_regex_string)
		log.Fatal(err)
	}

	md_lang_end_regex_string := "^```"
	md_regex.md_lang_end_regex, err = regexp.Compile(md_lang_end_regex_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling md_lang_end_regex: '%s'", md_lang_end_regex_string)
		log.Fatal(err)
	}

	md_regex.md_comment_start_regex, err = regexp.Compile("^<!--")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling md_comment_start_regex")
		log.Fatal(err)
	}

	md_regex.md_comment_end_regex, err = regexp.Compile("-->")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling md_comment_end_regex")
		log.Fatal(err)
	}

	return &md_regex
}
