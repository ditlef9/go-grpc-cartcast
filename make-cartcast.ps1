# make-cartcast.ps1
# Script to generate Pbs and build the project

# Configuration variables
$binDir = "bin"
$protoDir = "proto"
$serverDir = "server"
$clientDir = "client"

# Ensure project directory exists
$projectDir = "cartcast"
if (-Not (Test-Path $projectDir)) {
    Write-Host "Error 1: `e[31mDirectory '$projectDir' doesn't exist`e[0m" -ForegroundColor Red
    exit 1
}
Write-Host "Info: Project directory '$projectDir' found." -ForegroundColor Green

# Get Go package name from go.mod
$goModPath = "go.mod"
if (-Not (Test-Path $goModPath)) {
    Write-Host "Error 2: `e[31mFile '$goModPath' doesn't exist`e[0m" -ForegroundColor Red
    exit 1
}
$package = (Get-Content $goModPath -Head 1).Split(" ")[1]
Write-Host "Info: Go package name is '$package'." -ForegroundColor Green

# Check if proto directory exists
$protoPath = "$projectDir/$protoDir"
if (-Not (Test-Path $protoPath)) {
    Write-Host "Error 3: `e[31mProto directory '$protoPath' doesn't exist`e[0m" -ForegroundColor Red
    exit 1
}
Write-Host "Info: Proto directory '$protoPath' found." -ForegroundColor Green

# Run protoc to generate .pb.go files for project
Write-Host "Info: Running protoc to generate .pb.go files from '$protoPath'..." -ForegroundColor Green
protoc -I $protoPath --go_opt=module=$package --go_out=. --go-grpc_opt=module=$package --go-grpc_out=. "$protoPath/*.proto"
if ($LASTEXITCODE -ne 0) {
    Write-Host "`e[31mProtoc failed to generate pb files.`e[0m" -ForegroundColor Red
    exit 1
}
Write-Host "Info: Protoc completed successfully." -ForegroundColor Green

# Build the server binary
Write-Host "Info: Building the server binary..." -ForegroundColor Green
$serverBinPath = "$binDir/cartcast/$serverDir.exe"
go build -o $serverBinPath "./$projectDir/$serverDir"
if ($LASTEXITCODE -ne 0) {
    Write-Host "`e[31mFailed to build the server binary.`e[0m" -ForegroundColor Red
    exit 1
}
Write-Host "Info: Server binary built successfully at '$serverBinPath'." -ForegroundColor Green

# Build the client binary
Write-Host "Info: Building the client binary..." -ForegroundColor Green
$clientBinPath = "$binDir/cartcast/$clientDir.exe"
go build -o $clientBinPath "./$projectDir/$clientDir"
if ($LASTEXITCODE -ne 0) {
    Write-Host "`e[31mFailed to build the client binary.`e[0m" -ForegroundColor Red
    exit 1
}
Write-Host "Info: Client binary built successfully at '$clientBinPath'." -ForegroundColor Green

Write-Host "`e[32mBuild successful!`e[0m" -ForegroundColor Green
