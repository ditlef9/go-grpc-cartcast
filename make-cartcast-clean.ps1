# clean_greet.ps1
# Script to clean generated protobuf files for greet

# Configuration variables
$protoDir = "proto"
$greetDir = "greet"

# Check if the directory exists
if (-Not (Test-Path "$greetDir/$protoDir")) {
    Write-Host "`e[31mDirectory 'greet/$protoDir' doesn't exist`e[0m" -ForegroundColor Red
    exit 1
}

# Remove generated .pb.go files
$pbFiles = Get-ChildItem "$greetDir/$protoDir/*.pb.go"
if ($pbFiles.Count -eq 0) {
    Write-Host "`e[33mNo .pb.go files found in 'greet/$protoDir'.`e[0m" -ForegroundColor Yellow
} else {
    Remove-Item "$greetDir/$protoDir/*.pb.go" -Force
    Write-Host "`e[32mCleaned generated protobuf files.`e[0m" -ForegroundColor Green
}
