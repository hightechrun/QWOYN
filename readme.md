# Cosmic Horizon Network

# Installing `qwoynd`

## Hardware Requirements

Here are the minimal hardware configs required for running a validator/sentry node

- 16GB RAM
- 4vCPUs
- 500gb Disk space

## Software Requirements

- Ubuntu 18.04 or higher
- [Go v1.18.1](https://golang.org/doc/install) or higher

# Install `qwoynd`, Generate Wallet and Start your Node

Follow the steps below to compile `qwoynd`.

## Install Go version 1.18.1

```
sudo apt update
sudo apt install build-essential jq wget git -y

wget https://dl.google.com/go/go1.18.1.linux-amd64.tar.gz
tar -xvf go1.18.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

Now add go to your bashrc -

```
echo "" >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

# use this new bashrc configuration
source ~/.bashrc
```

## Build `qwoynd`

These steps install the `qwoynd` binary.

```
# Clone the Repo
git clone https://github.com/cosmic-horizon/QWOYN.git

# Install `Qwoynd`
cd QWOYN
git checkout v1.0.0-beta
make install

# check version (1.0.0-beta)
qwoynd version
```
