# Excercise 5: AWS Create S3 Resources

The goal of Excercise 5 is to use terraform module. A module is a container for multiple resources that are used together. We will use terraform module called "notify-slack". The test passes is on the module being added succesfully i.e. if terraform creates SNS and lambda.

We should have two input variable - Slack channel and webhook

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/gruntwork-io/terratest/modules/random
go get github.com/gruntwork-io/terratest/modules/aws
go test -v excercise_5_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/modules/develop/index.html
https://registry.terraform.io/modules/terraform-aws-modules/notify-slack/aws/latest
