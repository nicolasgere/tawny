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

### Generate Typescript web client
```
protoc -I ./protos  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:./dist/typescript-web/src" --ts_out="service=grpc-web:./dist/typescript-web/src"  protos/Push.proto  protos/Presence.proto 
```


### Generate Typescript node client
```
 protoc -I ./protos  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:./dist/typescript-node/src" --ts_out="./dist/typescript-node/src"  protos/Push.proto  protos/Presence.proto protos/Admin.proto 
```
