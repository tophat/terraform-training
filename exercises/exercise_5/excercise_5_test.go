package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExcercise4Test(t *testing.T) {
	t.Parallel()

	slack_channel_name := fmt.Sprintf("my-slack-channel")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"slack_channel_name": slack_channel_name,
		},
		VarFiles: []string{"varfile.tfvars"},
		NoColor:  true,
	})

	defer terraform.Destroy(t, terraformOptions)
	output := terraform.InitAndApply(t, terraformOptions)
	assert.Contains(t, output, "module.notify_slack", "Assert on using module `notify_slack`")
}
