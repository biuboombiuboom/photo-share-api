package model

import "time"

type Photo struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"userId"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"createdAt"`
	Star        int64     `json:"star"`
	Description string    `json:"description"`
	Title       string    `json:"title:"`
	IsPublic    bool      `json:"isPublic"`
	MD5         string    `json:"md5"`
}
