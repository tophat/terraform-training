package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExcercise8Test(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"create_bucket": true,
		},
		VarFiles: []string{"varfile.tfvars"},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedLength := 1
	actualResult := terraform.OutputList(t, terraformOptions, "buckets")
	assert.Equal(t, expectedLength, len(actualResult))

	terraformOptions2 := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"create_bucket": false,
		},
		VarFiles: []string{"varfile.tfvars"},
		NoColor: true,
	})

	terraform.InitAndApply(t, terraformOptions2)
	expectedLength2 := 0
	actualResult2 := terraform.OutputList(t, terraformOptions, "buckets")
	assert.Equal(t, expectedLength2, len(actualResult2))
}

