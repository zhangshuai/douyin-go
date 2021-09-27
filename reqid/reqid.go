package reqid

import (
	"context"
)

type reqidKey struct{}

// WithReqid 把reqid加入context中
func WithReqid(ctx context.Context, reqid string) context.Context {
	return context.WithValue(ctx, reqidKey{}, reqid)
}

// FromContext 从context中获取reqid
func FromContext(ctx context.Context) (reqid string, ok bool) {
	reqid, ok = ctx.Value(reqidKey{}).(string)
	return
}
