resource "aws_vpc" "vpc_mails" {
  cidr_block = "10.0.0.0/22"
  enable_dns_support = true
  enable_dns_hostnames = true
  tags = {
    Name = "vpc_mails"
  }
}

resource "aws_internet_gateway" "gateway" {
  vpc_id = aws_vpc.vpc_mails.id
}


resource "aws_subnet" "subnet_mails_client" {
  vpc_id            = aws_vpc.vpc_mails.id
  cidr_block        = "10.0.0.0/24"
  availability_zone = var.availability_zone
  map_public_ip_on_launch = true
}

resource "aws_route_table" "all_trafic" {
  vpc_id = aws_vpc.vpc_mails.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gateway.id
  }
  tags = {
    Name = "route-table-mails"
  }
}

resource "aws_route_table_association" "subnet_mails_client_association" {
  subnet_id      = aws_subnet.subnet_mails_client.id
  route_table_id = aws_route_table.all_trafic.id
}