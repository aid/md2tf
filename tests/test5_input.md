# md2tf Test File 5

## Language Tag Tests

Check for tf:

```tf
resource provider_element1 element1 {
    element1_variable = "element1_value"
}
```

Some tfvars, with spaces before:


```  tfvars
resource provider_element2 element2 {
    element2_variable = "element2_value"
}
```

A little python, with spaces after:

```python    
resource provider_element3 element3 {
    element3_variable = "element3_value"
}
```

A fenced block with some spaces in place of a language tab:

```   
resource provider_element3 element3 {
    element3_variable = "element3_value"
}
```

A fenced block with no language tab:

```   
resource provider_element3 element3 {
    element3_variable = "element3_value"
}
```

A fenced block with two languages tages (what should we do)?

``` this that
resource provider_element3 element3 {
    element3_variable = "element3_value"
}
```

## Finish

File test5_input.md finished.
