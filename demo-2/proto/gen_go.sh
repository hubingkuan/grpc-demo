#!/usr/bin/env bash

source ./proto_dir.cfg

# 生成待校验的proto文件以及html接口文档
for ((i = 0; i < ${#all_proto[*]}; i++)); do
  protoPath=${all_proto[$i]}
  protoName=$(basename "$protoPath" .proto)
  protoc  --doc_out=./docs --doc_opt=html,$protoName.html --validate_out="lang=go,paths=source_relative:./" --go_out=plugins=grpc,module=grpc-demo/demo-2/proto:./   $protoPath
  echo "protoc --go_out=plugins=grpc:." $protoPath
done
echo "proto file generate success..."