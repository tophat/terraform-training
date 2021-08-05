package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExercise10Test(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
		},
		VarFiles: []string{},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedOutputs := map[string]string{
		"one":"two",
		"two":"2",
		"three":"[string string]",
		"four": "map[five:5 six:foobar]",
	}


	outputFive := terraform.OutputMap(t, terraformOptions, "input")
	assert.Equal(t, expectedOutputs, outputFive)
}
