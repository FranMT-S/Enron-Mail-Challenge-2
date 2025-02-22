resource "aws_security_group" "sg_allow_zinc" {
  name        = "sg_allow_zinc"
  description = "allow connect zinc client security group"

  ingress {
    from_port   = 4080
    to_port     = 4080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] 
  }

}
