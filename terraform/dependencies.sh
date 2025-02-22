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