# ws-gateway

## build

### ws-gateway

```bash
docker build --build-arg APP_NAME=ws-gateway -f deploy/docker/Dockerfile -t ws-gateway .
```

### ws-api

```bash
docker build --build-arg APP_NAME=ws-api -f deploy/docker/Dockerfile -t ws-api .
```