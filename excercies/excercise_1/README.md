# Excercise 1: Input Variables (pt 1)

Input Variables are used in terraform to pass variables into a terraform stack. These can be set via the cli with `--var` or using a `.tfvars` file. Running excercise_1 you will see that the input variable is failing the assert. Update the code in `./terraform` to set the input variable to the correct value

Fix this test by editing the input var file in tfvars.

```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v -run excercise_1_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/values/variables.html