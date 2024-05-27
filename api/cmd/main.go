package main

import (
	"context"

	"github.com/mori-hisayuki/go-sutdy/api/internal/presenter"
)

// @title ユーザー管理サービスAPI
// @version 0.1.0
// @description ユーザー管理サービスAPIの昨日
// @host localhost:8080
func main() {
	server := presenter.NewUserPresenter()
	if err := server.Run(context.Background()); err != nil {
		panic(err)
	}
}
