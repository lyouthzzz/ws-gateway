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

## run
```bash
cd deploy/docker-compose
docker-compose up

Starting ws-api ... done
Starting ws-gateway ... done
Attaching to ws-api, ws-gateway
ws-api        | 2022/11/23 15:27:42 maxprocs: Updating GOMAXPROCS=1: determined from CPU quota
ws-api        | 2022/11/23 15:27:42 gRPC server serve :8081
ws-api        | 2022/11/23 15:27:42 gateway [172.20.0.3] gRPC streaming connect
ws-gateway    | 2022/11/23 15:27:42 maxprocs: Updating GOMAXPROCS=1: determined from CPU quota
ws-gateway    | 2022/11/23 15:27:42 HTTP server serve :8080

```


