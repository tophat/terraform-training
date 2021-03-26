# Excercise 6: Create Terraform module

The goal of Excercise 6 is to create terraform module that you can reuse accross different projects. This terraform module creates a S3 Bucket. The test passes is an a variable `s3_bucket_name` that should be used for the name of the bucket (as its used for assertion).

```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/gruntwork-io/terratest/modules/random
go get github.com/gruntwork-io/terratest/modules/aws
go test -v -run excercise_6_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/modules/develop/index.html
https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
