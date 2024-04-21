
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/**/*.proto



#protoc --go_out=paths=source_relative:proto/ proto/*.proto / 안됨 
#protoc --go_out=paths=import,plugins=grpc:proto/ proto/*.proto

protoc --go_out=. --go-grpc_out=. proto/*.proto

#protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/*.proto