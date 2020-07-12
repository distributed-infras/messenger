package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	res, err := conn.Do("SET", "key", "50")
	if err != nil {
		panic(err)
	}
	res, err = redis.String(conn.Do("GET", "key"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
