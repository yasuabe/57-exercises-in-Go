# MEMO
## How to check for Redis updates
```
$ docker exec -it redis-local redis-cli

127.0.0.1:6379> KEYS ex54:short_urls:*
1) "ex54:short_urls:aUKYON"
2) "ex54:short_urls:sequence"
3) "ex54:short_urls:aUKYOM"
127.0.0.1:6379> HGETALL ex54:short_urls:aUKYON
1) "long_url"
2) "https://en.wikipedia.org/wiki/Go_(programming_language)"
3) "visit_count"
4) "0"
127.0.0.1:6379> HGETALL ex54:short_urls:aUKYOM
1) "long_url"
2) "https://en.wikipedia.org/wiki/Redis"
3) "visit_count"
4) "1"
127.0.0.1:6379> 
```
