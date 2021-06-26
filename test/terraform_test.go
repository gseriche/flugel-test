package test

import (
  "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "log"
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
  bucketExist := aws.AssertS3BucketExistsE(t, bucketRegion, expectedBucketName)
  if bucketExist != nil {
    log.Println("The bucket doesn't exist")
  }else {
    log.Println("The bucket exists")
  }

  testFileOne, errOne := aws.GetS3ObjectContentsE(t, bucketRegion, expectedBucketName, "test1.txt")
  testFileTwo, errTwo := aws.GetS3ObjectContentsE(t, bucketRegion, expectedBucketName, "test2.txt")

  if errOne != nil {
    log.Println("The File test1 doesn't exist")
  }else {
    log.Println(testFileOne)
    log.Println("The File test1 exists")
  }

  if errTwo != nil {
    log.Println("The File test2 doesn't exist")
  }else {
    log.Println(testFileTwo)
    log.Println("The File test2 exists")
  }

}
