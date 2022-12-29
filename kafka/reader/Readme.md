# Kafka Reader

Kafka reader also known as consumers. When a consumer writes a log ont he stream, our consumer / reader reads the messsage and process the business logic required on the log.

## How to read a message / log?

We are using Go Lang Library `"github.com/segmentio/kafka-go"`

It offers an easy reading mechanism

```go
    reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   `[Kafka Brokers Address]`,
		Topic:     `Your Topic Name`,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
    kMessage, err := r.reader.ReadMessage(context.Background())
	if err != nil {
		log.Error(err)
	}

```

### A Good Practice

It is always good to divide your code base into managable positions, hence a reader / consumer should be responsible to read / consume the message forward it to a lobour doer and get work done from it. Hence a worker can be be used to process the business logic, as we know that multiple writers can write to a topic and to match the throughput we must have multiple workers processing the logs (incomming messages).
