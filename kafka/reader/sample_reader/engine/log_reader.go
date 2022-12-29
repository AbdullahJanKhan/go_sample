package engine

import (
	"context"
	"encoding/json"

	"github.com/abdullahjankhan/go_sample/utils"
	"github.com/segmentio/kafka-go"
)

type LogObserver interface {
	OnSampleRequestLog(log *SampleRequestLog, msg kafka.Message)
}

type LogReader interface {
	RegisterObserver(observer LogObserver)
	Run()
	CommitMessage(kafka.Message)
}

type KafkaLogReader struct {
	readerId  string
	productId string
	reader    *kafka.Reader
	observer  LogObserver
	Logger    utils.Logger
}

func NewKafkaLogReader(readerId string, brokers []string, logger utils.Logger) LogReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     KafkaSampleReqTopic,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
		Logger:    logger,
	})
	return &KafkaLogReader{
		readerId: readerId,
		reader:   reader,
		Logger:   logger,
	}
}

func (r *KafkaLogReader) RegisterObserver(observer LogObserver) {
	r.observer = observer
}

func (r *KafkaLogReader) Run() {
	r.Logger.Infof("%v:%v started read", r.productId, r.readerId)

	for {
		kMessage, err := r.reader.ReadMessage(context.Background())
		if err != nil {
			r.Logger.Error(err)
			continue
		}

		var base Base
		err = json.Unmarshal(kMessage.Value, &base)
		if err != nil {
			r.Logger.Error(err)
			continue
		}

		r.Logger.Infof("fetch Kafka log reader %v", base)
		r.Logger.Infof("Run log %v", base)

		switch base.Type {

		case LogTypeSampleRequestDone:
			var log SampleRequestLog
			err := json.Unmarshal(kMessage.Value, &log)
			if err != nil {
				panic(err)
			}
			r.observer.OnSampleRequestLog(&log, kMessage)
		}
	}
}

func (r *KafkaLogReader) CommitMessage(message kafka.Message) {
	err := r.commitMessage(retires, message)
	if err != nil {
		r.Logger.Errorf("unable to commit kafka log: %v\nError:%v", message, err)
	} else {
		r.Logger.Infof("kafka log %v committed successfully", message)
	}
}

func (r *KafkaLogReader) commitMessage(count int, msg kafka.Message) error {
	if count > 0 {
		err := r.reader.CommitMessages(context.Background(), msg)
		if err != nil {
			r.commitMessage(count-1, msg)
		}
		return err
	}
	return nil
}
