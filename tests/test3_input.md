# md2tf Test File 2

Here is some Terraform code:

```tf
resource provider_element1 element1 {
    element1_variable = "element1_value"
}
```

This is not Terraform code as it does not have the language identifier:


```
resource provider_element2 element2 {
    element2_variable = "element2_value"
}
```

Here is a block that is commented out at the Markdown level; it should not included as active Terraform code; despite the correct language identideir:


<!-- And now we move back to Terraform code:

```tf
resource provider_element3 element3 {
    element3_variable = "element3_value"
}
```
-->

## Comment Ending Not on Line Start

<!---

We're in a comment, with an otherwise valid TF Block, but with the comment end not starting at the beginning of a new line:

```tf
resource provider_element4 element4 {
    element4_variable = "element4_value"
}
``` -->

## Last Section
Finally let's introduce one more Terraform block:
```tf
resource provider_element5 element5 {
    element5_variable = "element5_value"
}
```
We're doing this with little whitespace.
## Finish

File test2_input.md finished.
