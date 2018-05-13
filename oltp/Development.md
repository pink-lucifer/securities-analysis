# Development Notes


## gRpc
1- install binary protoc
> [Protobuf release page](https://github.com/google/protobuf/releases)
```bash
cd ~
mkdir protoc && cd protoc
wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
unzip protoc-3.5.1-linux-x86_64.zip
cd  ~/protoc/include
sudo cp -rf google /usr/local/include/
cd  ~/protoc/bin
sudo cp protoc /usr/local/bin/
protoc --version
```
2- Install protoc plugins: protoc-gen-go, protoc-gen-grpc-gateway, protoc-gen-swagger
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
