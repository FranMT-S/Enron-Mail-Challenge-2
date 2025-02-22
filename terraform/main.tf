
resource "aws_instance" "ec2-mails" {
  ami             = var.ami  
  instance_type   = var.instance_type
  key_name        = var.key_name
  subnet_id     = aws_subnet.subnet_mails_client.id 
  security_groups = [
    aws_security_group.sg_allow_zinc.name,
    aws_security_group.sg_allow_ssh.name,
    aws_security_group.sg_allow_api.name,
    aws_security_group.sg_allow_client.name,
    aws_security_group.sg_egress_allow_all.name
  ]
  
  user_data = <<-EOF
    #!/bin/bash

    sudo apt update -y && sudo apt upgrade -y

    #  git
    sudo apt install -y git

    # docker
    sudo apt install -y docker.io
    sudo systemctl enable docker
    sudo systemctl start docker
    sudo usermod -aG docker ubuntu
    newgrp docker
    
    # golan

    sudo apt remove golang-go
    sudo rm -rf /usr/local/go
    wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
    sudo tar -C /usr/local -xvzf go1.21.0.linux-amd64.tar.gz
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
    source ~/.bashrc
    go version
  EOF
  
  tags = {
    Name = "ec2-mails"
  }

}
