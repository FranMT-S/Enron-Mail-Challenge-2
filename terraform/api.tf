# resource "aws_instance" "api_instance" {
#   ami             = var.ami  
#   instance_type   = var.instance_type
#   key_name        = var.mails-key 
#   security_groups = [aws_security_group.sg_app.name,aws_security_group.sg_app.name]
  
#   tags = {
#     Name = "api_instance"
#   }
# }