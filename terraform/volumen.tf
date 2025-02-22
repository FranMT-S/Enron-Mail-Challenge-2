resource "aws_ebs_volume" "mails_volumen" {
  availability_zone = var.availability_zone
  size              = 20
  type              = "gp2"
  encrypted         = false
  tags = {
    Name = "mails_volumen"
  }
}

resource "aws_volume_attachment" "mails_ebs_attachment" {
  device_name = "/dev/xvdf"
  volume_id   = aws_ebs_volume.mails_volumen.id
  instance_id = aws_instance.ec2-mails.id
}