output "instance_id" {
  value  = aws_instance.ec2-mails.id
}

output "instance_public_ip" {
  value =  aws_instance.ec2-mails.public_ip
}
