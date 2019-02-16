# stdk API

CRM is on 80 port, API on 8081, host http://178.128.37.226

Difference of versions:

```
curl http://178.128.37.226:8081/api/v1/modules -X POST -d '{"installedModules": [{"id": "proxy", "version": 3}, {"id": "ads", "version": 1}], "deviceId": "fojew"}' -v --header "Content-Type: application/json"
```

Getting binary by module name:

`http://178.128.37.226:8081/api/v1/modules/ads`

