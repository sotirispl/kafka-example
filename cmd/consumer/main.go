package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var (
	brokers = []string{"localhost:9092"}
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topicArg := flag.String("topic", "the_topic", "the topic that will consume messages from")
	flag.Parse()

	partitionCons, err := consumer.ConsumePartition(*topicArg, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})
	counter := 0

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		for {
			select {
			case err := <-partitionCons.Errors():
				log.Printf("Got error %v\n", err)
			case msg := <-partitionCons.Messages():
				log.Printf("Got key `%s` with message `%s`\n", string(msg.Key), string(msg.Value))
				counter++
			case <-sigs:
				log.Println("Got interruption signal")
				done <- struct{}{}
			}
		}
	}()
	<-done
	log.Printf("Got %d messages\n", counter)
}
