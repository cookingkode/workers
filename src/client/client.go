package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type job struct {
	Class string        `json:"class"`
	Args  []interface{} `json:"args"`
}

func AddJob(conn redis.Conn, queue, jobClass string, args ...interface{}) (int64, error) {
	// NOTE: Dirty hack to make a [{}] JSON struct
	if len(args) == 0 {
		args = append(make([]interface{}, 0), make(map[string]interface{}, 0))
	}

	jobJSON, err := json.Marshal(&job{jobClass, args})
	if err != nil {
		return -1, err
	}

	resp, err := conn.Do("RPUSH", queue, string(jobJSON))

	return redis.Int64(resp, err)

}

var (
	redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")
	queueName    = flag.String("queue-name", "goworkerqueue:sampleadd", "Queue name in the form <namespace>queue:<queuename>")
)

func main() {
	conn, err := redis.Dial("tcp", *redisAddress)
	if err != nil {
		fmt.Printf("Unable to connect to Redis %v\n", redisAddress)
		return
	}
	defer conn.Close()

	res, err := AddJob(conn, *queueName, "SampleAddJobClass", 1, 2, 3)

	fmt.Printf("Added Job %v  err=%v", res, err)

}
