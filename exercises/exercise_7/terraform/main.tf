
terraform {
  required_version = ">=0.13.2, <0.15"
}

resource "aws_s3_bucket" "s3_bucket" {
  bucket = "th-exercise"
}
