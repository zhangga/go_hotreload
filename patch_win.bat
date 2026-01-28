@echo off
setlocal enabledelayedexpansion

set GOCMD=go

pushd example
if errorlevel 1 (
    echo [ERROR] Failed to enter directory: example
    exit /b 1
)

echo Building patch_v1.dll ...

%GOCMD% build -buildmode=c-shared -o=patch_v1.dll ./patch_v1
if errorlevel 1 (
    echo [ERROR] Go build failed!
    popd
    exit /b 1
)

popd
echo [OK] Patch build success.
pause