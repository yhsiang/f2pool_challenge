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
10. detect application running in kubernetes
11. support Prometheus
12. support health check
13. add test for verifying maximum size of queries history

## TODO
1. move `/metrics` and `/health` into configure_interview_challenge

## Development

* prerequisite
   - redis

* execute server

```
$ godotenv -f .env -- go run cmd/interview-challenge-server/main.go --port 3000
```

## Deployment

use github actions to deploy to linode k8s

1. use github packages as container registry
2. create personal access token as password
3. setup secret of github action DOCKERHUB_TOKEN
4. setup a linode k8s
5. setup k8s secret
   `kubectl create secret docker-registry ghcr --docker-server=https://ghcr.io/v2/ --docker-username=yhsiang --docker-password=<GITHUB_PERSONAL_TOKEN> --docker-email=<EMAIL>`
6. create deployment
   `kubectl apply -f deployment.yaml`
7. download kube-config
8. cat kube-config file and encode to base64 string
9. setup secret of github action KUBE_CONFIG 
10. use port-forwarding to test 
   `kubectl port-forward <pod-name> 3000:3000`

**NOTE**

I use a sidecar pattern to deploy redis, and known issue is about persistence. It's better to change it into redis pod with persistent volume or redis instance (recommend). Redis instance can be Amazon ElastiCache for Redis, GCP Memorystore, RedisLab or self-hosted instance.