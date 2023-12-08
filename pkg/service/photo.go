package service

import (
	"context"

	"photo.share/pkg/model"
	"photo.share/pkg/store"
)

func NewPhoto(ctx context.Context, photoInfo model.Photo) (model.Photo, error) {
	query := "insert int pps.photo(user_id,path,created_at,star,description,title,is_public,md5) value(?,?,?,?,?,?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, photoInfo.Path, photoInfo.UserId, photoInfo.CreatedAt, photoInfo.Star, photoInfo.Description, photoInfo.Title, photoInfo.IsPublic, "")

	result, err := store.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Photo{}, err
	}

	photoInfo.Id, _ = result.LastInsertId()
	return photoInfo, nil

}
