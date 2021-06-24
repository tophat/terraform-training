
terraform {
  required_version = ">=0.13.2, <0.15"
}

module "my_module" {
    source = "./my_module"
    input = {
      "one": "two",
      "two": 2,
      "three": [ "string", "string" ],
      "four": {
          "five": 5
          "six": "foobar"
      }
    }
}