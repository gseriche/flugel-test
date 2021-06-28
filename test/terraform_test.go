package test

import (
  "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "testing"
)

func TestTerraformFlugel(t *testing.T) {
  t.Parallel()
  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../.",
  })

  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)

  expectedBucketName := "flugel-terraform-test"

  bucketRegion := terraform.Output(t, terraformOptions, "bucket_region")
  aws.AssertS3BucketExistsE(t, bucketRegion, expectedBucketName)
  aws.GetS3ObjectContents(t, bucketRegion, expectedBucketName, "test1.txt")
  aws.GetS3ObjectContentsE(t, bucketRegion, expectedBucketName, "test2.txt")
}
