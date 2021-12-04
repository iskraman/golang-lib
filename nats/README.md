# NATS Publish and Subscribe
Nats Server, Client를 활용하여 Pub / Sub 기능을 구현한다.

## Prepare
N/A

## Install
Docker 기반의 Nats Server를 구동한다.
```
$> docker run --rm -d -p 4222:4222 -p 8222:8222 --name nats nats:latest
$> docker logs nats
[1] 2021/12/04 14:35:32.000916 [INF] Starting nats-server
[1] 2021/12/04 14:35:32.001047 [INF]   Version:  2.6.6
[1] 2021/12/04 14:35:32.001050 [INF]   Git:      [878afad]
[1] 2021/12/04 14:35:32.001051 [INF]   Name:     NAEN65M23LTRJCOODRCKPCIIUWI6WXXL5E26OE3HA5LMFLBSAS7SM63Y
[1] 2021/12/04 14:35:32.001052 [INF]   ID:       NAEN65M23LTRJCOODRCKPCIIUWI6WXXL5E26OE3HA5LMFLBSAS7SM63Y
[1] 2021/12/04 14:35:32.001054 [INF] Using configuration file: nats-server.conf
[1] 2021/12/04 14:35:32.001880 [INF] Starting http monitor on 0.0.0.0:8222
[1] 2021/12/04 14:35:32.002027 [INF] Listening for client connections on 0.0.0.0:4222
[1] 2021/12/04 14:35:32.002287 [INF] Server is ready
[1] 2021/12/04 14:35:32.002318 [INF] Cluster name is bLxQHM7sQ5VZbT85Ki0zeO
[1] 2021/12/04 14:35:32.002321 [WRN] Cluster name was dynamically generated, consider setting one
[1] 2021/12/04 14:35:32.002388 [INF] Listening for route connections on 0.0.0.0:6222
```
Pub/Sub 연결을 테스트할 수 있는 natscli를 다운로드한다.
```
$> git clone https://github.com/nats-io/natscli
$> cd natscli
$> go build
```

## Run
CLI Pub/Sub Test
```
$> nats sub test.event
// Subscribe 역할을 담당할 CLI

$> nats pub test.event hello
// Publisher 역할을 담당할 CLI
```

Golang Pub/Sub Test
```
// nats/sub : Subscriber
$> go run main.go

// nats/pub : Publisher
$> go run main.go
```

## Reference
<ul>
<li>https://nats.io/</li>
<li>https://www.youtube.com/watch?v=Q7ZqQt87tz0</li>
<li>https://www.youtube.com/watch?v=h1K800PuBHM&t=286s</li>
</ul>
