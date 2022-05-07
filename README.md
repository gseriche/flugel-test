# IaC Testing

This is a simple example to create a bucket called "flugel-terraform-test" and
 upload two files with the timestamp of execution. You can change the name of
 the bucket inside of the file "main.tf".

## What I need to use this code

This Terraform script was executed with the version 1.1.9 and two providers:

-aws v4.13.0
-local v2.2.2

Also you need to create an AWS account to get access and give the possibility
to create the bucket and upload the file.
The account must have this permission:

- s3:GetBucketLocation
- s3:ListBucket
- s3:GetObject
- s3:PutObject

After that you should get the Access Key to fill the file called envrc, this file
has the environment variable to test the terraform in local.

~~~bash
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_DEFAULT_REGION=us-west-2
export TF_VAR_access_key=${AWS_ACCESS_KEY_ID}
export TF_VAR_secret_key=${AWS_SECRET_ACCESS_KEY}
~~~

## How to use this in the console (local)

In your favorite console you should load first the providers

~~~console
terraform init
~~~

Then execute the plan

~~~console
terraform plan
~~~

You must check if the plan its ok with this will do, create a bucket, upload
two files called test1.txt and test2.txt to the bucket, then when you see
everything its ok, execute the apply plan.

~~~console
terraform apply
~~~

You should write yes when its prompted, to apply the plan.

## How to test the code (local)

The test its located inside the test folder and you must have in your local Go
1.16 first, after that you must execute this command inside the test folder:

~~~console
go get -v -t -d && go mod tidy
~~~

Then to test the code you must execute only this:

~~~console
go test
~~~

## How to test the code in GitHub Action

You must create 4 secrets environment variables called:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_DEFAULT_REGION
- SSH_KEY

The three first variables must have the same condition explained in the first
paragraph, after that you should create a SSH_KEY variable having your private
key and get access to the repository, this variables was used inside the folder
".github/workflows" in the file called "terratest-action.yml"
