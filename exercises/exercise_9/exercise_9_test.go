package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExcercise7Test(t *testing.T) {
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

	outputOne := terraform.Output(t, terraformOptions, "input_one")
	assert.Equal(t, "food", outputOne)
	outputTwo := terraform.Output(t, terraformOptions, "input_two")
	assert.Equal(t, "false", outputTwo)
	outputThree := terraform.Output(t, terraformOptions, "input_three")
	assert.Equal(t, "1", outputThree)

	expectedOutputFour := []string{"asdasd", "Asdasd"}
	outputFour := terraform.OutputList(t, terraformOptions, "input_four")
	assert.Equal(t, expectedOutputFour, outputFour)

	expectedOutputFive := map[string]string{"foo": "Bar"}
	outputFive := terraform.OutputMap(t, terraformOptions, "input_five")
	assert.Equal(t, expectedOutputFive, outputFive)
}

