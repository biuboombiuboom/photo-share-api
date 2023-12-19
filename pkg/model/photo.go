package model

import "time"

type Photo struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"userId"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Title       string    `json:"title:"`
	IsPublic    bool      `json:"isPublic"`
	MD5         string    `json:"md5"`
}

type PhotoDTO struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	UserName    string `json:"username"`
	Path        string `json:"path"`
	Title       string `json:"tile"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	Star        int64  `json:"start"`
}

type PageQuery struct {
	OrderBy  string `json:"orderby"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}
