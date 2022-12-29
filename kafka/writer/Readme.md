# Kafka Writer

Kafka writer also known as a producer. Writer are responsible to write messages on the stream that will be read / consumed by the respective reader / consumer.

## How to write a message / log?

We are using Go Lang Library `"github.com/segmentio/kafka-go"`

It offers an easy writing mechanism

```go
    writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      `[Your Kafka Brokers Addresses]`,
		Topic:        `Your Topic Name`,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
    err =writer.WriteMessages(context.Background(), kafka.Message{Value: [`You Log In Bytes`]})
	if err != nil {
		log.Fatalf("Log Writing Failed%v", err)
	}

```
