#!/usr/bin/env bash

source ./proto_dir.cfg



for ((i = 0; i < ${#all_proto[*]}; i++)); do
  protoFile=${all_proto[$i]}
  protoName=$(basename "$protoFile" .proto)
  protoc  --doc_out=./docs --doc_opt=html,$protoName.html --go_out=plugins=grpc,module=grpc-demo/demo-2/proto:./   $protoFile
  echo "protoc --go_out=plugins=grpc:." $protoFile
done
echo "proto file generate success..."


#  --validate_out="lang=go:./gen"