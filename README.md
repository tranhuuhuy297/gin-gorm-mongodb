# Project stater with Golang and Gin

```
go run main.go
```
## Docker for running Mongo and Redis
- Mongo
```
docker run -d -p 27017:27017 –name=mongo mongo:latest
```
- Redis
```
docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```