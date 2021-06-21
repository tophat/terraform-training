# Exercise 5: AWS Create S3 Resources

The goal of Exercise 5 is to use a Terraform module. A module is a container for multiple resources that are used together. We will use the public terraform module called "notify-slack".

The test passes if the module is successfully referenced i.e. if terraform creates SNS and Lambda resources.

Note: do NOT create the resources manually.

The Test passes in 4 input variables `slack_channel_name`, `slack_webhook_url`, `slack_username`, `sns_topic_name`

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
