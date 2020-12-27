# GO Contacts

## Configuration
Config file resides in ```config/config.yaml```

Values:
``` yaml
Host: localhost:8080
Dsn: gocontacts:gocontacts@tcp(localhost:3306)/gocontacts?charset=utf8&parseTime=true
```

## Fetching project dependencies
``` shell
go mod vendor
```


### TODO
- [x] endpoints
- [x] flags: timeouts, ports
- [x] error handling with messages
- [x] Persistence  
- [x] metrics -> prometheus
- [ ] logging fulfilling request time
- [ ] test coverage, 70%
- [ ] e2e test
- [ ] Dockerfile config

