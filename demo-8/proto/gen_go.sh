protoc  --doc_out=./docs --doc_opt=html,friend.html \
          --validate_out="lang=go,paths=source_relative:./" \
          --grpc-gateway_out ./ \
          --grpc-gateway_opt paths=source_relative \
          --grpc-gateway_opt logtostderr=true \
          --grpc-gateway_opt generate_unbound_methods=true \
          --go_out=plugins=grpc,module=grpc-demo/demo-8/proto:./   friend/friend.proto

protoc-go-inject-tag -input=friend/friend.pb.go