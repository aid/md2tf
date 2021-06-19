# md2tf

Process Markdown files into Terraform .tf and .tfvars files.

## Overview

The `md2tf` tool allows you to create "Infrastructure as Documentation" - document your infrastrcture in expressive Markdown using titles, lists tables and any other element; but also including Terraform directives within Fenced Code Blocks with a language type of 'tf'.

Fenced Code Blocks start at the begging of a line with three backticks, following by a language tag with the code block terminated by a a further three backticks on a line of their own.

For example if we had a file `main.md`:

```
    # My Infastructure

    ## Router 1

    Here we define router 1 which acts as the gateway router for our network:

    ```tf
    module router1 {
        source = "./myrouter"
        name = "router1"
        loopback_ip = "10.1.1.1"
    }
    ```

```

The `md2tf` tool would convert the above file to `main.tf`.  The `main.tf` file would then be usable as part of your `terraform plan` or `terraform apply` commands.

## Line Numbers

Note the `md2tf` tool does not extract the Terraform code from the Mardown files and place just those entries into the `.tf` files.  

Instead, `md2tf` copies all lines of the `.md` file int the corresponding `.tf` file; while commenting out all of the Markdown file **except** for the Terraform content.

Through this operation; any line numbers referenced in error output of Terraform will be valid line numbers for the original Markdown source — allowing you to rapidly and accurately find the location of source of the error.


## Usage

Call this tool with a list of files and/or directories.  Only files with permitted extensions will be processed.  By default, these are `.md` files.  Any directories listed are recusively searched for files within that match the permitted extensions:

```
md2tf <filename>|<directory> ...
```



