## 실행방법

```
$ go build
$ go run main.go
```

## API Endpoint
### 토픽생성
```
curl --location 'http://localhost:8080/topics' \
--header 'Content-Type: application/json' \
--data '{
    "name": "토픽"
}'
```

### 토픽 목록 조회
```
curl --location 'http://localhost:8080/topics/'
```

### 토픽 조회
```
curl --location 'http://localhost:8080/topics/1/'
```

### 토픽 구독
```
curl --location 'http://localhost:8080/topics/1/subscribe' \
--header 'Content-Type: application/json' \
--data '{
    "UserID": "9"
}'
```

### 토픽 발행
```
curl --location 'http://localhost:8080/topics/1/publish' \
--header 'Content-Type: application/json' \
--data '{
    "message": "hi"
}'
```