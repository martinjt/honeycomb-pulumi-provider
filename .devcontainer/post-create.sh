#!/bin/zsh
set -e

curl -fsSL https://get.pulumi.com | sh

wget https://github.com/pulumi/pulumictl/releases/download/v0.0.45/pulumictl-v0.0.45-linux-amd64.tar.gz
tar xf pulumictl-v0.0.45-linux-amd64.tar.gz
sudo mv pulumictl ~/.pulumi/bin
sudo chmod a+x ~/.pulumi/bin/pulumictl
rm -rf pulumictl-v0.0.45-linux-amd64.tar.gz

export PATH=$PATH:~/.pulumi/bin