package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExercise7Test(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{},
		VarFiles: []string{"varfile.tfvars"},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedText := []string{"th-exercise-0","th-exercise-1","th-exercise-2", "th-exercise-3" }
	actualTextExample := terraform.OutputList(t, terraformOptions, "buckets")
	assert.Equal(t, expectedText, actualTextExample)
}

