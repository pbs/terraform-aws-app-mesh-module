package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testAppMesh(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "id")
	name := terraform.Output(t, terraformOptions, "name")
	arn := terraform.Output(t, terraformOptions, "arn")

	region := getAWSRegion(t)
	accountID := getAWSAccountID(t)

	expectedID := fmt.Sprintf("example-tf-app-mesh-%s", variant)
	expectedName := expectedID
	expectedARN := fmt.Sprintf("arn:aws:appmesh:%s:%s:mesh/%s", region, accountID, expectedName)

	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedName, name)
	assert.Equal(t, expectedARN, arn)
}
