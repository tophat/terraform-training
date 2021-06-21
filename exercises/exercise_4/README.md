# Excercise 4: AWS Create S3 Resources

The goal of Excercise 4 is a terraform template that creates a S3 Bucket. The test passes in an a variable `s3_bucket_name` that should be used for the name of the bucket (as its used for assertion). 


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