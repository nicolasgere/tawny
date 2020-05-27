### Requirements

- Go 1.14+


### Generate core Golang
```  
protoc -I=./protos --go_out=plugins=grpc:tawny protos/Presence.proto protos/Push.proto 
```
### Generate documentation
```
docker run --rm -v $(pwd)/docs:/out -v $(pwd)/protos:/protos pseudomuto/protoc-gen-doc --doc_opt=markdown,README.md
```

### Generate Typescript client
```
protoc -I ./protos  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:./clients/typescript/src" --ts_out="service=grpc-web:./clients/typescript/src"  protos/Push.proto  protos/Presence.proto 
```
