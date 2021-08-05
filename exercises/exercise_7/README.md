# Exercise 7: Count Meta Variable

The `count` meta variable can be used to launch multiple of a single resource without having to copy/paste resource blocks.
exercise_7 currently has an output setup that will return all buckets defined in a list of string.
The test expects for this to return 4 buckets with the name format of `th-exercise-0, th-exercise-1, th-exercise-2, th-exercise-3`

Fix this by defining 4 buckets without copy/pasting the same code.

## Invoke
```
inv test
```

## Manually
```
# Only need to be ran once for all tests
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v exercise_7_test.go
```

# Additional Readings
https://www.terraform.io/docs/language/meta-arguments/count.html
