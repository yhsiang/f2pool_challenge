# Interview challenge

## development histories
1. https://github.com/go-swagger/go-swagger
2. swagger generate server -f ./swagger.json
3. implement lookup domain (net.LookupIP)
4. implement validate ip (net.ParseIP)
5. implement queries history (redis list)
   - I don't really test for twenty items
6. add root endpoint
7. add dockerfile
8. add docker-compose.yml
9. add some test cases
10. detect kubernetes
11. Prometheus
12. health check

## TODO
1. move `/metrics` and `/health` into configure_interview_challenge

## Development

* prerequisite
   - redis

* execute server

```
$ godotenv -f .env -- go run cmd/interview-challenge-server/main.go --port 3000
```