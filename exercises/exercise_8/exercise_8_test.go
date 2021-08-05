package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExercise8Test(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"create_bucket": true,
		},
		VarFiles: []string{},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedText := []string{"th-exercise-eight"}
	actualTextExample := terraform.OutputList(t, terraformOptions, "buckets")
	assert.Equal(t, expectedText, actualTextExample)

	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"create_bucket": false,
		},
		NoColor: true,
	})
	terraform.InitAndApply(t, terraformOptions)
	actualTextExample = terraform.OutputList(t, terraformOptions, "buckets")
	expectedText = []string{}

	assert.Equal(t, expectedText, actualTextExample)
}

