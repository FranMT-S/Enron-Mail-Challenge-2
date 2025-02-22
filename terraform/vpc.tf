resource "aws_vpc" "vpc_mails" {
  cidr_block = "10.0.0.0/22"
  enable_dns_support = true
  enable_dns_hostnames = true
  tags = {
    Name = "vpc_mails"
  }
}

resource "aws_subnet" "subnet_mails_client" {
  vpc_id            = aws_vpc.vpc_mails.id
  cidr_block        = "10.0.0.0/24"
  availability_zone = var.availability_zone
}
