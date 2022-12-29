# Kafka Workers

Labour doers of kafka logs. The workers are programmed to process the business logic binded to a given log.

They have a simpe purpose to apply the business logic and commit the message.

### How to Commit a message?

It is simple but necessary to avoid re-reading of logs.

```go

	err := reader.CommitMessages(context.Background(), msg `// Kafka Message`)
    if err != nil {
        log.Error(err)
    }

```
