variable "aws_access_key" {
  description = "Your AWS access key"
  default = env("AWS_ACCESS_KEY_ID")
}

variable "aws_secret_key" {
  description = "Your AWS secret key"
  default = env("AWS_SECRET_ACCESS_KEY")
}


variable "region" {
  description = "The AWS region"
  default = "ap-south-1"
}

# export AWS_ACCESS_KEY_ID=my-access-key-id
# export AWS_SECRET_ACCESS_KEY=my-secret-access-key
