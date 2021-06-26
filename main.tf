# lets create a bucket
resource "aws_s3_bucket" "my_bucket" {
  bucket = "flugel-terraform-test"
  acl    = "private"
  tags = {
    Name        = "My Flugel test bucket"
    Environment = "Test"
  }
}
