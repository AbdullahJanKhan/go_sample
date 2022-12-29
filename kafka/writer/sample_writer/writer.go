package samplewriter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/abdullahjankhan/go_sample/models"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/abdullahjankhan/go_sample/utils"
	"github.com/segmentio/kafka-go"
)

type KafkaSampleWriter interface {
	Submit(log *models.SampleRequest) error
}
type kafkaSampleWriter struct {
	Logger utils.Logger
	Config service.GlobleConfigService
}

func NewKafkaSampleWriter(Logger utils.Logger, Config service.GlobleConfigService) KafkaSampleWriter {
	return &kafkaSampleWriter{
		Logger: Logger,
		Config: Config,
	}
}

func newSampleLog(log *models.SampleRequest) *SampleLog {
	return &SampleLog{
		Base: Base{
			Type: LogTypeSampleLog,
			Time: time.Now(),
		},
		Name:     log.Name,
		ID:       log.ID,
		Email:    log.Email,
		Password: log.Password,
	}
}

// SubmitUser submit or write user on kafka
func (w *kafkaSampleWriter) Submit(log *models.SampleRequest) error {

	buf, err := json.Marshal(newSampleLog(log))
	if err != nil {
		w.Logger.Fatalf("Unable to marshal log%v", err)
		return err
	}
	err = w.getWriter().WriteMessages(context.Background(), kafka.Message{Value: buf})
	if err != nil {
		w.Logger.Fatalf("Log Writing Failed%v", err)
		return err
	}
	return nil
}

// getGoldenTicketWriter returns new kafka writer instance
func (w *kafkaSampleWriter) getWriter() *kafka.Writer {
	writer, found := kafkaId2Writer.Load(KafkaSampleReqTopic)
	if found {
		return writer.(*kafka.Writer)
	}

	config := w.Config.GetConfig()
	newWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      config.Kafka.Brokers,
		Topic:        KafkaSampleReqTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
		Logger:       w.Logger,
	})
	kafkaId2Writer.Store(KafkaSampleReqTopic, newWriter)
	return newWriter
}
