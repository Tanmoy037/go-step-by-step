resource "aws_rds_instance" "my_instance" {
  name = "my-rds-instance"
  engine = "mysql"
  instance_class = "db.t2.micro"
  allocated_storage = 5
  username = "admin"
  password = "admin123"
}