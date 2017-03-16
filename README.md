Cascade
=======

A consul cluster management shell

Set up the build on Ubuntu
============

### Install go and checkout source
1. install go, version should be 1.6.
```
curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
tar xvf go1.6.linux-amd64.tar 
sudo mv go /usr/local
```
2. append /usr/local/go/bin to your PATH
```
export PATH=$PATH:/usr/local/go/bin
```
3. setup the GOPATH to the directory where you are going to checkout the source, assume your home directory is /home/boundary
```
mkdir git/casdade
export GOPATH=/home/boundary/git/cascade
```
4. checkout the source 
```
go get github.com/boundary/cascade
```
### install fpm
```
sudo apt-get update
sudo apt-get install ruby-dev build-essential
sudo gem install fpm
```

### run build
```
cd /home/boundary/git/cascade/src/github.com/boundary/cascade
make
```
