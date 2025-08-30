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
aws eks update-kubeconfig --name v0
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
Edited podmakerv0_types.go - run above make generate
#
Added/removed CRD fields - - run above make generate
#
make manifests (generate crds)
#
internal/controller/podmakerv0_controller.go (implement controller logic here)
#
kubectl apply -f config/crd/bases/z0.v0.com_podmakerv0s.yaml
#
bash-5.2# pwd
#
/root/operatorv0/v3
#
bash-5.2# kubectl apply -f config/crd/bases/z0.v0.com_podmakerv0s.yaml
#
customresourcedefinition.apiextensions.k8s.io/podmakerv0s.z0.v0.com created
#
bash-5.2# go run ./cmd
#
Compiles the code (including all imports and dependencies)
#
ctrl + c (if on terminal operator stops no new reconcilation old ones remain)
#
Starts operator
#
Initializes the controller-runtime manager
#
Registers custom controller (PodMakerv0Reconciler)
#
Starts health probes and metrics endpoints
#
Begins watching for PodMakerv0 resources in the cluster
#
kubectl apply -f v3/config/samples/zv0.yml
#
Operator(creates pod) - A packaged application that includes one or more controllers
#
Controller - The actual logic that watches resources and takes action (like Pod creation)
