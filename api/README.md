# API proto
## Make Api Service Directory
mk dir auth/service/v1
mk dir user/service/v1
mk dir product/service/v1

## Make .proto files 
auth/service/v1/auth.proto
user/service/v1/user.proto
product/service/v1/product.proto

## Generate protoc
protoc --proto_path=. `
       --proto_path=C:\Users\HF\go\pkg\mod\github.com\go-kratos\kratos@v1.0.1\third_party `
       --go_out=. --go_opt=paths=source_relative `
       --go-http_out=. --go-http_opt=paths=source_relative `
       product.proto

protoc --proto_path=. `
       --proto_path="C:\Users\HF\go\pkg\mod\github.com\go-kratos\kratos@v1.0.1\third_party" `
       --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       --go-http_out=. --go-http_opt=paths=source_relative `
       kratos/api/user/service/v1/user.proto

protoc --proto_path=. `
       --proto_path="C:\Users\HF\go\pkg\mod\github.com\go-kratos\kratos@v1.0.1\third_party" `
       --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       --go-http_out=. --go-http_opt=paths=source_relative `
       kratos/api/product/service/v1/product.proto

protoc --proto_path=. `
       --proto_path="C:\Users\HF\go\pkg\mod\github.com\go-kratos\kratos@v1.0.1\third_party" `
       --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       --go-http_out=. --go-http_opt=paths=source_relative `
       kratos/api/auth/service/v1/auth.proto
