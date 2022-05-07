# lets create a bucket
resource "aws_s3_bucket" "my_bucket" {
  bucket = "flugel-terraform-test"
  acl    = "private"
  tags = {
    Name        = "My Flugel test bucket"
    Environment = "Test"
  }
}

resource "local_file" "test1" {
  content  = formatdate("YYYY-MM-DD'T'hh:mm:ssZZZZZ",timestamp())
  filename = "test1.txt"
}

resource "local_file" "test2" {
  content  = formatdate("YYYY-MM-DD'T'hh:mm:ssZZZZZ",timestamp())
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
