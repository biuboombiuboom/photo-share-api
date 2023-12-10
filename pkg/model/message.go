package model

import "time"

type Message struct {
	Id           int64
	Content      string
	SendUserId   int64
	SendUsername string
	ToUserId     int64
	ToUsername   int64
	SendTime     time.Time
	ReadState    bool
}
