package engine

import "time"

type LogType string

const (
	LogTypeSampleRequestDone = LogType("Sample_Request_Log")
	KafkaSampleReqTopic      = "sample_requets"
)

const (
	retires = 5
)

type Base struct {
	Type LogType
	Time time.Time
}

type SampleRequestLog struct {
	Base
	Name     string `json:"name" binding:"required"`
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
