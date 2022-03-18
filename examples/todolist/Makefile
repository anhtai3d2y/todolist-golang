gen-todolist: protoc --go_out=. --go_opt=paths=source_relative .\todolist\todolist.proto
gen-todolist-grpc: protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\todolist\todolist.proto

protoc --go_out=. --go_opt=paths=source_relative .\todolist\todolist.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\todolist\todolist.proto
protoc -I . --grpc-gateway_out ./gen/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    /todolist/todolist.proto


go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

mkdir -p google/api
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto