resource "aws_security_group" "sg_allow_client" {
  name        = "sg_allow_client"
  description = "allow connect to api"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]  
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]  
  }

}

resource "aws_security_group" "sg_egress_allow_all" {
  name        = "sg_egress_allow_all"
  description = "allow all egress trafic "

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"  
    cidr_blocks = ["0.0.0.0/0"]
  }

}
