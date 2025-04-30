# Memo
### How to Run
Start Redis in your local environment with the command below, then run `go run ./ex53`:
```
$ docker run --name redis-local -d -p 6379:6379 redis
```
The todo items are stored in a Redis hash with the key `ex53:tasks`, viewable using redis-cli.
```
127.0.0.1:6379> hgetall ex53:tasks 
1) "5"
2) "my todo 5"
```