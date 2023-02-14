# Auth Service Documentation 

### App build and run: 
```
docker-compose up
```

#### Delete app: 
```
docker-compose down 
```

### Generate proto: 
```
protoc -I ./proto \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
  ./proto/auth.proto
```
##### If proto packages not found enter in console: 
```
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
```

### Migrations 
For applying migrations use goose library.

Applying all migrations:
```
goose up 
```


Roll back a single migration from the current version:
```
goose down
```
