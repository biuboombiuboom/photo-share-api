package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	"photo.share/pkg/model"
	"photo.share/pkg/store"
)

func Login(ctx context.Context, login string, pwd string) (model.User, error) {
	// hash := md5.Sum([]byte(pwd))
	encodePwd := pwd
	fmt.Printf("%s,%s\n", login, pwd)
	user := model.User{}
	empty := model.User{}

	sql := "select id,username,email,password from pps.user where (username=? or email=?)"
	args := make([]interface{}, 0)
	args = append(args, login, login)

	row := store.DB.QueryRowContext(ctx, sql, args...)
	err := row.Err()
	if err != nil {
		return user, err
	}
	var userPassword string
	row.Scan(&user.Id, &user.UserName, &user.Email, &userPassword)
	if encodePwd == userPassword {
		return user, nil
	} else {
		return empty, errors.New("错误的登录信息")
	}

}

func Exists(ctx context.Context, login string) (bool, error) {
	sql := "select count() c from pps.user where (username=? or email=?) "
	args := make([]interface{}, 0)
	args = append(args, login, login)
	rows, err := store.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func NewUser(ctx context.Context, user model.User) (int64, error) {
	hash := md5.Sum([]byte(user.Password))
	encodePwd := hex.EncodeToString(hash[:])

	sql := "insert into pps.user(username,email,password)value(?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, user.UserName, user.Email, encodePwd)
	fmt.Printf("%s,%s,%s\n", user.UserName, user.Email, encodePwd)
	result, err := store.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
