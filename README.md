# GO Contacts

## Configuration
Config file resides in ```config/config.yaml```

Values:
``` yaml
Host: localhost
Port: 8080
Dsn: gocontacts:gocontacts@tcp(localhost:3306)/gocontacts?charset=utf8&parseTime=true
```

## Fetching project dependencies
``` shell
go mod vendor
```


### TODO
- [x] Endpoints
- [ ] flags: timeouts, ports
- [x] logging
- [x] error handling
- [x] metrics -> prometheus
- [ ] test coverage, 70%
- [ ] e2e test
- [ ] Dockerfile config
- [x] Database
