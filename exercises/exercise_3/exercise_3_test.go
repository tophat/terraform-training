package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExercise3Test(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"random_input_string": "test",
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedOutputs := make(map[string]interface{})
	expectedOutputs["random_output_string"] = "test"
	actualTextExample := terraform.OutputAll(t, terraformOptions)
	assert.Equal(t, expectedOutputs, actualTextExample, "Assert onOutput variable `random_output_string` ")

	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"random_input_string": "test2",
		},
		NoColor: true,
	})
	terraform.InitAndApply(t, terraformOptions)
	expectedOutputs = make(map[string]interface{})
	expectedOutputs["random_output_string"] = "test2"
	actualTextExample = terraform.OutputAll(t, terraformOptions)
	assert.Equal(t, expectedOutputs, actualTextExample, "Assert onOutput variable `random_output_string` ")

}
