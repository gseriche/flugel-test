package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformFlugel(t *testing.T) {
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../.",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	expectedBucketName := "flugel-terraform-test"
	expectedContent := time.Now().Format(time.RFC3339)

	bucketRegion := terraform.Output(t, terraformOptions, "bucket_region")
	aws.AssertS3BucketExistsE(t, bucketRegion, expectedBucketName)

	actualContentA := aws.GetS3ObjectContents(t, bucketRegion, expectedBucketName, "test1.txt")
	assert.Equal(t, expectedContent, actualContentA)
	actualContentB := aws.GetS3ObjectContents(t, bucketRegion, expectedBucketName, "test2.txt")
	assert.Equal(t, expectedContent, actualContentB)
}
