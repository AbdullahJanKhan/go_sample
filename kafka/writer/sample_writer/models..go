package samplewriter

import (
	"sync"
	"time"
)

type LogType string

var (
	kafkaId2Writer sync.Map
)

const (
	LogTypeSampleLog    = LogType("Sample_Request_Log")
	KafkaSampleReqTopic = "sample_requets"
)

type Base struct {
	Type LogType
	Time time.Time
}

type SampleLog struct {
	Base
	Name     string `json:"name" binding:"required"`
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
