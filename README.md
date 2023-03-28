# terraform-training
![terraform-training](https://user-images.githubusercontent.com/42545233/123446215-a8e80f00-d5a6-11eb-87d0-be359b2641d8.png)

> # :warning: This material is archived and will not receive further updates. For the latest information on Terraform, it is advised to look at Terraform's [documentation](https://developer.hashicorp.com/terraform/docs) and [tutorials](https://developer.hashicorp.com/terraform/tutorials).

Learn to write terraform by actually writing code (urm, i mean templates).

If wanting to read through general terraform docs refer to here: https://www.terraform.io/docs/index.html

# What is terraform-training?

The goal of this repo is to get you familiar with writing terraform templates. Currently the focus is on using AWS as your cloud provider. Each excercise's goal is to introduce you to a singular concept within terraform. Will start out super simple and grow more complex as you do more excercies.

Keep in mind that each excercise is designed specifically to be solved with only a couple lines of code, so keep your solutions lean and simple.

# Exercise Layout

`./excercies/exercise_#`                       : Root Folder.   
`./excercies/exercise_#/README.md`             : Description of exercise.   
`./excercies/exercise_#/excercise_#_test.go`   : Unit Test written in GO.   
`./excercies/excercise_#/terraform/*`          : terraform resources directory.   

# Install Dependenecies

## Install Go

We recommend using Go `1.16.2` because we've seen issues installing Go packages with versions above this.

- **Manual Install:** download the appropriate installer/archive from https://golang.org/dl/#go1.16.2 and follow the remaining steps at https://golang.org/doc/install
- **Homebrew:** `brew install go`

## Install Terraform
- `. ./scripts/bootstrap` (NOTE: will install terraform binary to `/usr/local/bin/`)

# Usage
- `cd exercises/exercise_1`
- Read the README
- `inv test`
- Fix the issue by updating the templates in `exercises/exercise_1/terraform`
