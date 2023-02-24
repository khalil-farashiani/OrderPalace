package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (q RedisQueue) ReceiveOrder(data chan []byte) {
	// Subscribe to the Redis queue
	psc := redis.PubSubConn{Conn: q.conn}
	psc.Subscribe("orders")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Println(string(v.Data))
			data <- v.Data
			// send acknowledgement to Redis server
			psc.Conn.Do("PACK", "orders", v.Pattern, v.Channel)
		case redis.Subscription:
			log.Printf("Subscribed to %s\n", v.Channel)
		case redis.Pong:
			// ignore pong message
		case error:
			log.Println("Error receiving message:", v)
			psc.Conn.Do("NACK", "orders", v.(redis.Error).Error())
		}
		psc.Conn.Flush()
	}
}
