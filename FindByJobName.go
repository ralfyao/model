package model

type FindMyJobName struct {
	JobName string `bson:"jobName"`
	Command string `bson:"command"`
}
