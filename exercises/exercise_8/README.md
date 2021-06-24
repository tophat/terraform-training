# Exercise 8: Conditional Resources 

While `count` as used in Excercise 7 is used to launch multiple of a given resource, it can also used in Terraform dynamically launch or not set a resource. When wanting to create a resource that may or may not exist , you can utilize the count variable to be `0` or `1` based on your logic. This is how you can achieve "If X, then created Y" logic. 

Excercise 8 is currently setup to launch a single `s3_bucket` and has in an boolean input variable called `launch_bucket` which is a boolean. Use this input var to dynamically set the creation of the s3 bucket. 

The test will assert that when  `launch_bucket` is true, the bucket is created and when `launch_bucket` is `false` no bucket is created


## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v excercise_8_test.go
```

# Note

The one thing to keep in my when using this conditional logic, is it does change how you can refer to a resource.  Previously, after creating a resource you were able to refer to it by using the `<RESOURCE>.<NAME>` syntax, if using the count variable this changes to `<RESOURCE>.<NAME>[<INDEX>]`. A very common mistake is using count to dynamically set a resource, then forgetting to update the references to the resource to account for this change of syntax. 

```
resource "aws_instance" "example" {
  ami           = "ami-abc123"
}
```
`aws_instance.example.id` : id of instance.


This syntax changes when using the count variable, as the resource now became a list of objects, instead of a single object.
```
resource "aws_instance" "example" {
  count = 1
  ami           = "ami-abc123"
}
```
`aws_instance.example[*].id`: returns a list of all of the ids of each of the instances.
`aws_instance.example[0].id`: returns just the id of the first instance.

