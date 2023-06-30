#!/usr/bin/env bash

source ./proto_dir.cfg

doc_dir="./docs"

if [ ! -d "$doc_dir" ]; then
  mkdir "$doc_dir"
  echo "mkdir $doc_dir"
fi

# 生成待校验的proto文件以及html接口文档
for ((i = 0; i < ${#all_proto[*]}; i++)); do
  protoPath=${all_proto[$i]}
  protoName=$(basename "$protoPath" .proto)
  protoc  --doc_out=$doc_dir --doc_opt=html,$protoName.html \
          --validate_out="lang=go,paths=source_relative:./" \
          --go_out=plugins=grpc,module=grpc-demo/demo-2/proto:./   $protoPath
  echo "protoc --go_out=plugins=grpc:." $protoPath
done
echo "proto file generate success"