# about.ps1
# Display info related to the build environment

# OS information
$osInfo = Get-ComputerInfo -Property OsVersion, OsArchitecture
$os = "{0} {1}" -f "Windows", ($osInfo | Format-Table -HideTableHeaders | Out-String).Trim()

# Shell information
$shellVersion = (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim()

# Protoc version
$protocVersion = (protoc --version)

# Go version
$goVersion = (go version)

# Go package (extract the first line from go.mod and split to get the package name)
$package = (Get-Content go.mod -Head 1).Split(" ")[1]

# OpenSSL version
$opensslVersion = (openssl version)

# Display the information
Write-Host "OS: $os"
Write-Host "Shell: powershell $shellVersion"
Write-Host "Protoc version: $protocVersion"
Write-Host "Go version: $goVersion"
Write-Host "Go package: $package"
Write-Host "Openssl version: $opensslVersion"

# Pause to keep the window open
Read-Host -Prompt "Press Enter to exit"
