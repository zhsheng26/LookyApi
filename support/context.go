package support

import (
	"github.com/kataras/iris"
	"sync"
)

/**
* @author zhangsheng
* @date 2019/6/17
 */

type UsContext struct {
	iris.Context
}

var contextPool = sync.Pool{New: func() interface{} {
	return &UsContext{}
}}

func acquire(original iris.Context) *UsContext {
	ctx := contextPool.Get().(*UsContext)
	ctx.Context = original
	return ctx
}

func release(ctx *UsContext) {
	contextPool.Put(ctx)
}
func Handler(h func(*UsContext)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

func (ctx *UsContext) UserId() string {
	header := ctx.GetHeader("userId")
	if header == "" {
		return "888888"
	}
	return header
}
