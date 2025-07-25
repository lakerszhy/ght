# Ght Installer for Windows

# This script installs ght on Windows using PowerShell.
#
# Usage:
#   irm https://raw.githubusercontent.com/lakerszhy/ght/main/scripts/install.ps1 | iex
#
# The script will:
# 1. Detect the user's architecture (amd64 or arm64).
# 2. Download the latest release of ght for Windows.
# 3. Create a directory at '$HOME\AppData\Local\ght'.
# 4. Extract the binary to the new directory.
# 5. Add the directory to the user's PATH environment variable.

param()

$ErrorActionPreference = 'Stop'

function Main {
    # --- Detect Architecture ---
    $arch = switch ($env:PROCESSOR_ARCHITECTURE) {
        'AMD64' { 'amd64' }
        'ARM64' { 'arm64' }
        default { Write-Host "Unsupported architecture: $env:PROCESSOR_ARCHITECTURE" -ForegroundColor Red; exit 1 }
    }

    # --- Get Latest Version ---
    try {
        $response = Invoke-WebRequest -Uri "https://github.com/lakerszhy/ght/releases/latest"
        $latest_version_url = $response.BaseResponse.ResponseUri.AbsoluteUri
        $latest_version = $latest_version_url.Split('/')[-1]
    } catch {
        Write-Host "Failed to fetch the latest version of ght." -ForegroundColor Red
        exit 1
    }

    if (-not $latest_version) {
        Write-Host "Failed to fetch the latest version of ght." -ForegroundColor Red
        exit 1
    }

    # --- Download and Extract ---
    $download_url = "https://github.com/lakerszhy/ght/releases/download/$latest_version/ght-$latest_version-windows_$($arch).zip"
    $install_dir = "$env:LOCALAPPDATA\ght"
    $temp_zip_path = "$env:TEMP\ght.zip"

    Write-Host "Downloading ght $latest_version for Windows/$arch..." -ForegroundColor Green

    try {
        Invoke-WebRequest -Uri $download_url -OutFile $temp_zip_path
    } catch {
        Write-Host "Failed to download ght. Please check the URL and your network connection." -ForegroundColor Red
        exit 1
    }

    Write-Host "Download complete." -ForegroundColor Green

    # --- Installation ---
    $temp_extract_dir = Join-Path -Path $env:TEMP -ChildPath "ght_extracted"
    if (Test-Path -Path $temp_extract_dir) {
        Remove-Item -Path $temp_extract_dir -Recurse -Force
    }
    New-Item -ItemType Directory -Path $temp_extract_dir | Out-Null

    Write-Host "Extracting files..." -ForegroundColor Yellow
    Expand-Archive -Path $temp_zip_path -DestinationPath $temp_extract_dir -Force

    $archive_dir_name = "ght-$latest_version-windows_$($arch)"
    $source_exe_path = Join-Path -Path $temp_extract_dir -ChildPath "$archive_dir_name\ght.exe"

    if (-not (Test-Path -Path $install_dir)) {
        New-Item -ItemType Directory -Path $install_dir | Out-Null
    }

    Write-Host "Installing ght to $install_dir..." -ForegroundColor Yellow
    Move-Item -Path $source_exe_path -Destination $install_dir -Force

    # --- Add to PATH ---
    $currentUserPath = [System.Environment]::GetEnvironmentVariable('Path', 'User')
    if (-not ($currentUserPath -split ';' -contains $install_dir)) {
        Write-Host "Adding $install_dir to your PATH." -ForegroundColor Yellow
        $newPath = "$currentUserPath;$install_dir"
        [System.Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
        $env:Path = $newPath # Update for current session
    }

    # --- Cleanup ---
    Remove-Item -Path $temp_zip_path
    Remove-Item -Path $temp_extract_dir -Recurse -Force

    Write-Host "Installation complete!" -ForegroundColor Green
    Write-Host "Please restart your terminal for the PATH changes to take full effect."
    Write-Host "You can now run 'ght' to start the application."
}

Main