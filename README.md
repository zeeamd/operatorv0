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
# Initializes a new Operator project
operator-sdk init --domain=v0.com --repo=pod-operator
# Sets the API domain for your custom resources. This will be part of the full API group (e.g., demo.v0.com/v1
# Sets the Go module path. This should match Git repo name or desired import path
# This generated all files in current dir and no pod-operator dir was created
operator-sdk create api --group=z0 --version=v1 --kind=PodMakerv0 --resource --controller
#
creates api/v1/ → where we define custom resource spec
#
controllers/ → where we write logic to create Pod
#
--group=demo: Sets API group to z0.v0.com
#
--version=v1: API version
#
--kind=PodMaker: Custom resource name
#
--resource: Generates the CRD definition
#
--controller: Creates the controller logic
#
add fields to api/v1/podmakerv0_types.go
#
PodMakerv0Spec struct, define the fields you custom resource should accept
#
regenerate Go code and update the CRD YAML files to reflect new spec
#
make generate
#
make manifests
#
internal/controller/podmakerv0_controller.go (implement controller logic here)
