package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
)

// AddOrder adds an order to the Redis queue.
func (q RedisQueue) AddOrder(ctx context.Context, order *entities.Order) error {
	conn := q.pool.Get()
	defer conn.Close()

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal order: %v", err)
	}

	//connWithContext := conn(ctx)
	//if err := connWithContext.Err(); err != nil {
	//	return fmt.Errorf("failed to create context with Redis connection: %v", err)
	//}

	if err := conn.Send("PUBLISH", "orders", orderJSON); err != nil {
		return fmt.Errorf("failed to send order to Redis queue: %v", err)
	}

	if err := conn.Flush(); err != nil {
		return fmt.Errorf("failed to flush Redis connection: %v", err)
	}
	return nil
}
