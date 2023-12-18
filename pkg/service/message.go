package service

import (
	"context"
	"strings"
	"time"

	"photo.share/pkg/model"
	"photo.share/pkg/store"
)

func NewMessage(ctx context.Context, message model.Message) error {
	sql := "insert into pps.message(content,from_user_id,from_username,to_user_id,to_username,created_at,send_time,read_state)values(?,?,?,?,?,?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, message.Content, message.FromUserId, message.FromUsername, message.ToUserId, message.ToUsername, time.Now(), time.Now(), false)
	_, err := store.DB.ExecContext(ctx, sql, args...)
	return err
}

func ReadMessages(ctx context.Context, messageIds []int64) error {
	sql := "update pps.message set read_state=true where id in"
	p := make([]string, len(messageIds))
	args := make([]interface{}, 0)
	for i, id := range messageIds {
		p[i] = "?"
		args = append(args, id)
	}
	sql = sql + "(" + strings.Join(p, ",") + ")"

	_, err := store.DB.ExecContext(ctx, sql, args...)
	return err
}

func ReadAllsMessages(ctx context.Context, to_user_id int64) error {
	sql := "update pps.message set read_state=true where to_user_id =?"

	_, err := store.DB.ExecContext(ctx, sql, to_user_id)
	return err
}
