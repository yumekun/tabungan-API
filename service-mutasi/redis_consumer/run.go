package redis_consumer

import (
	"context"
	"fmt"
)

func (consumer *RedisConsumer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			{
				// create new context everytime core logic executed
				ctx := context.Background()

				// execute core logic
				err := consumer.consume(ctx)
				if err != nil {
					fmt.Printf("failed to consume: %s", err)
				}
			}
		}
	}
}