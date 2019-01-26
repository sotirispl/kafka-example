package main

import (
	"flag"
	"log"

	"github.com/Shopify/sarama"
)

var (
	brokers = []string{"localhost:9092"}
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topicArg := flag.String("topic", "the_topic", "the topic that the message will be produced into")
	msgArg := flag.String("msg", "an important message", "the message that will be produced")
	keyArg := flag.String("key", "a_key", "the key of the message")
	flag.Parse()

	msg := &sarama.ProducerMessage{
		Topic: *topicArg,
		Value: sarama.StringEncoder(*msgArg),
		Key:   sarama.StringEncoder(*keyArg),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	log.Printf("Send a message in topic %s - partition %d  - offset %d\n", *topicArg, partition, offset)
}
