# Gorilla / websocket 패키지를 사용한 Go WebSocket
gorilla 패키지를 활용한 웹소켓 통신 방법

## Prepare
N/A

## Install
```
go get github.com/gorilla/websocket
```

## Run
Web Socket server를 기동한다.
```
$> go run main.go
```

Web Socket client를 기동한다.
```
$ python -m SimpleHTTPServer
Serving HTTP on 0.0.0.0 port 8000 ...
```

## Test
크롬 브라우저 -> F12 -> Console
```
http://localhost:8000
```

## Reference
https://www.youtube.com/watch?v=dniVs0xKYKk
