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
	Star        int64     `json:"star"`
	Like        int64     `json:"like"`
	Comment     int64     `json:"comment"`
}

type PhotoComment struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	UserName  string    `json:"username"`
	PhotoId   int64     `json:"photoId"`
	ReplyTo   int64     `json:"replyTo"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}

type PhotoCollect struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	PhotoId   int64     `json:"photoId"`
}

type PhotoStar struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	PhotoId   int64     `json:"photoId"`
}

type PhotoDTO struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	UserName    string `json:"username"`
	Path        string `json:"path"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	Star        int64  `json:"star"`
	Like        int64  `json:"like"`
	Comment     int64  `json:"comment"`
}

type PageQuery struct {
	OrderBy  string `json:"orderby"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}
