package db

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//InitConfig is initialize method for db connection pool
func InitConfig(hostName string, port int) {
	var redisAddress = fmt.Sprintf("%s:%s", hostName, strconv.Itoa(port))
	pool = newPool(redisAddress)
	cleanupHook()
}

func newPool(redisAddress string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		pool.Close()
		os.Exit(0)
	}()
}

//Ping db connection
func Ping() error {
	conn := pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

//Get value from db by key
func Get(key string) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

//GetByPattern value from db by pattern key
func GetByPattern(pattern string) ([][]byte, error) {
	conn := pool.Get()
	defer conn.Close()

	var data [][]byte

	keys, err := redis.Strings(conn.Do("KEYS", pattern))
	if err != nil {
		return data, fmt.Errorf("error getting pattern %s: %v", pattern, err)
	}
	for _, key := range keys {
		dataByKey, err := redis.Bytes(conn.Do("GET", key))
		if err != nil {
			return data, fmt.Errorf("error getting key %s: %v", key, err)
		}

		data = append(data, dataByKey)
	}

	return data, err
}

//Set value from db by key
func Set(key string, value []byte) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

//Exists is check value in db by key
func Exists(key string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

//Delete value from db by key
func Delete(key string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

//GetKeys is method return all keys from db
func GetKeys(pattern string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}
