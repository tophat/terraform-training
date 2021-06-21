# Exercise 4: AWS Create S3 Resources

The goal of Exercise 4 is to create a Terraform template that creates an S3 Bucket.
The test passes in a value for the variable `s3_bucket_name` that should be used for the name of the bucket (as it's used for assertion).


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
go test -v excercise_4_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/values/inputs.html
https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
