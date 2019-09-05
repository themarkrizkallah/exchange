package kafka

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"

	"matcher/engine"
	"matcher/env"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	Ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	orderbook := engine.GetOrderBook()

	for message := range claim.Messages() {
		order, err := engine.ProtoToOrder(message.Value)
		if err != nil {
			log.Printf("Error decoding message: %v\n", message)
			continue
		}

		log.Printf("Processing: %+v\n", order)
		trades := orderbook.Process(order) // Process the order

		if len(trades) > 0 {
			fmt.Printf("Completed Trade(s): %+v\n", trades)
		}

		for _, trade := range trades {
			producer.Input() <- &sarama.ProducerMessage{
				Topic: "trades",
				Value: sarama.ByteEncoder(trade.ToProto()),
			}
		}

		// Mark the message as processed
		session.MarkMessage(message, "")
	}

	return nil
}

func SetupConsumer(brokers []string) (Consumer, sarama.ConsumerGroup, error){
	log.Println("Starting a new Sarama consumer...")

	version, err := sarama.ParseKafkaVersion(env.KafkaVersion)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/*
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	//Setup a new Sarama consumer group
	consumer := Consumer{
		Ready: make(chan bool),
	}

	consumerClient, err := sarama.NewConsumerGroup(brokers, env.KafkaConsGroup, config)

	return consumer, consumerClient, err
}
