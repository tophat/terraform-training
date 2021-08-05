# Exercise 10: Variable Typing (Part 2)

Expanding on Exercise 9, we will now look at complex typing in the form of maps and lists with multiple types supported. When using `type = map` and `type = list` like we did in exercise 9, there is an implication that all values stored in these objects share the same type. (a map of all strings, a list of all numbers, etc)

Following a similar format, `./terraform/my_module` is set up with a single input variable `input` with incorrect typing defined. Based on the map of various types being passed in to the module in `./terraform/main.tf` fix the typing of this complex map being passed in.

Get this exercise passing by updating the type in `./terraform/my_module/variables.tf` to its correct values based on what's currently being passed to it. Do not change the values being passed to the module, you should only have to update the type.

```
module "my_module" {
    source = "./my_module"
    input = {
      "one": "two",
      "two": 2,
      "three": [ "string", "string" ],
      "four": {
          "five": 5
          "six": "foobar"
      }
    }
}
```

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v exercise_10_test.go
```


# Additional Readings
https://www.terraform.io/docs/language/expressions/type-constraints.html#complex-types
