package service

import (
	"context"

	"photo.share/pkg/model"
	"photo.share/pkg/store"
)

func NewPhoto(ctx context.Context, photoInfo model.Photo) (model.Photo, error) {
	query := "insert into pps.photo(user_id,path,created_at,description,title,is_public) value(?,?,?,?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, photoInfo.UserId, photoInfo.Path, photoInfo.CreatedAt, photoInfo.Description, photoInfo.Title, photoInfo.IsPublic)

	result, err := store.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Photo{}, err
	}

	photoInfo.Id, err = result.LastInsertId()
	return photoInfo, err
}

func GetPhotosByUserId(ctx context.Context, userId int64) ([]model.PhotoDTO, error) {
	returning := make([]model.PhotoDTO, 0)
	query := "select id,user_id,title,path,description,is_public from pps.photo where deleted=false and user_id=?"
	rows, err := store.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return returning, err
	}
	defer rows.Close()
	for rows.Next() {
		p := model.PhotoDTO{}
		rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Path, &p.Description, &p.IsPublic)
		returning = append(returning, p)
	}

	return returning, nil
}

func DeletePhoto(ctx context.Context, photoId int64) error {
	sql := "update pps.photo set deleted=true where id =?"

	_, err := store.DB.ExecContext(ctx, sql, photoId)
	return err
}
