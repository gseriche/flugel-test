# lets create a bucket
resource "aws_s3_bucket" "my_bucket" {
  bucket = "flugel-terraform-test"
  tags = {
    Name        = "My Flugel test bucket"
    Environment = "Test"
  }
}

resource "aws_s3_bucket_acl" "my_bucket" {
  bucket = aws_s3_bucket.my_bucket
  acl    = "private"
}

resource "local_file" "test1" {
  content  = formatdate("YYYY-MM-DD'T'hh:mm:ssZ",timestamp())
  filename = "test1.txt"
}

resource "local_file" "test2" {
  content  = formatdate("YYYY-MM-DD'T'hh:mm:ssZ",timestamp())
  filename = "test2.txt"
}

resource "aws_s3_bucket_object" "copy_test1" {
  bucket = aws_s3_bucket.my_bucket.id
  key    = local_file.test1.filename
  source = local_file.test1.filename
}

resource "aws_s3_bucket_object" "copy_test2" {
  bucket = aws_s3_bucket.my_bucket.id
  key    = local_file.test2.filename
  source = local_file.test2.filename
}
