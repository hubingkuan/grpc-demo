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
           --grpc-gateway_opt standalone=true ^
           --go_out=plugins=grpc,module=grpc-demo/demo-2/proto:./ !protoPath!
    @echo protoc --go_out=plugins=grpc:. "%%i"

    set v=!protoName!/!protoName!%suffix%
    protoc-go-inject-tag -input=!v!
)

@echo proto file generate success
pause