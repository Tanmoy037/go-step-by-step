provider "aws" {
    region     = "${var.region}"
    access_key = "${var.access_key}"
    secret_key = "${var.secret_key}"
}
resource "aws_db_instance" "my_instance" {
  engine = "mysql"
  identifier = "myrdsinstance"
  instance_class = "db.t2.micro"
  allocated_storage = 20
  username = "admin"
  password = "admin123"
  skip_final_snapshot = true
  publicly_accessible =  true
}
