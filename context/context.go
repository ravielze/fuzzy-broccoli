package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

type (
	Parameters map[string]string
	Queries    map[string]interface{}
	Context    struct {
		ginCtx *gin.Context

		err      error
		params   Parameters
		code     string
		httpCode int
		query    Queries
		isError  bool
	}
)

func New(ctx *gin.Context) *Context {
	return &Context{
		ginCtx:   ctx,
		isError:  false,
		err:      nil,
		params:   Parameters{},
		code:     "",
		httpCode: -1,
		query:    Queries{},
	}
}

func (ctx *Context) ClientIP() string {
	return ctx.ginCtx.ClientIP()
}

func (ctx *Context) Gin() *gin.Context {
	return ctx.ginCtx
}

func (ctx *Context) Context() context.Context {
	return ctx.ginCtx.Request.Context()
}