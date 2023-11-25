package redis_consumer

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"service-mutasi/dto"
)

// Consume consumes stream data sent to redis stream
func (consumer *RedisConsumer) consume(ctx context.Context) error {
	streams, err := consumer.store.redis.GetStreams(ctx, consumer.config.RedisStreamRequest, 1)
	if err != nil {
		fmt.Printf("failed to get stream (%s): %s\n", consumer.config.RedisStreamRequest, err)

		return err
	}

	for _, stream := range streams {
		for _, message := range stream.Messages {
			ctx := context.Background()

			values := message.Values

			err = consumer.store.redis.DeleteFromStream(ctx, consumer.config.RedisStreamRequest, message.ID)
			if err != nil {
				fmt.Printf("failed to delete message stream (%s): %s\n", message.ID, err)

				return err
			}

			fmt.Printf("values: %+v\n", values)

			reqType, ok := values["req_type"].(string)
			if !ok {
				fmt.Printf("error assert req_type got type %T", values["req_type"])
			}

			switch reqType {
			case "tabung":
				{
					request := &dto.TabungRequest{}

					nominalString, ok := values["nominal"].(string)
					if !ok {
						fmt.Printf("error assert nominal, got type %T", values["nominal"])

						return errors.New("error assert nominal")
					}
					nominal, err := strconv.ParseInt(nominalString, 10, 64)
					if err != nil {
						fmt.Println("error converting string to int64:", err)

						return errors.New("error assert nominal")
					}
					request.Nominal = nominal

					noRekening, ok := values["no_rekening"].(string)
					if !ok {
						fmt.Printf("error assert no_rekening got type %T", values["no_rekening"])

						return errors.New("error assert no_rekening")
					}
					request.NoRekening = noRekening

					_, err = consumer.service.Tabung(ctx, *request)
					if err != nil {
						fmt.Printf("error: %s", err)
					}
				}
			case "tarik":
				{
					request := &dto.TarikRequest{}

					nominalString, ok := values["nominal"].(string)
					if !ok {
						fmt.Printf("error assert nominal, got type %T", values["nominal"])

						return errors.New("error assert nominal")
					}
					nominal, err := strconv.ParseInt(nominalString, 10, 64)
					if err != nil {
						fmt.Println("error converting string to int64:", err)

						return errors.New("error assert nominal")
					}
					request.Nominal = nominal

					noRekening, ok := values["no_rekening"].(string)
					if !ok {
						fmt.Printf("error assert no_rekening got type %T", values["no_rekening"])

						return errors.New("error assert no_rekening")
					}
					request.NoRekening = noRekening

					_, err = consumer.service.Tarik(ctx, *request)
					if err != nil {
						fmt.Printf("error: %s", err)
					}
				}
			default:
				{
					fmt.Printf("unknown req type: %s", reqType)
				}
			}
		}
	}

	return nil
}