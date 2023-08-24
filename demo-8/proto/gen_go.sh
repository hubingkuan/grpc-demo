doc_dir="./docs"

if [ ! -d "$doc_dir" ]; then
  mkdir "$doc_dir"
  echo "mkdir $doc_dir"
fi

swagger_dir="./swagger"
if [ ! -d "$swagger_dir" ]; then
  mkdir "$swagger_dir"
  echo "mkdir $swagger_dir"
fi


protoc  --doc_out=$doc_dir --doc_opt=html,proto_doc.html \
          --validate_out="lang=go,paths=source_relative:./" \
          --grpc-gateway_out ./ \
          --grpc-gateway_opt paths=source_relative \
          --grpc-gateway_opt logtostderr=true \
          --grpc-gateway_opt generate_unbound_methods=true \
          --openapiv2_out=$swagger_dir  --openapiv2_opt logtostderr=true --openapiv2_opt generate_unbound_methods=true  --openapiv2_opt allow_merge=true,merge_file_name=demo-8 \
          --go_out=plugins=grpc:./ --go_opt=paths=source_relative   */*.proto


protoc-go-inject-tag -input=*/*.pb.go