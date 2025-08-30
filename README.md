# operatorv0
golang

# run vs in wsl
code .

# create go module
go mod init v0
#
go run .

# standard lib go
https://pkg.go.dev/std

# launch vs code in wsl
/mnt/c/Users/lenovo/Downloads/VSCode-win32-x64-1.87.2/code.exe "$@"

# 
$env:GIT_SSL_NO_VERIFY = "true"
#
sudo apt update
#
sudo apt install golang-go
#
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
#
chmod +x kubectl
#
sudo mv kubectl /usr/local/bin/
#
curl -LO https://github.com/operator-framework/operator-sdk/releases/latest/download/operator-sdk_linux_amd64
#
chmod +x operator-sdk_linux_amd64
#
mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk
#
go version
#
kubectl version --client
#
operator-sdk version
#
apt update
#
apt install docker.io



