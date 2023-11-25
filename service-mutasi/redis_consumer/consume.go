package redis_consumer

import (
	"context"
	"errors"
	"fmt"

	"service-mutasi/dto"
)

// Consume consumes stream data sent to redis stream
func (consumer *RedisConsumer) consume(ctx context.Context) error {
	streams, err := consumer.store.redis.GetStreams(ctx, consumer.streamName, 1)
	if err != nil {
		fmt.Printf("failed to get stream (%s): %s\n", consumer.streamName, err)

		return err
	}

	for _, stream := range streams {
		for _, message := range stream.Messages {
			ctx := context.Background()

			values := message.Values

			err = consumer.store.redis.DeleteFromStream(ctx, consumer.streamName, message.ID)
			if err != nil {
				fmt.Printf("failed to delete message stream (%s): %s\n", message.ID, err)

				return err
			}

			fmt.Printf("values: %+v\n", values)

			switch consumer.streamName {
			case consumer.config.RedisStreamRequestTabung:
				{
					request := &dto.TabungRequest{}

					nominal, ok := values["nominal"].(int64)
					if !ok {
						fmt.Printf("error assert nominal, got type %T", values["nominal"])

						return errors.New("error assert nominal")
					}
					request.Nominal = nominal

					noRekening, ok := values["no_rekening"].(string)
					if !ok {
						fmt.Printf("error assert no_rekening")

						fmt.Printf("error assert no_rekening got type %T", values["no_rekening"])
					}
					request.NoRekening = noRekening

					_, err := consumer.service.Tabung(ctx, *request)
					if err != nil {
						fmt.Printf("error: %s", err)
					}
				}
			case consumer.config.RedisStreamRequestTarik:
				{
					request := &dto.TarikRequest{}

					nominal, ok := values["nominal"].(int64)
					if !ok {
						fmt.Printf("error assert nominal, got type %T", values["nominal"])

						return errors.New("error assert nominal")
					}
					request.Nominal = nominal

					noRekening, ok := values["no_rekening"].(string)
					if !ok {
						fmt.Printf("error assert no_rekening got type %T", values["no_rekening"])

						return errors.New("error assert no_rekening")
					}
					request.NoRekening = noRekening

					_, err := consumer.service.Tarik(ctx, *request)
					if err != nil {
						fmt.Printf("error: %s", err)
					}
				}
			default:
				{
					fmt.Printf("unknown stream name: %s", consumer.streamName)
				}
			}
		}
	}

	return nil
}