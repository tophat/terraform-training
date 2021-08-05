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
	slack_webhook_url := fmt.Sprintf("https://hooks.slack.com/services/AAA/BBB/CCC")
	slack_username := fmt.Sprintf("reporter")
	sns_topic_name :=  fmt.Sprintf("slack-topic")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		TerraformDir: "./terraform/",
		Vars: map[string]interface{}{
			"slack_channel_name": slack_channel_name,
			"slack_webhook_url": slack_webhook_url,
			"slack_username": slack_username,
			"sns_topic_name": sns_topic_name,
		},
		VarFiles: []string{"varfile.tfvars"},
		NoColor:  true,
	})

	defer terraform.Destroy(t, terraformOptions)
	output := terraform.InitAndApply(t, terraformOptions)
	assert.Contains(t, output, "module.notify_slack", "Assert on using module `notify_slack`")
}
