
### Quickstart
```
docker-compose up -d

# For IDE support
docker cp __go_container_id__:/usr/local/go ./pkg/

# For GoLand create an external tool and map it to F10/your favourite button
# /usr/local/bin/docker-compose exec -T golang go run $FilePathRelativeToProjectRoot$ 
# https://www.jetbrains.com/help/idea/configuring-third-party-tools.html
```

### Format
`source aliases && go fmt src/hello.go`
TODO: autoformat on git commit

### Build
`source aliases && go build src/hello.go`



### TODO
- Use package manager to install:
```
docker-compose exec golang go get -d github.com/GoesToEleven/go-programming
docker-compose exec golang go get -d github.com/GoesToEleven/golang-web-dev
docker-compose exec golang go get -d github.com/GoesToEleven/GolangTraining
```
