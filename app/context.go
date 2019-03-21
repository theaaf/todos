package app

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/theaaf/todos/db"
	"github.com/theaaf/todos/model"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
	User          *model.User
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}

func (ctx *Context) WithUser(user *model.User) *Context {
	ret := *ctx
	ret.User = user
	return &ret
}

func (ctx *Context) AuthorizationError() *UserError {
	return &UserError{Message: "unauthorized", StatusCode: http.StatusForbidden}
}
