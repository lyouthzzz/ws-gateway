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

## benchmark
### 安装压测工具
```bash
go get github.com/lyouthzzz/websocket-benchmark-cli@main
```

### case - 1

- websocket客户端：1w
- 发送间隔：1s
- 数据大小：50b

```bash
 websocket-benchmark-cli message --file testdata/50b.txt  --interval 1s --times 10000 --user 10000 --host 127.0.0.1:8080 --path /gateway/ws
```

服务客户端表现稳定

```bash
docker stats ws-gateway

CONTAINER ID   NAME         CPU %     MEM USAGE / LIMIT   MEM %     NET I/O         BLOCK I/O   PIDS
7144b1c4efff   ws-gateway   6.86%     380.4MiB / 2GiB     18.58%    854MB / 593MB   0B / 0B     6
```

![统计图.png](docs/benchmark-50b.png)

### case - 2

- websocket客户端：1w
- 发送间隔：1s
- 数据大小：1k

```bash
websocket-benchmark-cli message --file testdata/1k.txt  --interval 1s --times 10000 --user 10000 --host 127.0.0.1:8080 --path /gateway/ws
```

出现瓶颈，客户端报错：write: connection timed out 原因：服务端处理消息出现瓶颈（可能是 ws-gateway -> ws-api出现问题，
但是内网gRPC-Streaming流量支持不会那么差），导致客户端写入超时。

解决方案：

- ws-gateway -> ws-api benchmark
- 增加ws-api节点

![统计图.png](docs/benchmark-1k.png)



