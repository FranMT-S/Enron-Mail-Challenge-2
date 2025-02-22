resource "aws_security_group" "sg_allow_api" {
  name        = "sg_allow_api"
  description = "allow connect api"

  ingress {
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]  
  }

}

