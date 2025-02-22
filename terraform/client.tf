# resource "aws_instance" "app_instance" {
#   ami             = var.ami  
#   instance_type   = var.instance_type
#   key_name        = var.mails-key 
#   security_groups = [aws_security_group.api.name]
  
#   tags = {
#     Name = "app_instance"
#   }
# }