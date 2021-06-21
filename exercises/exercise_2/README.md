# Exercise 2: Input Variables (pt 2)

Continuing on from exercise_1, we will now go about adding a new input variable. The test here will fail due to input variable not existing.

Fix this by adding the wanted input variable.

The input variable should be named `random_input_string` and you should give it a string value of `test`

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v excercise_2_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/values/variables.html
