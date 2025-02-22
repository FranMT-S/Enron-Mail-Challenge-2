
resource "aws_instance" "ec2-mails" {
  ami             = var.ami  
  instance_type   = var.instance_type
  key_name        = var.key_name
  subnet_id     = aws_subnet.subnet_mails_client.id 
  associate_public_ip_address = true

  vpc_security_group_ids  = [
    aws_security_group.sg_allow_zinc.id,
    aws_security_group.sg_allow_ssh.id,
    aws_security_group.sg_allow_api.id,
    aws_security_group.sg_allow_client.id,
    aws_security_group.sg_egress_allow_all.id
  ]
  
  user_data = file("dependencies.sh")

  tags = {
    Name = "ec2-mails"
  }

}
