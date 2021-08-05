package test

import (
	"fmt"
	"strings"
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestExercise4Test(t *testing.T) {
	t.Parallel()

	s3_bucket_name := fmt.Sprintf("terraform-training-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"s3_bucket_name": s3_bucket_name,
		},
		VarFiles: []string{"varfile.tfvars"},
		NoColor: true,
	})


	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	aws.AssertS3BucketExists(t, "us-east-1", s3_bucket_name)
}
