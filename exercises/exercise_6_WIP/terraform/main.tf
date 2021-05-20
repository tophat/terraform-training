
terraform {
  required_version = ">=0.13.2, <0.15"
}

module "custom_s3_bucket" {
  source = "./modules/s3_bucket"
}
