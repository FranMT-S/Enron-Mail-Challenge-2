resource "aws_security_group" "sg_allow_ssh"{
  name        = "sg_allow_ssh"
  description = "ssh security group"
  vpc_id      = aws_vpc.vpc_mails.id
  
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}