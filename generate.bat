@echo off
echo Installing required tools...
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

echo Getting googleapis...
if not exist "third_party\google\api" (
    mkdir third_party\google\api
    curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o third_party/google/api/annotations.proto
    curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o third_party/google/api/http.proto
)

echo Generating Protobuf code...
for %%f in (api\proto\*.proto) do (
    echo Generating %%f
    protoc -I ./api/proto -I ./third_party ^
      --go_out . --go_opt paths=source_relative ^
      --go-grpc_out . --go-grpc_opt paths=source_relative ^
      --grpc-gateway_out . --grpc-gateway_opt paths=source_relative ^
      %%f
)

echo Installing sqlc...
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo Generating DB code with sqlc...
sqlc generate

echo Generating Swagger/OpenAPI docs...
protoc -I ./api/proto -I ./third_party --openapiv2_out ./api/proto --openapiv2_opt logtostderr=true api/proto/*.proto

echo Done!
