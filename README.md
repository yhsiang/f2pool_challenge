# Interview challenge

## development histories
1. https://github.com/go-swagger/go-swagger
2. swagger generate server -f ./swagger.json
3. implement lookup domain (net.LookupIP)
4. implement validate ip (net.ParseIP)
5. implement queries history (redis list)
6. add root endpoint
7. add dockerfile
8. add docker-compose.yml
9. add some test cases

## TODO
1. detect kubernetes
2. Prometheus
3. health check

## Development

* prerequisite
   - redis

* execute server

```
$ godotenv -f .env -- go run cmd/interview-challenge-server/main.go --port 3000
```