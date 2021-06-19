package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
)

func process_and_compare(
	t *testing.T,
	test_name string,
	language string,
	input_filename string,
	actual_output_filename string,
	prepared_output_filename string) {

	// process the file
	temp_filename, _, err := process_file_with_language(input_filename, language)
	if err != nil {
		t.Errorf("%s error recieved from process_file_with_language: %s", test_name, err)
	}

	// test that the temp file was created
	if _, err := os.Stat(temp_filename); errors.Is(err, os.ErrNotExist) {
		t.Errorf("%s did not produce expected output file: %s", test_name, temp_filename)
	}

	// load the actual ouptput
	actual_output, err := ioutil.ReadFile(temp_filename)
	if err != nil {
		t.Errorf("%s could not load actual output from file: %s", test_name, temp_filename)
	}

	// load the prepared output
	prepared_output, err := ioutil.ReadFile(prepared_output_filename)
	if err != nil {
		t.Errorf("%s could not load prepared output from file: %s", test_name, prepared_output_filename)
	}

	actual_output_str := string(actual_output)
	prepared_output_str := string(prepared_output)

	// check that the prepared output and actual output match
	if actual_output_str != prepared_output_str {
		t.Errorf("%s actual content does not match prepared content", test_name)
	}

	// delete the temporary file
	syscall.Unlink(temp_filename)
}

func Test1(t *testing.T) {
	process_and_compare(
		t,
		"test1",
		"tf",
		"tests/test1_input.md",
		"tests/test1_input.tf",
		"tests/test1_target.tf")
}

func Test2(t *testing.T) {
	process_and_compare(
		t,
		"test2",
		"tf",
		"tests/test2_input.md",
		"tests/test2_input.tf",
		"tests/test2_target.tf")
}

func Test3(t *testing.T) {
	process_and_compare(
		t,
		"test3",
		"tf",
		"tests/test3_input.md",
		"tests/test3_input.tf",
		"tests/test3_target.tf")
}

func Test4(t *testing.T) {
	process_and_compare(
		t,
		"test4",
		"tf",
		"tests/test4_input.md",
		"tests/test4_input.tf",
		"tests/test4_target.tf")
}

func Test5(t *testing.T) {
	const test_name string = "test5"
	var test5_filename = "tests/test5_input.md"
	var test5_target = []string{}

	test5_target = append(test5_target, "tf")
	test5_target = append(test5_target, "tfvars")
	test5_target = append(test5_target, "python")

	var languages = scan_md_languages(test5_filename)

	if !unordered_lists_are_equal(&languages, &test5_target) {
		// print_string_list("test5_target = ", test5_target)
		// print_string_list("languages    = ", languages)
		fmt.Printf("test5_target = %#v\n", test5_target)
		fmt.Printf("languages    = %#v\n", languages)
		t.Errorf("%s returned language list does not match target language list", test_name)
	}
}
