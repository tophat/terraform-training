# Exercise 1: Input Variables (pt 1)

Input Variables are used in terraform to pass variables into a Terraform stack. These can be set via the cli with `--var` or using a `.tfvars` file.
Running exercise_1 you will see that the input variable is failing the assertion.
Update the code in `./terraform` to set the input variable to the correct value

Fix this test by editing the input var file `varfile.tfvars`.

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v excercise_1_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/values/variables.html
