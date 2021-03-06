# fiber-redis-go
Golang-based pub/sub model using fiber and redis.<br/>

## Install
```
$> docker-compose up --build -d
```

## Run
### Publisher run
```
$> go run publ.go
```

### Subscriber run
```
$> go run subs.go
```

### Chrome postman run
```
1. launch Postman
2. POST, JSON setting

{
    "name" : "MyName",
    "email": "My-Email"
}

```

## Reference
<ul>
<li>https://dev.to/franciscomendes10866/using-redis-pub-sub-with-golang-mf9</li>
<li>fiber: https://gofiber.io/</li>
<li>go-redis: https://redis.uptrace.dev/</li>
</ul>
