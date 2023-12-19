package service

import (
	"context"
	"fmt"
	"time"

	"photo.share/pkg/model"
	"photo.share/pkg/store"
)

func StarPhoto(ctx context.Context, star model.PhotoStar) (int, error) {
	tx, err := store.DB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return 0, err
	}
	query := "select count(1) from pps.photo_star where user_id=? and photo_id=?"
	args := []interface{}{star.UserId, star.PhotoId}
	row := tx.QueryRowContext(ctx, query, args...)
	count := 0
	row.Scan(&count)
	if count > 0 {
		return 0, nil
	}

	sql := "insert into pps.photo_star(user_id,username,photo_id,created_at)	value(?,?,?,?)"
	_, err = tx.ExecContext(ctx, sql, star.UserId, star.UserName, star.PhotoId, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	updateStarQuery := "update pps.photo set star=star+1 where id=?"
	_, err = tx.ExecContext(ctx, updateStarQuery, star.PhotoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func CollectPhoto(ctx context.Context, collect model.PhotoCollect) (int, error) {
	tx, err := store.DB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return 0, err
	}
	query := "select count(1) from pps.photo_collect where user_id=? and photo_id=?"
	args := []interface{}{collect.UserId, collect.PhotoId}
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Err()
	if err != nil {
		return 0, err
	}
	count := 0
	row.Scan(&count)
	if count > 0 {
		return 0, nil
	}

	sql := "insert into pps.photo_collect(user_id,username,photo_id,created_at)	value(?,?,?,?)"
	_, err = tx.ExecContext(ctx, sql, collect.UserId, collect.UserName, collect.PhotoId, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	updateStarQuery := "update pps.photo set collect=collect+1 where id=?"
	_, err = tx.ExecContext(ctx, updateStarQuery, collect.PhotoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return 1, nil
}

func GetPublicPhotos(ctx context.Context, orderby string, limit0, limit1 int64) ([]model.PhotoDTO, int, error) {
	total := 0
	returning := make([]model.PhotoDTO, 0)

	query := "select count(1) from pps.photo where is_public=true and deleted=false"

	row := store.DB.QueryRowContext(ctx, query)
	if err := row.Scan(&total); err != nil {
		return returning, total, err
	}

	sql := fmt.Sprintf("select p.id,p.path,p.title,p.description, p.user_id,u.username,p.star,p.collect from pps.photo as p inner join pps.user as u on p.user_id=u.id where p.is_public=true and p.deleted=false order by p.%s desc limit %d,%d", orderby, limit0, limit1)
	rows, err := store.DB.QueryContext(ctx, sql)
	if err != nil {
		return returning, total, err
	}
	defer rows.Close()
	for rows.Next() {
		p := model.PhotoDTO{}
		rows.Scan(&p.Id, &p.Path, &p.Title, &p.Description, &p.UserId, &p.UserName, &p.Star, &p.Like)
		returning = append(returning, p)
	}

	return returning, total, nil
}

func NewPhoto(ctx context.Context, photoInfo model.Photo) (model.Photo, error) {
	query := "insert into pps.photo(user_id,path,created_at,description,title,is_public,deleted) value(?,?,?,?,?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, photoInfo.UserId, photoInfo.Path, photoInfo.CreatedAt, photoInfo.Description, photoInfo.Title, photoInfo.IsPublic, false)

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
