package main

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/tidwall/gjson"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := []string{"localhost:9092"}
	topic := "localhost.db_test_cdc.tbl_test"

	consumer, err := sarama.NewConsumerGroup(brokers, "cdc-group", config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}

	handler := ConsumerGroupHandler{}

	ctx := context.TODO()
	go func() {
		for {
			if err := consumer.Consume(ctx, []string{topic}, &handler); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
		}
	}()

	select {}
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		data := string(message.Value)
		id := gjson.Get(data, "payload.after.id").Int()
		name := gjson.Get(data, "payload.after.name").String()
		email := gjson.Get(data, "payload.after.email").String()
		createdAt := gjson.Get(data, "payload.after.created_at").String()
		fmt.Printf("ID: %d, Name: %s, Email: %s, CreatedAt: %s\n", id, name, email, createdAt)
		session.MarkMessage(message, "")
	}
	return nil
}
