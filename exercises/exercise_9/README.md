# Exercise 9: Variable Typing (Part 1)

While for the most part Typing of variables is rather intuitive, there are some unique properties to how Terraform's template language defines them. The goal of Exercise 9 + 10 is to explore some of these behaviours.

Generally speaking, the only time you will run into issues with Variable Typing is when defining new input variables of a new module.

With that in mind, Exercise is set up with a module called `my_module` defined under `./terraform/my_module`. This Module currently takes in 5 input variables (input_one, input_two, input_three, input_four, input_five) all of which currently have values being passed to them via `terraform/main.tf`. In its current state, this test will fail because the types of each of these input variables are currently misconfigured.

Get this exercise passing by updating the types in `./terraform/my_module/variables.tf` to their correct values based on what's currently being passed to them. Do not change the values being passed to the module, you should only have to update the types.

```
module "my_module" {
    source = "./my_module"
    input_one = "food"
    input_two = false
    input_three = 1
    input_four = ["asdasd", "Asdasd"]
    input_five = {"foo": "Bar"}
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
go test -v exercise_9_test.go
```


# Additional Readings
https://www.terraform.io/docs/language/expressions/types.html
