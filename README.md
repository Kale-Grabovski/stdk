# Stfk

Initializing: `cp .bin/config.dist.json .bin/config.json && make`

Run: `docker-compose up`

Open CRM:

`http://localhost:8080`

API - Difference of versions:

```
curl http://localhost:8081/api/v1/modules -X POST -d '{"installedModules": [{"id": "proxy", "version": 3}, {"id": "ads", "version": 1}], "deviceId": "fojew"}' -v --header "Content-Type: application/json"
```

API - Getting binary by module name:

`http://localhost:8081/api/v1/modules/ads`

## Optionally

Run watcher to rebuild go container on any change to *.go files:

`watchexec --restart --exts go --watch . "make build && docker-compose restart crm && docker-compose restart api"`
