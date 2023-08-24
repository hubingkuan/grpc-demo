@echo off
setlocal enabledelayedexpansion

set "doc_dir=./docs"

IF NOT EXIST "%doc_dir%" (
  mkdir "%doc_dir%"
  @echo "mkdir %doc_dir%"
)

set "config_file=proto_dir.cfg"

set suffix=.pb.go

for /f "tokens=*" %%a in ('type "%config_file%" ^| findstr /V /C:")" ^| findstr /V /C:"all_proto=("') do (
   set "all_proto=!all_proto!%%a "
)


for %%i in (%all_proto%) do (
    set "protoPath=%%i"
    for %%j in ("!protoPath!") do set "protoName=%%~nj"
    protoc --doc_out=!doc_dir! --doc_opt=html,!protoName!.html ^
           --validate_out=lang=go,paths=source_relative:./ ^
           --grpc-gateway_out ./ ^
           --grpc-gateway_opt paths=source_relative ^
           --grpc-gateway_opt logtostderr=true ^
           --openapiv2_out=./ --openapiv2_opt logtostderr=true ^
           --go_out=plugins=grpc:./ --go_opt=paths=source_relative !protoPath!
    @echo protoc generate "%%i"

    set v=!protoName!/!protoName!%suffix%
    protoc-go-inject-tag -input=!v!
)

@echo proto file generate success
pause