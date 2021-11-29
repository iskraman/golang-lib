# Cache ALL THE THINGS using Redis with Golang / Go
Golang 기반으로 Redis 연결, HSet, HGet등을 처리하는 기본 모듈

## Prepare
<ul>
  <li>Redis Server</li>
</ul>

URL, Password 정보를 확인한다.
```
        opt := redis.DialPassword("changeme")
        conn, err := redis.Dial("tcp", "localhost:6379", opt)
```

## Install
N/A

## Run
```
$> go run main.go
```

## Reference
https://www.youtube.com/watch?v=jqI4g8_11eY
