package model

type Contextsource interface {
	UserId() string
}

type Context struct {
	userId string
}

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) UserId() string {
	return ctx.userId
}
