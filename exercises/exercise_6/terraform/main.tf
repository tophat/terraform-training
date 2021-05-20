
terraform {
  required_version = ">=0.13.2, <0.15"
}

module "my_bucket" {
  source = "./module/custom_s3_bucket"
}
