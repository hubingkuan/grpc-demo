source ./proto_dir.cfg

doc_dir="./docs"

suffix=".pb.go"

if [ ! -d "$doc_dir" ]; then
  mkdir "$doc_dir"
  echo "mkdir $doc_dir"
fi

for ((i = 0; i < ${#all_proto[*]}; i++)); do
  protoPath=${all_proto[$i]}
  protoName=$(basename "$protoPath" .proto)
  protoc  --doc_out=$doc_dir --doc_opt=html,$protoName.html \
          --validate_out="lang=go,paths=source_relative:./" \
          --grpc-gateway_out ./ \
          --grpc-gateway_opt paths=source_relative \
          --grpc-gateway_opt logtostderr=true \
          --openapiv2_out=./ --openapiv2_opt logtostderr=true \
          --go_out=plugins=grpc:./ --go_opt=paths=source_relative   $protoPath
  echo "protoc generate" $protoPath

  v=$protoName/$protoName$suffix
  protoc-go-inject-tag -input=$v
done
echo "proto file generate success"