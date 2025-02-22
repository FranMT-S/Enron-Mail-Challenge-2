resource "aws_security_group" "sg_allow_ssh"{
  name        = "sg_allow_ssh"
  description = "ssh security group"
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}