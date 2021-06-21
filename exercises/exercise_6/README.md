# Excercise 6: Create Terraform module

The goal of Excercise 6 is to CREATE a custom terraform module that you can reuse accross different projects. This terraform module should create a S3 Bucket.  You can use the folder in `terraform/module` to add the module.

The test passes if a s3_bucket with the name given from the  variable `s3_bucket_name` is created via a module.

Your Module should be named `my_bucket` and have an output variable called `aws_s3_bucket`. The test will assert the s3 bucket is created 

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
go test -v excercise_6_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/modules/develop/index.html
https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
