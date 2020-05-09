package myredis

import (
	"fmt"
	"sync"
	"time"

	"github.com/PuerkitoBio/redisc"
	"github.com/gomodule/redigo/redis"
)

var once sync.Once
var cluster redisc.Cluster
var pool *redis.Pool

const (
	ttl      = 3
	redisUrl = "localhost:6379"
)

func GetRedisCluster() redisc.Cluster {
	once.Do(func() {
		fmt.Println("Initializing redis conn cluster!")

		cluster := redisc.Cluster{
			StartupNodes: []string{redisUrl},
			DialOptions:  []redis.DialOption{redis.DialConnectTimeout(2 * time.Second)},
		}

		DoPingTestCluster(cluster)

	})
	return cluster
}

func GetRedisPool() *redis.Pool {
	once.Do(func() {
		fmt.Println("Initializing redis conn pool!")
		pool = newPool()
		DoPingTestPool(pool)
	})
	return pool
}

func DoPingTestPool(pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()
	err := ping(conn)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func DoPingTestCluster(cluster redisc.Cluster) error {
	conn := cluster.Get()
	defer conn.Close()
	err := ping(conn)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisUrl
		ppppp
	
	
	pl)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func MyHset(hash, key, value string) {
	//c := GetRedisPool().Get()
	cl := GetRedisCluster()
	c := cl.Get()
	defer c.Close()
	_, err := c.Do("HSET", hash, key, value)
	if err != nil {
		fmt.Println("Error in HSET , hash - ", hash, " key - ", key, " Error - ", err)
	}
}

func MySetExp(hash string) {
	//c := GetRedisPool().Get()
	cl := GetRedisCluster()
	c := cl.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", hash, ttl)
	if err != nil {
		fmt.Println("Error in EXPIRE , hash - ", hash, " Error - ", err)
	}
}

func MyHget(hash, key string) string {
	//c := GetRedisPool().Get()
	cl := GetRedisCluster()
	c := cl.Get()
	defer c.Close()
	value, err := redis.String(c.Do("HGET", hash, key))
	if err != nil {
		fmt.Println("Error - ", err)
		return ""
	}
	return value
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}

	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG
	return nil
}
