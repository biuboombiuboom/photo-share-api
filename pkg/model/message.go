package model

import "time"

type Message struct {
	Id           int64
	Content      string
	FromUserId   int64
	FromUsername string
	ToUserId     int64
	ToUsername   int64
	CreatedAt    time.Time
	SendTime     time.Time
	ReadState    bool
}
