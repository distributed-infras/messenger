# Redis Prototype
This example shows how to interact with Redis with Golang

## Getting Started

### Start Redis server

```bash
redis-server
```

You should see the following output in the terminal:

```
35065:C 11 Jul 2020 23:27:12.482 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
35065:C 11 Jul 2020 23:27:12.482 # Redis version=6.0.5, bits=64, commit=00000000, modified=0, pid=35065, just started
35065:C 11 Jul 2020 23:27:12.482 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
35065:M 11 Jul 2020 23:27:12.484 * Increased maximum number of open files to 10032 (it was originally set to 256).
                _._                                                  
           _.-``__ ''-._                                             
      _.-``    `.  `_.  ''-._           Redis 6.0.5 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._                                   
 (    '      ,       .-`  | `,    )     Running in standalone mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
 |    `-._   `._    /     _.-'    |     PID: 35065
  `-._    `-._  `-./  _.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |           http://redis.io        
  `-._    `-._`-.__.-'_.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |                                  
  `-._    `-._`-.__.-'_.-'    _.-'                                   
      `-._    `-.__.-'    _.-'                                       
          `-._        _.-'                                           
              `-.__.-'                                               

35065:M 11 Jul 2020 23:27:12.485 # Server initialized
35065:M 11 Jul 2020 23:27:12.485 * Loading RDB produced by version 6.0.5
35065:M 11 Jul 2020 23:27:12.485 * RDB age 2 seconds
35065:M 11 Jul 2020 23:27:12.485 * RDB memory usage when created 0.99 Mb
35065:M 11 Jul 2020 23:27:12.485 * DB loaded from disk: 0.000 seconds
35065:M 11 Jul 2020 23:27:12.485 * Ready to accept connections
```

### Run example

```bash
go run main.go
```

You should see the follow output:

```
50
```
