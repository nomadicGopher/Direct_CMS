@echo off
set LOG_FILE=build.log
type nul > %LOG_FILE%

echo ---------------------------------------------------------------------------------------------------- >> %LOG_FILE%
echo Build process started at: %date% %time% >> %LOG_FILE%

echo Removing web\go.wasm if it exists... >> %LOG_FILE%
if exist web\go.wasm del web\go.wasm >> %LOG_FILE%

echo Building a new web\go.wasm... >> %LOG_FILE%
set GOOS=js
set GOARCH=wasm
go build -o=web\go.wasm -buildvcs=false main.go >> %LOG_FILE% 2>&1
if %errorlevel% neq 0 (
    echo Failed to build web\go.wasm >> %LOG_FILE%
    exit /b 1
)
echo web\go.wasm was built. >> %LOG_FILE%

echo Removing web\wasm_exec.js if it exists... >> %LOG_FILE%
if exist web\wasm_exec.js del web\wasm_exec.js >> %LOG_FILE%

echo Fetching a new web\wasm_exec.js... >> %LOG_FILE%
copy "%GOROOT%\misc\wasm\wasm_exec.js" web\ >> %LOG_FILE% 2>&1
echo web\wasm_exec.js was fetched from %GOROOT%\misc\wasm\ >> %LOG_FILE%

echo Build process ended at: %date% %time% >> %LOG_FILE%
exit /b