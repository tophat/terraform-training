# Exercise 3: Output Variables

As the inverse to Exercise 2, this test will fail due to a missing output variable.

The output variable should be named `random_output_string` and should be the value of the input variable `random_input_string`

Fix this by adding the wanted output variable needed and giving it the correct value

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v exercise_3_test.go
```


# Additional Readings
https://www.terraform.io/docs/language/values/outputs.html
