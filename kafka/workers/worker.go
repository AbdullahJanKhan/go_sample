package workers

import (
	"github.com/abdullahjankhan/go_sample/kafka/reader/sample_reader/engine"
	"github.com/abdullahjankhan/go_sample/utils"
	"github.com/segmentio/kafka-go"
)

type SampleLogReader struct {
	logReader engine.LogReader
	logger    utils.Logger
}

func NewSampleLogReader(logReader engine.LogReader, logger utils.Logger) *SampleLogReader {
	t := &SampleLogReader{
		logReader: logReader,
		logger:    logger,
	}

	t.logReader.RegisterObserver(t)
	return t
}

func (t *SampleLogReader) Start() {
	go t.logReader.Run()
}

func (t *SampleLogReader) OnSampleRequestLog(log *engine.SampleRequestLog, msg kafka.Message) {
	t.logger.Infof("Log Read and Processed %v", log)
	t.logReader.CommitMessage(msg)
}
