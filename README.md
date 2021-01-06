# GO Contacts

## Fetching project dependencies
``` shell
go mod vendor
```

## Build project
``` shell
go build
```

## Build and start docker image
``` shell
docker build -t go-contacts .
docker run -t go-contacts
```

## TODO
- [x] endpoints
- [x] flags: timeouts, ports
- [x] error handling with messages
- [x] Persistence  
- [x] metrics -> prometheus
- [x] logging fulfilling request time
- [x] Dockerfile config
- [ ] test coverage, 70%
- [ ] e2e test

