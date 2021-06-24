
terraform {
  required_version = ">=0.13.2, <0.15"
}

module "my_module" {
    source = "./my_bucket"
    input_one = "food"
    input_two = false
    input_three = 1
    input_four = ["asdasd", "Asdasd"]
    input_five = {"foo": "Bar"}
}