package model

type ExecTime struct {
	StartTime int64 `bson:"startTime"`
	EndTIme   int64 `bsin:"EndTIme"`
}
type LogRecord struct {
	JobName string `bson:"jobName"`
	Command string `bson:command`
}
